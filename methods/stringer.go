package main
import (
	"fmt"
)
type Person struct {
	Name string
	Age int
}
//实现了String()接口
//类似于重写了java的toString
func (p Person)String()string  {
	return fmt.Sprintf("%v (%v years)",p.Name,p.Age)
}
func main() {
	a:=Person{"Arthur Dent",42}
	z:=Person{"Zaphod Beeblerox",9001}
	fmt.Println(a,z)
}