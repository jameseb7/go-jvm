package main

import "fmt"
import "os"
import "github.com/jameseb7/go-jvm/class"

func main(){
	if len(os.Args) < 2 {
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	classFile, err := class.ReadFile(file)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(classFile)
}
