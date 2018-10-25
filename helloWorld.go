package main

import (
	"awesomeProject/1023show"
	console "fmt"
	"unsafe"
)

func init() {
	console.Print("init函数\n")
}
func main() {
	console.Println("hello pojun")
	_023show.Show()
	/*var i uint8 = 2
	console.Print(unsafe.Sizeof(i))*/
	/*var j int32 = 1
	console.Print(unsafe.Sizeof(j))*/
	var i int = 3
	console.Print(unsafe.Sizeof(i))
	var j float32 = 2.4
	console.Print(unsafe.Sizeof(j))
	var b bool = false
	console.Print(b)
	console.Print('\n')
	var r rune = 3
	console.Print(r)
}
