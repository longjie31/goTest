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

type Vertex struct {
	X int
	y string
}

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
	k, l := 31, 47
	n := &k         // 指向k
	fmt.Println(*n) // 通过n打出k的值
	*n = 33         // 通过n设置k的值
	fmt.Println(k)  // 打出k
	m := &l         // 指向l
	*m = *m / 47    // 通过n设置l的值
	fmt.Println(l)  // 打出l
	fmt.Println(Vertex{1, "pojun"})
	v := Vertex{1, "muyan"}
	v.X = 2
	fmt.Println(v.X)
	v.y = "jiansheng"
	fmt.Println(v.y)
	o := &Vertex{1, "libai"}
	o.X = 31
	fmt.Println(o)
	var array1 [2]string
	array1[0] = "hello"
	array1[1] = "world"
	fmt.Println(array1[0], array1[1])
	fmt.Println(array1)
	array2 := [3]int{21, 31, 47}
	fmt.Println(array2)
	var array3 []int = array2[0:2]
	fmt.Println(array3)
	// 切片
	array4 := [4]string{
		"hello", "world", "libai", "dufu",
	}
	fmt.Println(array4)
	array41 := array4[0:2]
	array42 := array4[1:3]
	fmt.Println(array41, array42)
	array42[0] = "XXX"
	fmt.Println(array41, array42)
	fmt.Println(array4)
	qiepian1 := []int{2, 3, 4, 5, 6}
	fmt.Println(qiepian1)
	qiepian2 := []bool{true, false, true, false, true}
	fmt.Println(qiepian2)
	str1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{31, true},
	}
	fmt.Println(str1)
	qiepian3 := []int{2, 4, 5, 6, 7, 31}
	qiepian3 = qiepian3[1:4]
	fmt.Println(qiepian3)
	qiepian3 = qiepian3[:3]
	fmt.Println(qiepian3)
	qiepian3 = qiepian3[1:]
	fmt.Println(qiepian3)
	qiepian4 := []int{2, 4, 6, 7, 8, 31, 47}
	printSlice(qiepian4)
	// 截取切片使其长度为0
	printSlice(qiepian4[:0])
	// 拓展其长度
	printSlice(qiepian4[:4])
	// 舍弃前两个值
	printSlice(qiepian4[2:])
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
