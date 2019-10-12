package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func add(x, y int) int {
	return x*2 + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var (
	tobe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Print('\n')
	v := 47
	fmt.Printf("v is of type %T\n", v)
	fmt.Print('\n')
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)
	/*const f = "%T(%v)\n"
	fmt.Printf(f, tobe, tobe)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, z, z)
	fmt.Print('\n')
	var i int
	fmt.Println(i, c, python, java)
	fmt.Print('\n')
	fmt.Println(split(31))
	fmt.Print('\n')
	a, b := swap("hello", "world")
	fmt.Print(a, b)
	fmt.Println(add(31, 47))
	fmt.Println(math.Pi);
	fmt.Printf("现在有 %g 项目", math.Nextafter(2, 3))
	fmt.Println("我最喜欢的数字是：", rand.Intn(10))
	fmt.Println("欢迎来到练习场")
	fmt.Println("现在时间是：", time.Now())*/
}
