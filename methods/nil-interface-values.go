package main

import "fmt"

type I interface {
	M()
}

func describe(i I) {
	//%v获得值,%T获得类型
	fmt.Println("(%v,%T)", i, i)
}
func main() {
	var i I
	//空接口不能被调用
	describe(i)
	i.M()
}
