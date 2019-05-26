package main
import (
	"fmt"
)
var pow =[]int{1,2,4,8,16,32,64,128}
func main() {
	//range 用法类似于python range的用法,会有两个返回值
	for  i,v:=range pow{
		fmt.Printf("2**%d=%d\n",i,v)
	}
}