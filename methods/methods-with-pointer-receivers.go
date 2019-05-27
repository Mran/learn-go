package main
import (
	"fmt"
	"math"
)
type Vertex struct {
	X,Y float64
}
//推荐通过指针来实现扩展函数来获得避免值传递,获得更好的性能
//同时也可以对原对象进行修改
func (v *Vertex)Scale(f float64)  {
	fmt.Println("Scale",v)
	v.X=v.X*f
	v.Y=v.Y*f
}
func (v *Vertex)Abs()float64  {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}
func (v Vertex)Scale1(f float64)  {
	fmt.Println("Scale1",&v)
	v.X=v.X*f
	v.Y=v.Y*f
}
func (v Vertex)Abs2()float64  {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func main() {
v:=Vertex{3,4}
fmt.Println(&v)
fmt.Println("Before scaling: %+v,Abs:%v\n",v,v.Abs())
v.Scale(5)	
v.Scale1(10)

fmt.Println("Before scaling: %+v,Abs:%v\n",v,v.Abs())

}