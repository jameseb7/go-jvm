package main

import "fmt"
import "os"
import "github.com/jameseb7/go-jvm/class"
import "strings"
import "path/filepath"
import "errors"

func main() {
	if len(os.Args) < 2 {
		return
	}

	err := loadClass(os.Args[1], "", nil)
	if err != nil {
		panic(err)
	}

	class := classes[classNamePair{os.Args[1], ""}].classDef
	super := classes[classNamePair{class.SuperClassName(), ""}].classDef

	fmt.Println("")
	fmt.Println(class.Name())
	fmt.Println(class.SuperClassName())

	fmt.Println("")
	fmt.Println(super.Name())
	fmt.Println(super.SuperClassName())
}

type classNamePair struct{
	className string
	initiatingLoader string //"" denotes a class initiated by the bootstrap loader
}

type classDefPair struct{
	definingLoader string //"" denotes a class defined by the bootstrap loader
	classDef *class.File
}

var classes = make(map[classNamePair]classDefPair, 100)

func loadClass(name string, loader string, chain []string) (err error) {
	for _, v := range chain {
		if name == v {
			err = errors.New("ClassCircularityError")
			return
		}
	}
	
	if loader != "" {
		panic(errors.New("User-defined loaders are not yet supported"))
	}

	_, ok := classes[classNamePair{name, loader}]
	if ok {
		return //class already loaded, no loading required
	}

	//search for the class in the filesystem
	var file *os.File
	filepath1 := strings.Join([]string{name,".class"},"")
	file, err = os.Open(filepath1)
	if err != nil {
		classpath := os.Getenv("CLASSPATH")
		fmt.Println(classpath)
		classpaths := filepath.SplitList(classpath)
		fmt.Println(classpaths)
		for _, v := range classpaths {
			filepath2 := filepath.Join(v, filepath1)
			file, err = os.Open(filepath2)
			if file == nil {
				continue
			}
			err = nil
			break
		}
	}
	if err != nil {
		return
	}

	classFile, err := class.ReadFile(file)
	if err != nil {
		return
	}

	if  classFile.Name() != name {
		err = errors.New(fmt.Sprintf("NoClassDefFoundError: %v", name))
		return
	}

	if (classFile.MajorVersion > 52) || 
		((classFile.MajorVersion == 52) && (classFile.MinorVersion > 0)) {
		err = errors.New("UnsupportedClassVersionError")
		return
	}
	
	//load superclasses
	if classFile.SuperClass != 0 {
		if chain == nil {
			chain = make([]string, 5)
		}
		superclass := classFile.SuperClassName()
		loadClass(superclass, loader, append(chain, name))
		if classes[classNamePair{superclass, loader}].classDef.IsInterface() {
			err = errors.New("IncompatibleClassChangeError")
			return
		}
	} else {
		if classFile.Name() != "java/lang/Object" {
			err = errors.New("Only Object may have no superclass")
			return
		}
	}

	classes[classNamePair{name, loader}] = classDefPair{loader, &classFile}

	return
}
