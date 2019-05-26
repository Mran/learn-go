package main
import "fmt"
func split(sum int)(x,y,c int)  {
	x=sum*4/9
	y=sum-x
	c=100
	return
}
func main()  {
	fmt.Println(split(17))
}