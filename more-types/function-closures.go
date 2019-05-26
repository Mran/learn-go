package main

import "fmt"

//定义的这个函数是把一个函数当作返回值返回了,在这个函数里
//使用了外层的函数值,并且保留了下来
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

}
