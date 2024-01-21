package main

import (
	"fmt"
	"reflect"
)

func main() {
	var exampleOne string
	exampleTwo := ""
	var exampleThree int

	exampleThree++

	fmt.Println(reflect.TypeOf(exampleOne))
	fmt.Println(reflect.TypeOf(exampleTwo))
	fmt.Println(reflect.TypeOf(exampleThree))
	fmt.Println(exampleThree)
}
