package main

import (
	"fmt"
)

type I interface {
	M()
}
type T struct {
	S string
}

func (t *T) M() {
	if  t==nil{
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	//接口为空值也能实现
	var t *T
	i=t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	//接口能获得实例里的值
	i.M()
}
func describe(i I) {
	fmt.Println("(%v, %T)\n", i, i)
}
