package main
import "fmt"
type Vertex struct {
	Lat,Long float64
}
//忽略了单个初始化时的类型名
var m=map[string]Vertex{
	"Bell Labs":{40,-74,},
	"Google":{37,-122},
}
func main() {
	fmt.Println(m)
}