package main

import (
	"fmt"
	"math"
)

//只是数据类型
type Vertex struct {
	X, Y float64
}

//类似于kt里的扩展函数,给这个数据类型加上一个扩展函数
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
