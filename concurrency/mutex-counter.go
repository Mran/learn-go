package main

import (
	"fmt"
	"time"
	"sync"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	//锁住
	c.mux.Lock()
	c.v[key]++
	//解锁
	c.mux.Unlock()
}
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	//使用defer操作避免忘记解锁
	defer c.mux.Unlock()
	return c.v[key]
}
func main() {
	c:=SafeCounter{v:make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
