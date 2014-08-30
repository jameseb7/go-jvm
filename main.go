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

	err := loadClass(os.Args[1], "")
	if err != nil {
		panic(err)
	}
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

func loadClass(name string, loader string) (err error) {
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

	classes[classNamePair{name, loader}] = classDefPair{loader, &classFile}

	return
}
