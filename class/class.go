package class

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
	Bytes uint16
}

type ConstantFloatInfo struct {
	Bytes uint16
}

type ConstantLongInfo struct {
	HighBytes uint32
	LowBytes  uint32
}

type ConstantDoubleInfo struct {
	HighBytes uint32
	LowBytes  uint32
}

type ConstantNameAndTypeInfo struct {
	NameIndex       uint16
	DescriptorIndex uint16
}

type ConstantUtf8Info struct {
	Bytes []byte
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

type AttributeInfo struct {
	AttributeNameIndex uint16
	Info               []byte
}
