package main
import "fmt"
type Vertex struct {
	X int
	Y int
}
//通过点操作符获取结构体里的内容

func main() {
	v:=Vertex{1,2}
	v.X=4
	fmt.Println(v.X)
}