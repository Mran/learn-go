package main
import (
	"fmt"
	"math"
)
type Vertex struct {
	X,Y float64
}
func (v Vertex)Abs() float64 {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}
//扩展函数可以使用指针对象,可以对原对象进行处理
func (v *Vertex)Scale(f float64)  {
	v.X=v.X*f
	v.Y=v.Y*f
}
//这样子并不能修改原对象,只能使用指针
func (v Vertex)Change()  {
	v.X=100
}
func main() {
	v:= Vertex{3,4}
	v.Change()
	// v.Scale(10)
	fmt.Println(v.Abs())
}