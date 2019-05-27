package main

import (
	"fmt"
	"math"
)

//非stuuct类型也可以加入扩展函数
//相当于通过继承形成了一个新的类型
type MyFloat float64
type MyInt int

type Parent struct {
	P int
}

func (p Parent) PM() string {
	return "Parent"
}
type Child Parent
func (c Child)CM()string  {
	return "Child"
}
func (i MyInt) Abs() int {
	if i < 0 {
		return int(-i)
	}
	return int(i)
}
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func main() {
	f := MyFloat(-math.Sqrt(2))
	fmt.Println(f.Abs())
	i := MyInt(-10)
	fmt.Println(i.Abs())
	p:=Parent{}
	c:=Child{P:1}
	fmt.Println(p.PM())
	fmt.Println(c.P)
}
