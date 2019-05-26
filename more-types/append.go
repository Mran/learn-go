package main
import "fmt"
func printfSlice(s []int)  {
	fmt.Printf("len=%d cap=%d %v \n",len(s),cap(s),s)
}
func main() {
	var s []int
	printfSlice(s)
	
	s=append(s,0)
	printfSlice(s)
}