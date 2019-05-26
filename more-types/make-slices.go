package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
func main() {
	//创建一个容量和长度都是为5的数组,返回的是引用
	a := make([]int, 5)
	printSlice("a", a)
	//创建一个容量为5,长度为0的数组,返回0的引用
	b := make([]int, 0, 5)
	printSlice("b", b)
	//对于一个无长度但是容量的数组的切片操作,实际上是创造了新的数组
	c := a[:2]
	printSlice("c", c)

	d := c[2:5]
	d[0]=1
	c[0]=4
	printSlice("d", d)
	fmt.Println(d)
	fmt.Println(c)
	fmt.Println(b)
	printSlice("a", a)

}
