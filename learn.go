package main

import (
	"fmt"
)
const NAME = "learn"// 常量一般建议大写
var name = "lj" // 在main函数外面定义的变量就是全局变量，在包里调用
// 单行注释
/*多行注释*/
func main() {
	fmt.Println("learn1")
	fmt.Println(NAME)
	fmt.Println(name)
}
