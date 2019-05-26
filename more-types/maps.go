package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}
//声明
var m map[string]Vertex
func main() {
	//初始化
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -7439967,
	}
	fmt.Println(m["Bell Labs"])
}
