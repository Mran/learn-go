package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) String() string {
	return fmt.Sprintf("a1111111t %v %s", e.When, e.What)
}

//实现错误类型的接口
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}
func main() {
	//run()返回了error类型的数据,那么就要实现Error()的接口
	if err := run(); err != nil {
		fmt.Println(err)
	}

}
