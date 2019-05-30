package main

import (
	"fmt"
)

func main() {

	//构造了一个缓冲的管道,缓冲区为2个int,缓冲区满之后才回继续运行,否则会死锁,如果满了还是继续添加,也会死锁
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
