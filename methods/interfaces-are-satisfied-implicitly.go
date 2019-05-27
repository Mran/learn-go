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
//T通过函数实现了I的接口描述,而且不用其他的声明
func (t T)M()  {
	fmt.Println(t.S)
}
func main() {
	var i I=T{"hello"}
	i.M()
}