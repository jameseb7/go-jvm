package class

func (cf File) Name() string {
	thisIndex := cf.ThisClass
	classInfo, ok := cf.ConstantPool[thisIndex].Info.(ConstantClassInfo)
	if !ok {
		return ""
	}

	name, ok := cf.ConstantPool[classInfo.NameIndex].Info.(ConstantUtf8Info)
	if !ok {
		return ""
	}
	
	return name.Bytes
}

func (cf File) SuperClassName() string {
	superIndex := cf.SuperClass

	if superIndex == 0 {
		return ""
	}

	classInfo, ok := cf.ConstantPool[superIndex].Info.(ConstantClassInfo)
	if !ok {
		return ""
	}

	name, ok := cf.ConstantPool[classInfo.NameIndex].Info.(ConstantUtf8Info)
	if !ok {
		return ""
	}
	
	return name.Bytes
}

func (cf File) IsPublic() bool {
	return (cf.AccessFlags & 0x0001) != 0
}

func (cf File) IsFinal() bool {
	return (cf.AccessFlags & 0x0010) != 0
}

func (cf File) IsInterface() bool {
	return (cf.AccessFlags & 0x0200) != 0
}

func (cf File) IsAbstract() bool {
	return (cf.AccessFlags & 0x0400) != 0
}

func (cf File) IsSynthetic() bool {
	return (cf.AccessFlags & 0x1000) != 0
}

func (cf File) IsAnnotation() bool {
	return (cf.AccessFlags & 0x2000) != 0
}

func (cf File) IsEnum() bool {
	return (cf.AccessFlags & 0x4000) != 0
}
