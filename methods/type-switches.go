package main
import "fmt"
//使用i.(type) 来获取该泛型是那种泛型,自定义结构类型的也是可以的
type I struct {
	s string
}
func do(i interface{})  {
	switch v:=i.(type){
	case int:
		fmt.Printf("Twice %v is %v\n",v,v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n",v,len(v))
	case I:
		fmt.Println(v)
	default:
		fmt.Printf("I dont know about type %T\n",v)
	}
}
func main() {
	do(12)
	do("hello")
	do(true)
	i:=I{"123"}
	fmt.Println("%v",i)
	do(i)
}