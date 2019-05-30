package main

import (
	"fmt"
	"time"
)
//看不懂到底是个什么意思
//我知道如果不加sleep,就会导致代码走完了,但是有一个线程还没有走完
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	go say("world")
	say("hello")
}
