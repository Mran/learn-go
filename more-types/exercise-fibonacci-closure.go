package main
import "fmt"
//用闭包实现一个斐波那契数列
func flbonacci()func()int{
	i:=0
	sum:=0
	return func ()int  {
		temp:=sum
		sum=i+sum
		i=temp
		if sum==0{
			sum=1
		}
		return sum
	}
}
func main() {
	f:=flbonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}