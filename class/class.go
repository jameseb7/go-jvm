package class

import "fmt"

const MAGIC = 0xCAFEBABE

type File struct {
	Magic        uint32
	MinorVersion uint16
	MajorVersion uint16
	ConstantPool []CPInfo
	AccessFlags  uint16
	ThisClass    uint16
	SuperClass   uint16
	Interfaces   []uint16
	Fields       []FieldInfo
	Methods      []MethodInfo
	Attributes   []AttributeInfo
}

type CPTag uint8

const NullTag CPTag = 0
const ClassTag CPTag = 7
const FieldRefTag CPTag = 9
const MethodRefTag CPTag = 10
const InterfaceMethodRefTag CPTag = 11
const StringTag CPTag = 8
const IntegerTag CPTag = 3
const FloatTag CPTag = 4
const LongTag CPTag = 5
const DoubleTag CPTag = 6
const NameAndTypeTag CPTag = 12
const Utf8Tag CPTag = 1
const MethodHandleTag CPTag = 15
const MethodTypeTag CPTag = 16
const InvokeDynamicTag CPTag = 18

type CPInfo struct {
	Tag  CPTag
	Info interface{}
}

type ConstantClassInfo struct {
	NameIndex uint16
}

type ConstantFieldRefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type ConstantMethodRefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type ConstantInterfaceMethodRefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type ConstantStringInfo struct {
	StringIndex uint16
}

type ConstantIntegerInfo struct {
	Bytes int32
}

type ConstantFloatInfo struct {
	Bytes float32
}

type ConstantLongInfo struct {
	Bytes int64
}

type ConstantDoubleInfo struct {
	Bytes float64
}

type ConstantNameAndTypeInfo struct {
	NameIndex       uint16
	DescriptorIndex uint16
}

type ConstantUtf8Info struct {
	Bytes string
}

type ConstantMethodHandleInfo struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}

type ConstantMethodTypeInfo struct {
	DescriptorIndex uint16
}

type ConstantInvokeDynamicInfo struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

type FieldInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []AttributeInfo
}

type MethodInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []AttributeInfo
}

func (t CPTag) String() string {
	switch t {
	case ClassTag:
		return "Class"
	case FieldRefTag:
		return "FieldRef"
	case MethodRefTag:
		return "MethodRef"
	case InterfaceMethodRefTag:
		return "InterfaceMethodRef"
	case StringTag:
		return "String"
	case IntegerTag:
		return "Integer"
	case FloatTag:
		return "Float"
	case LongTag:
		return "Long"
	case DoubleTag:
		return "Double"
	case NameAndTypeTag:
		return "NameAndType"
	case Utf8Tag:
		return "Utf8"
	case MethodHandleTag:
		return "MethodHandle"
	case MethodTypeTag:
		return "MethodType"
	case InvokeDynamicTag:
		return "InvokeDynamic"
	default:
		return fmt.Sprintf("%d", t)
	}
}
