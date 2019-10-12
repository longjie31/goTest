package main

import (
	"fmt"
	"reflect"
)

var a int
var b int = 456
var (
	c string = "pojun"
	d int    = 1993
)

func main() {
	a = 123
	e, f, g := 1, 2, 3
	fmt.Print(a)
	fmt.Print(b, "\n")
	fmt.Print(c, "\n")
	fmt.Print(d, "\n")
	fmt.Print(e, f, g)
	var h int = 4
	var i = float32(h)
	fmt.Print(i)
	fmt.Print(reflect.TypeOf(i))
	fmt.Print("\n")
	var j float32 = 3.14
	k := int32(j)
	fmt.Print(k)
	fmt.Print(reflect.TypeOf(k))
	var a1, b1, c1 = true, false, "分别赋值"
	fmt.Println(a1, b1, c1)
	for i := 0; i < 4; i++ {
		fmt.Println("for循环", i)
	}
}
