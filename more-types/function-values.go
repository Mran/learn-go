package main

import (
	"fmt"
	"math"
)
//把函数当作参数来传入
func compute(fn func(float64,float64)float64)float64  {
	return fn(3,4)
}
func main() {
	
	hypot:=func (x,y float64)float64  {
		return math.Sqrt(x*x+y*y)
	}

	fmt.Println(hypot(5,12))
	//把函数当成参数时,参数类型要一一对应
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

}