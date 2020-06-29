// Author: Tong Zicheng
// Date: 2020/6/18 23:02
// Describe: fmt包的详细说明

package basics

import (
	"fmt"
	"testing"
)

func Test_fmt(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	user := User{
		"let-it-go",
		99,
	}

	fmt.Printf("%%\n")                   // %
	fmt.Printf("%b\n", 16)               // 10000
	fmt.Printf("%c\n", 65)               // A
	fmt.Printf("%c\n", 0x4f60)           // 你
	fmt.Printf("%U\n", '你')              // U+4f60
	fmt.Printf("%x\n", '你')              // 4f60
	fmt.Printf("%X\n", '你')              // 4F60
	fmt.Printf("%d\n", 'A')              // 65
	fmt.Printf("%t\n", 1 > 2)            // false
	fmt.Printf("%e\n", 4396.7777777)     // 4.396778e+03 默认精度6位
	fmt.Printf("%20.3e\n", 4396.7777777) //            4.397e+03 设置宽度20,精度3,宽度一般用于对齐
	fmt.Printf("%E\n", 4396.7777777)     // 4.396778E+03
	fmt.Printf("%f\n", 4396.7777777)     // 4396.777778
	fmt.Printf("%o\n", 16)               // 20
	fmt.Printf("%p\n", []int{1})         // 0xc000016110
	fmt.Printf("Hello %s\n", "World")    // Hello World
	fmt.Printf("Hello %q\n", "World")    // Hello "World"
	fmt.Printf("%T\n", 3.0)              // float64
	fmt.Printf("%v\n", user)             // {let-it-go 99}
	fmt.Printf("%+v\n", user)            // {Name:let-it-go Age:99}
	fmt.Printf("%#v\n", user)            // main.User{Name:"let-it-go", Age:99}
}
