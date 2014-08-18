package class

import "fmt"

type AttributeInfo struct {
	AttributeNameIndex uint16
	Info               interface{}
}

type ConstantValueAttribute struct {
	ConstantValueIndex uint16
}

type CodeAttribute struct {
	MaxStack       uint16
	MaxLocals      uint16
	Code           []byte
	ExceptionTable []ExceptionTableEntry
	Attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	StartPC   uint16
	EndPC     uint16
	HandlerPC uint16
	CatchType uint16
}

type StackMapTableAttribute struct {
	Entries []StackMapFrame
}

type ItemTag uint8

const (
	ItemTop               ItemTag = 0
	ItemInteger           ItemTag = 1
	ItemFloat             ItemTag = 2
	ItemNull              ItemTag = 5
	ItemUninitializedThis ItemTag = 6
	ItemObject            ItemTag = 7
	ItemUninitialized     ItemTag = 8
	ItemLong              ItemTag = 4
	ItemDouble            ItemTag = 3
)

func (it ItemTag) String() string {
	switch it {
	case ItemTop:
		return "Top"
	case ItemInteger:
		return "Integer"
	case ItemFloat:
		return "Float"
	case ItemNull:
		return "Null"
	case ItemUninitializedThis:
		return "UninitializedThis"
	case ItemObject:
		return "Object"
	case ItemUninitialized:
		return "Uninitialized"
	case ItemLong:
		return "Long"
	case ItemDouble:
		return "Double"
	default:
		return fmt.Sprintf("%d", it)
	}
}

type VerificationTypeInfo struct {
	Tag  ItemTag
	Info interface{}
}

type ObjectVariableInfo struct {
	CpoolIndex uint16
}

type UninitializedVariableInfo struct {
	Offset uint16
}

type FrameTypeType uint8

func (ft FrameTypeType) String() string {
	if ft < 64 {
		return fmt.Sprintf("Same, OffsetDelta = %d", ft)
	}
	if ft < 128 {
		return fmt.Sprintf("SameLocals1StackItem, OffsetDelta = %d", ft-64)
	}
	if ft < 247 {
		return fmt.Sprintf("Reserved %d", ft)
	}
	if ft == 247 {
		return "SameLocals1StackItemExtended"
	}
	if ft < 251 {
		return fmt.Sprintf("Chop %d", 251-ft)
	}
	if ft == 251 {
		return "SameFrameExtended"
	}
	if ft < 255 {
		return fmt.Sprintf("Append %d", ft-251)
	}
	return "FullFrame"
}

type StackMapFrame struct {
	FrameType FrameTypeType
	Info      interface{}
}

type SameLocals1StackItemFrame struct {
	Stack [1]VerificationTypeInfo
}

type SameLocals1StackItemFrameExtended struct {
	OffsetDelta uint16
	Stack       [1]VerificationTypeInfo
}

type ChopFrame struct {
	OffsetDelta uint16
}

type SameFrameExtended struct {
	OffsetDelta uint16
}

type AppendFrame struct {
	OffsetDelta uint16
	Locals      []VerificationTypeInfo
}

type FullFrame struct {
	OffsetDelta uint16
	Locals      []VerificationTypeInfo
	Stack       []VerificationTypeInfo
}

type ExceptionsAttribute struct {
	ExceptionIndexTable []uint16
}

type BootstrapMethodsAttribute struct {
	BootstrapMethods []struct {
		BootstrapMethodRef uint16
		BootstrapArguments []uint16
	}
}
