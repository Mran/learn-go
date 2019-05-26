package main

import "fmt"

//deferd的作用就是可以在函数返回之后再执行你定义的代码
func c() (i int) {
    defer func() { i++ }()
    return 1
}
func main() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	} 
	fmt.Println("done")
	fmt.Println(c())
}
