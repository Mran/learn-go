package main

import "fmt"
//类型断言,可以用这种方法来判断是否是某种类型
func main() {
	var i interface{}="hello"

	s:=i.(string)
	fmt.Println(s)

	s,ok:=i.(string)
	fmt.Println(s,ok)

	f,ok:=i.(float64)
	fmt.Println(f,ok)

	f=i.(float64)
	fmt.Println(f)
}