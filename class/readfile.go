package class

import "io"
import "encoding/binary"
import "errors"
import "fmt"

func ReadFile(r io.Reader) (file File, err error){
	var count uint16
	binary.Read(r, binary.BigEndian, &file.Magic)
	if file.Magic != MAGIC {
		err = errors.New(fmt.Sprint("Invalid class file magic number: ", file.Magic))
		return
	}
	
	binary.Read(r, binary.BigEndian, &file.MinorVersion)
	binary.Read(r, binary.BigEndian, &file.MajorVersion)
	
	binary.Read(r, binary.BigEndian, &count)
	file.ConstantPool = make([]CPInfo, count)
	for i,_ := range file.ConstantPool  {
		if i == 0 {
			continue
		}
		binary.Read(r, binary.BigEndian, &file.ConstantPool[i].Tag)
		switch file.ConstantPool[i].Tag {
		case ClassTag:
			var info ConstantClassInfo
			binary.Read(r, binary.BigEndian, &info.NameIndex)
			file.ConstantPool[i].Info = info
		case FieldRefTag:
			var info ConstantFieldRefInfo
			binary.Read(r, binary.BigEndian, &info.ClassIndex)
			binary.Read(r, binary.BigEndian, &info.NameAndTypeIndex)
			file.ConstantPool[i].Info = info
		case MethodRefTag:
			var info ConstantMethodRefInfo
			binary.Read(r, binary.BigEndian, &info.ClassIndex)
			binary.Read(r, binary.BigEndian, &info.NameAndTypeIndex)
			file.ConstantPool[i].Info = info
		case InterfaceMethodRefTag:
			var info ConstantInterfaceMethodRefInfo
			binary.Read(r, binary.BigEndian, &info.ClassIndex)
			binary.Read(r, binary.BigEndian, &info.NameAndTypeIndex)
			file.ConstantPool[i].Info = info
		case StringTag:
			var info ConstantStringInfo
			binary.Read(r, binary.BigEndian, &info.StringIndex)
			file.ConstantPool[i].Info = info
		case IntegerTag:
			var info ConstantIntegerInfo
			binary.Read(r, binary.BigEndian, &info.Bytes)
			file.ConstantPool[i].Info = info
		case FloatTag:
			var info ConstantFloatInfo
			binary.Read(r, binary.BigEndian, &info.Bytes)
			file.ConstantPool[i].Info = info
		case LongTag:
			var info ConstantLongInfo
			binary.Read(r, binary.BigEndian, &info.HighBytes)
			binary.Read(r, binary.BigEndian, &info.LowBytes)
			file.ConstantPool[i].Info = info
		case DoubleTag:
			var info ConstantDoubleInfo
			binary.Read(r, binary.BigEndian, &info.HighBytes)
			binary.Read(r, binary.BigEndian, &info.LowBytes)
			file.ConstantPool[i].Info = info
		case NameAndTypeTag:
			var info ConstantNameAndTypeInfo
			binary.Read(r, binary.BigEndian, &info.NameIndex)
			binary.Read(r, binary.BigEndian, &info.DescriptorIndex)
			file.ConstantPool[i].Info = info
		case Utf8Tag:
			var info ConstantUtf8Info
			binary.Read(r, binary.BigEndian, &count)
			info.Bytes = make([]byte, count)
			for j,_ := range info.Bytes {
				binary.Read(r, binary.BigEndian, &info.Bytes[j])
			}
			file.ConstantPool[i].Info = info
		case MethodHandleTag:
			var info ConstantMethodHandleInfo
			binary.Read(r, binary.BigEndian, &info.ReferenceKind)
			binary.Read(r, binary.BigEndian, &info.ReferenceIndex)
			file.ConstantPool[i].Info = info
		case MethodTypeTag:
			var info ConstantMethodTypeInfo
			binary.Read(r, binary.BigEndian, &info.DescriptorIndex)
			file.ConstantPool[i].Info = info
		case InvokeDynamicTag:
			var info ConstantInvokeDynamicInfo
			binary.Read(r, binary.BigEndian, &info.BootstrapMethodAttrIndex)
			binary.Read(r, binary.BigEndian, &info.NameAndTypeIndex)
			file.ConstantPool[i].Info = info
		default:
			err = errors.New("Bad tag in class file")
			return
		}
	}

	binary.Read(r, binary.BigEndian, &file.AccessFlags)
	binary.Read(r, binary.BigEndian, &file.ThisClass)
	binary.Read(r, binary.BigEndian, &file.SuperClass)
	
	binary.Read(r, binary.BigEndian, &count)
	file.Interfaces = make([]uint16, count)
	for i,_ := range file.Interfaces {
		binary.Read(r, binary.BigEndian, &file.Interfaces[i])
	}

	binary.Read(r, binary.BigEndian, &count)
	file.Fields = make([]FieldInfo, count)
	for i,_ := range file.Fields {
		binary.Read(r, binary.BigEndian, &file.Fields[i].AccessFlags)
		binary.Read(r, binary.BigEndian, &file.Fields[i].NameIndex)
		binary.Read(r, binary.BigEndian, &file.Fields[i].DescriptorIndex)
		binary.Read(r, binary.BigEndian, &count)
		file.Fields[i].Attributes = make([]AttributeInfo, count)
		for j,_ := range file.Fields[i].Attributes {
			binary.Read(r, binary.BigEndian, 
				&file.Fields[i].Attributes[j].AttributeNameIndex)
			binary.Read(r, binary.BigEndian, &count)
			file.Fields[i].Attributes[j].Info = make([]byte, count)
			for k,_ := range file.Fields[i].Attributes[j].Info {
				binary.Read(r, binary.BigEndian, 
					&file.Fields[i].Attributes[j].Info[k])
			}
		}
	}

	binary.Read(r, binary.BigEndian, &count)
	file.Methods = make([]MethodInfo, count)
	for i,_ := range file.Methods {
		binary.Read(r, binary.BigEndian, &file.Methods[i].AccessFlags)
		binary.Read(r, binary.BigEndian, &file.Methods[i].NameIndex)
		binary.Read(r, binary.BigEndian, &file.Methods[i].DescriptorIndex)
		binary.Read(r, binary.BigEndian, &count)
		file.Methods[i].Attributes = make([]AttributeInfo, count)
		for j,_ := range file.Methods[i].Attributes {
			binary.Read(r, binary.BigEndian, 
				&file.Methods[i].Attributes[j].AttributeNameIndex)
			binary.Read(r, binary.BigEndian, &count)
			file.Methods[i].Attributes[j].Info = make([]byte, count)
			for k,_ := range file.Methods[i].Attributes[j].Info {
				binary.Read(r, binary.BigEndian, 
					&file.Methods[i].Attributes[j].Info[k])
			}
		}
	}

	binary.Read(r, binary.BigEndian, &count)
	file.Attributes = make([]AttributeInfo, count)
	for i,_ := range file.Attributes {
		binary.Read(r, binary.BigEndian, &file.Attributes[i].AttributeNameIndex)
		binary.Read(r, binary.BigEndian, &count)
		file.Attributes[i].Info = make([]byte, count)
		for j,_ := range file.Attributes[i].Info {
			binary.Read(r, binary.BigEndian, &file.Attributes[i].Info[j])
		}
	}
	
	return
}
