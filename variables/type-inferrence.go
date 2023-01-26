package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name = "Lucas"
	var age = 22
	fmt.Println("Hello", name, "you are", age, "years old")

	// this works too, but it's not recommended because the compiler doesn't really know the type
	var version = 1.18
	fmt.Println("This is version", version)

	fmt.Println("----------")
	fmt.Println("name:", reflect.TypeOf(name))
	fmt.Println("age:", reflect.TypeOf(age))
	fmt.Println("version:", reflect.TypeOf(version))
}
