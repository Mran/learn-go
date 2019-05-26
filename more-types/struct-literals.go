package main
import "fmt"
type Vertex struct {
	X,Y int
}
var(
	//指定初始化
	v1=Vertex{1,2}
	//缺省初始化
	v2=Vertex{X:1}
	v3=Vertex{}
	//初始化并获得指针
	p=&Vertex{1,2}
)
func main() {
	fmt.Println(v1,p,v2,v3)
	p.Y=100
	fmt.Println(v1,p,v2,v3)
}