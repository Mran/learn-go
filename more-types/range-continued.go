package main
import "fmt"
func main() {
	pow:=make([]int,10)
	//只返回index
	for i:=range pow{
		pow[i]=1<<uint(i)
	}
	//忽略index,只取value
	for _,value:= range pow{
		fmt.Printf("%d\n",value)
	}
}