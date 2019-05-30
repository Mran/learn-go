package main
import (
	"fmt"
)
func finbonacci(n int,c chan int)  {
	x,y:=0,1
	for i := 0; i < n; i++ {
		//向管道中压入c的值
		c<-x
		x,y=y,x+y
	}
	//接受者可以关闭管道,如果管道没有满,可以手动关闭管道以避免死锁
	close(c)
}
func main() {
	//制造一个10个缓冲区大小的管道
	c:=make(chan int,10)
	//开启线程
	go finbonacci(cap(c),c)
	//虽然开了线程,但是被阻塞,因为c的值没有满
	for i:=range c{
		fmt.Println(i)
	}
}