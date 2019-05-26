package main
import (
	"fmt"
)
func main() {
	names:=[4]string{"1","2","3","dfg"}
	fmt.Println(names)
	//操作的并不是原数组的复制,而是原数组本身
	a:=names[0:2]
	b:=names[1:3]
	fmt.Println(a,b)
	b[0]="xxx"
	fmt.Println(a,b)
	fmt.Println(names)
}