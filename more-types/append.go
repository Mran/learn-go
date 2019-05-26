package main
import "fmt"
func printfSlice(s []int)  {
	fmt.Printf("len=%d cap=%d %v \n",len(s),cap(s),s)
}
func main() {
	var s []int
	printfSlice(s)
	fmt.Println(&s)
	
	s=append(s,0)
	printfSlice(s)
	fmt.Println(&s)

	s=append(s,2,3,4)
	printfSlice(s)
	fmt.Printf("%v",&s)

}