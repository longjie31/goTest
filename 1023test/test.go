package main

import (
	"fmt"
	"reflect"
)
type pojun int32

func main() {
	var i int32
	var j float32
	var b bool
	var d complex64
	fmt.Print(i)
	fmt.Print("\n")
	fmt.Print(j)
	fmt.Print("\n")
	fmt.Print(b)
	fmt.Print("\n")
	fmt.Print(d)
	fmt.Print("\n")
	 var c pojun
	var a int32
	fmt.Print(c)
	fmt.Print(a)
	fmt.Print("\n")
	fmt.Print(reflect.TypeOf(c))
	
}
