package class

import "io"
import "encoding/binary"
import "errors"
import "fmt"

func ReadFile(r io.Reader) (file File, err error) {
	var count uint16
	binary.Read(r, binary.BigEndian, &file.Magic)
	if file.Magic != MAGIC {
		err = errors.New(fmt.Sprint("Invalid class file magic number: ", file.Magic))
		return
	}
	fmt.Printf("magic = %X \n", file.Magic)

	binary.Read(r, binary.BigEndian, &file.MinorVersion)
	fmt.Println("minor_version = ", file.MinorVersion)
	binary.Read(r, binary.BigEndian, &file.MajorVersion)
	fmt.Println("major_version = ", file.MajorVersion)

	binary.Read(r, binary.BigEndian, &count)
	fmt.Println("constant_pool_count = ", count)
	file.ConstantPool = make([]CPInfo, count)
	for i := 1; i < len(file.ConstantPool); i++ {
		binary.Read(r, binary.BigEndian, &file.ConstantPool[i].Tag)
		fmt.Println("constant_pool[", i, "].tag = ", file.ConstantPool[i].Tag)
		switch file.ConstantPool[i].Tag {
		case ClassTag:
			var info ConstantClassInfo
			binary.Read(r, binary.BigEndian, &info.NameIndex)
			fmt.Println("constant_pool[", i, "].name_index = ", info.NameIndex)
			file.ConstantPool[i].Info = info
		case FieldRefTag:
			var info ConstantFieldRefInfo
			binary.Read(r, binary.BigEndian, &info.ClassIndex)
			fmt.Println("constant_pool[", i, "].class_index = ", info.ClassIndex)
			binary.Read(r, binary.BigEndian, &info.NameAndTypeIndex)
			fmt.Println("constant_pool[", i, "].name_and_type_index = ", info.NameAndTypeIndex)
			file.ConstantPool[i].Info = info
		case MethodRefTag:
			var info ConstantMethodRefInfo
			binary.Read(r, binary.BigEndian, &info.ClassIndex)
			fmt.Println("constant_pool[", i, "].class_index = ", info.ClassIndex)
			binary.Read(r, binary.BigEndian, &info.NameAndTypeIndex)
			fmt.Println("constant_pool[", i, "].name_and_type_index = ", info.NameAndTypeIndex)
			file.ConstantPool[i].Info = info
		case InterfaceMethodRefTag:
			var info ConstantInterfaceMethodRefInfo
			binary.Read(r, binary.BigEndian, &info.ClassIndex)
			fmt.Println("constant_pool[", i, "].class_index = ", info.ClassIndex)
			binary.Read(r, binary.BigEndian, &info.NameAndTypeIndex)
			fmt.Println("constant_pool[", i, "].name_and_type_index = ", info.NameAndTypeIndex)
			file.ConstantPool[i].Info = info
		case StringTag:
			var info ConstantStringInfo
			binary.Read(r, binary.BigEndian, &info.StringIndex)
			fmt.Println("constant_pool[", i, "].string_index = ", info.StringIndex)
			file.ConstantPool[i].Info = info
		case IntegerTag:
			var info ConstantIntegerInfo
			binary.Read(r, binary.BigEndian, &info.Bytes)
			fmt.Println("constant_pool[", i, "].bytes = ", info.Bytes)
			file.ConstantPool[i].Info = info
		case FloatTag:
			var info ConstantFloatInfo
			binary.Read(r, binary.BigEndian, &info.Bytes)
			fmt.Println("constant_pool[", i, "].bytes = ", info.Bytes)
			file.ConstantPool[i].Info = info
		case LongTag:
			var info ConstantLongInfo
			binary.Read(r, binary.BigEndian, &info.Bytes)
			fmt.Println("constant_pool[", i, "].bytes = ", info.Bytes)
			file.ConstantPool[i].Info = info
			i++
		case DoubleTag:
			var info ConstantDoubleInfo
			binary.Read(r, binary.BigEndian, &info.Bytes)
			fmt.Println("constant_pool[", i, "].bytes = ", info.Bytes)
			file.ConstantPool[i].Info = info
			i++
		case NameAndTypeTag:
			var info ConstantNameAndTypeInfo
			binary.Read(r, binary.BigEndian, &info.NameIndex)
			fmt.Println("constant_pool[", i, "].name_index = ", info.NameIndex)
			binary.Read(r, binary.BigEndian, &info.DescriptorIndex)
			fmt.Println("constant_pool[", i, "].descriptor_index = ", info.DescriptorIndex)
			file.ConstantPool[i].Info = info
		case Utf8Tag:
			var info ConstantUtf8Info
			binary.Read(r, binary.BigEndian, &count)
			bytes := make([]byte, count)
			for j, _ := range bytes {
				binary.Read(r, binary.BigEndian, &bytes[j])
			}
			info.Bytes = string(bytes)
			fmt.Printf("constant_pool[ %v ].bytes = %q \n", i, info.Bytes)
			file.ConstantPool[i].Info = info
		case MethodHandleTag:
			var info ConstantMethodHandleInfo
			binary.Read(r, binary.BigEndian, &info.ReferenceKind)
			fmt.Println("constant_pool[", i, "].reference_kind = ", info.ReferenceKind)
			binary.Read(r, binary.BigEndian, &info.ReferenceIndex)
			fmt.Println("constant_pool[", i, "].reference_index = ", info.ReferenceIndex)
			file.ConstantPool[i].Info = info
		case MethodTypeTag:
			var info ConstantMethodTypeInfo
			binary.Read(r, binary.BigEndian, &info.DescriptorIndex)
			fmt.Println("constant_pool[", i, "].descriptor_index = ", info.DescriptorIndex)
			file.ConstantPool[i].Info = info
		case InvokeDynamicTag:
			var info ConstantInvokeDynamicInfo
			binary.Read(r, binary.BigEndian, &info.BootstrapMethodAttrIndex)
			fmt.Println("constant_pool[", i, "].bootstrap_method_attr_index = ", info.BootstrapMethodAttrIndex)
			binary.Read(r, binary.BigEndian, &info.NameAndTypeIndex)
			fmt.Println("constant_pool[", i, "].name_and_type_index = ", info.NameAndTypeIndex)
			file.ConstantPool[i].Info = info
		default:
			err = errors.New(fmt.Sprintf("Bad tag in class file: %v", file.ConstantPool[i].Tag))
			return
		}
	}

	binary.Read(r, binary.BigEndian, &file.AccessFlags)
	fmt.Printf("access_flags = %X \n", file.AccessFlags)
	binary.Read(r, binary.BigEndian, &file.ThisClass)
	fmt.Println("this_class = ", file.ThisClass)
	binary.Read(r, binary.BigEndian, &file.SuperClass)
	fmt.Println("super_class = ", file.SuperClass)

	binary.Read(r, binary.BigEndian, &count)
	fmt.Println("intefaces_count = ", count)
	file.Interfaces = make([]uint16, count)
	for i, _ := range file.Interfaces {
		binary.Read(r, binary.BigEndian, &file.Interfaces[i])
		fmt.Println("intefaces[", i, "] = ", file.Interfaces[i])
	}

	binary.Read(r, binary.BigEndian, &count)
	file.Fields = make([]FieldInfo, count)
	for i, _ := range file.Fields {
		binary.Read(r, binary.BigEndian, &file.Fields[i].AccessFlags)
		binary.Read(r, binary.BigEndian, &file.Fields[i].NameIndex)
		binary.Read(r, binary.BigEndian, &file.Fields[i].DescriptorIndex)
		binary.Read(r, binary.BigEndian, &count)
		file.Fields[i].Attributes = make([]AttributeInfo, count)
		for j, _ := range file.Fields[i].Attributes {
			file.Fields[i].Attributes[j], err = readAttribute(r, file.ConstantPool)
			if err != nil {
				return
			}
		}
	}
	fmt.Println("Fields read")

	binary.Read(r, binary.BigEndian, &count)
	file.Methods = make([]MethodInfo, count)
	for i, _ := range file.Methods {
		binary.Read(r, binary.BigEndian, &file.Methods[i].AccessFlags)
		binary.Read(r, binary.BigEndian, &file.Methods[i].NameIndex)
		binary.Read(r, binary.BigEndian, &file.Methods[i].DescriptorIndex)
		binary.Read(r, binary.BigEndian, &count)
		file.Methods[i].Attributes = make([]AttributeInfo, count)
		for j, _ := range file.Methods[i].Attributes {
			file.Methods[i].Attributes[j], err = readAttribute(r, file.ConstantPool)
			if err != nil {
				return
			}
		}
	}

	binary.Read(r, binary.BigEndian, &count)
	file.Attributes = make([]AttributeInfo, count)
	for i, _ := range file.Attributes {
		file.Attributes[i], err = readAttribute(r, file.ConstantPool)
		if err != nil {
			return
		}
	}

	return
}
