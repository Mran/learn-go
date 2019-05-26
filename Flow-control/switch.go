package main

import (
	"fmt"
	"runtime"
)
//switch的条件可以是变量,只匹配一个情况,不用break
func main()  {
	fmt.Print("Go runs on ")
	switch os:=runtime.GOOS;os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n",os)
	}
}