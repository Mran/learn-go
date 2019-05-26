package main
import "fmt"
type Vertex struct {
	Lat,Long float64
}
//map[键类型]值类型
var m=map[string]Vertex{
	"Bell Labs":Vertex{
		40,123,
	},
	"Google":Vertex{
		37,-123,
	},
}
func main() {
	fmt.Println(m)
}