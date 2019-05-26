package main
import (
	"fmt"
)
//len()的作用是看分片的长度,cap的作用是看原分片的长度
func printSlice(s []int)  {
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
}
func main() {
	s:=[]int{2,3,5,7,11,13}
	printSlice(s[:0])

	s=s[:4]
	printSlice(s)
	s=s[2:]
	printSlice(s)
	//编译不通过,数组初始化只能使用常量值
	size :=4
	a:=[size]int

	fmt.Println(a)
}