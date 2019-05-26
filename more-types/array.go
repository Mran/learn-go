package main
import "fmt"
func main() {
	//数组的长度在初始化后就不可改变
	var a [2]string
	a[0]="Hello"
	a[1]="World"
	fmt.Println(a[0],a[1])
	fmt.Println(a)
	//声明时初始化,
	primes:=[6]int{2,3,4,5}
	fmt.Println(primes)
}