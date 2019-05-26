package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}
//使用指针指向原对象,否则建立新的对象
func main() {
	v := Vertex{1, 2}
	p := v
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(p)
}
