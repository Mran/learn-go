package main
import "fmt"
func main() {
	m:=make(map[string]int)

	m["Answer"]=42
	fmt.Println("The value:",m["Answer"])

	m["Answer"]=48
	fmt.Println("The Value:",m["Answer"])

	delete(m,"Answer")
	fmt.Println("The value:",m["Answer"])

	//如果answer在m内则v为相应的value,ok为true.
	//否则v为相应类型的zero值,ok为false
	v,ok:=m["Answer"]
	fmt.Println("The value:",v,"present?",ok)
}