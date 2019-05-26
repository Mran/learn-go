package main
import "fmt"
type Vertex struct {
	X int
	Y int
}
type Mynew struct{
	Name string
}
func main() {
	fmt.Println(Vertex{1,2})
	fmt.Println(Mynew{"123"})
}