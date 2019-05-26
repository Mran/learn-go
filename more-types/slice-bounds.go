package main
import "fmt"
//类似于python中的切片操作,忽略起始则从0开始,忽略结束则从最后结束
func main() {
	s:=[]int{2,3,4,5,6,11}
	s=s[1:4]
	fmt.Println(s)
	s=s[:2]
	fmt.Println(s)
	s=s[1:]
	fmt.Println(s)

}