// Author: Tong Zicheng
// Date: 2020/6/26 21:16
// Describe: 菲波那切数列 （https://studygolang.com/articles/7473）

package basics

import (
	"fmt"
	"time"
)

// 递归实现
func fibonacci(num int) int {
	if num < 2 {
		return 1
	}

	return fibonacci(num-1) + fibonacci(num-2)
}

func main() {
	for i := 0; i < 10; i++ {
		nums := fibonacci(i)
		fmt.Println(nums)
	}
}

// 必包实现
func fibonacci2() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main2() {
	f := fibonacci2()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// 计算Go运行的时间
func StartCac() {
	t1 := time.Now() // get current time
	//logic handlers
	for i := 0; i < 1000; i++ {
		fmt.Print("*")
	}
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}
