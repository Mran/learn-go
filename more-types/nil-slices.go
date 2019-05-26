package main
import "fmt"
func main() {
	//空数组
	var s []int
	//长度为0,容量为0
	fmt.Println(s,len(s),cap(s))
	//实例是nil
	if s==nil{
		fmt.Println("nil!")
	}
}