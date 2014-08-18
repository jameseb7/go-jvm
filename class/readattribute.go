package class

import "io"
import "encoding/binary"
import "errors"
import "fmt"

func readAttribute(r io.Reader, cp []CPInfo) (ai AttributeInfo, err error) {
	binary.Read(r, binary.BigEndian, &ai.AttributeNameIndex)
	fmt.Println("AttributeNameIndex =", ai.AttributeNameIndex)

	anc := cp[ai.AttributeNameIndex]
	if anc.Tag != Utf8Tag {
		err = errors.New(fmt.Sprint("Invalid attribute name tag ", anc.Tag))
		return
	}
	an := anc.Info.(ConstantUtf8Info).Bytes
	fmt.Println("AttributeName =", an)

	var size uint32
	binary.Read(r, binary.BigEndian, &size)

	if an == "ConstantValue" {
		info := ConstantValueAttribute{}
		binary.Read(r, binary.BigEndian, &info.ConstantValueIndex)

		ai.Info = info
		return
	}

	if an == "Code" {
		info := CodeAttribute{}
		binary.Read(r, binary.BigEndian, &info.MaxStack)
		binary.Read(r, binary.BigEndian, &info.MaxLocals)

		var codeLength uint32
		binary.Read(r, binary.BigEndian, &codeLength)
		info.Code = make([]byte, codeLength)
		binary.Read(r, binary.BigEndian, &info.Code)

		var exceptionTableLength uint16
		binary.Read(r, binary.BigEndian, &exceptionTableLength)
		info.ExceptionTable = make([]ExceptionTableEntry, exceptionTableLength)
		binary.Read(r, binary.BigEndian, &info.ExceptionTable)

		var attributesCount uint16
		binary.Read(r, binary.BigEndian, &attributesCount)
		info.Attributes = make([]AttributeInfo, attributesCount)
		for i, _ := range info.Attributes {
			info.Attributes[i], err = readAttribute(r, cp)
			if err != nil {
				return
			}
		}

		ai.Info = info
		return
	}

	if an == "StackMapTable" {
		info := StackMapTableAttribute{}
		var numberOfEntries uint16
		binary.Read(r, binary.BigEndian, &numberOfEntries)
		info.Entries = make([]StackMapFrame, numberOfEntries)
		for i, _ := range info.Entries {
			var ft FrameTypeType
			binary.Read(r, binary.BigEndian, &ft)
			info.Entries[i].FrameType = ft
			if ft < 64 {
				//FrameType = SAME - no more to do
				continue
			}
			if ft < 128 {
				//FrameType = SAME_LOCALS_1_STACK_ITEM
				frame := SameLocals1StackItemFrame{}
				frame.Stack[0] = readVerificationTypeInfo(r)
				info.Entries[i].Info = frame
				continue
			}
			if ft < 247 {
				//FrameType = RESERVED
				continue
			}
			if ft == 247 {
				//FrameType = SAME_LOCALS_1_STACK_ITEM_EXTENDED
				frame := SameLocals1StackItemFrameExtended{}
				binary.Read(r, binary.BigEndian, &frame.OffsetDelta)
				frame.Stack[0] = readVerificationTypeInfo(r)
				info.Entries[i].Info = frame
				continue
			}
			if ft < 251 {
				//FrameType = CHOP
				frame := ChopFrame{}
				binary.Read(r, binary.BigEndian, &frame.OffsetDelta)
				info.Entries[i].Info = frame
				continue
			}
			if ft == 251 {
				//FrameType = SAME_FRAME_EXTENDED
				frame := SameFrameExtended{}
				binary.Read(r, binary.BigEndian, &frame.OffsetDelta)
				info.Entries[i].Info = frame
				continue
			}
			if ft < 255 {
				//FrameType = APPEND_FRAME
				frame := AppendFrame{}
				binary.Read(r, binary.BigEndian, &frame.OffsetDelta)
				frame.Locals = make([]VerificationTypeInfo, ft-251)
				for j, _ := range frame.Locals {
					frame.Locals[j] = readVerificationTypeInfo(r)
				}
				info.Entries[i].Info = frame
				continue
			}
			//FrameType = FULL_FRAME
			frame := FullFrame{}
			binary.Read(r, binary.BigEndian, &frame.OffsetDelta)
			var localsSize uint16
			frame.Locals = make([]VerificationTypeInfo, localsSize)
			for j, _ := range frame.Locals {
				frame.Locals[j] = readVerificationTypeInfo(r)
			}
			var stackSize uint16
			frame.Stack = make([]VerificationTypeInfo, stackSize)
			for j, _ := range frame.Stack {
				frame.Stack[j] = readVerificationTypeInfo(r)
			}
			info.Entries[i].Info = frame
			continue
		}

		ai.Info = info
		return
	}

	//default handling for an unrecognised attribute
	info := make([]byte, size)
	for i, _ := range info {
		binary.Read(r, binary.BigEndian, &info[i])
	}
	ai.Info = info
	return
}

func readVerificationTypeInfo(r io.Reader) VerificationTypeInfo {
	vtinfo := VerificationTypeInfo{}
	binary.Read(r, binary.BigEndian, &vtinfo.Tag)
	switch vtinfo.Tag {
	case ItemObject:
		info := ObjectVariableInfo{}
		binary.Read(r, binary.BigEndian, &info.CpoolIndex)
		vtinfo.Info = info
	case ItemUninitialized:
		info := UninitializedVariableInfo{}
		binary.Read(r, binary.BigEndian, &info.Offset)
		vtinfo.Info = info
	}
	return vtinfo
}
