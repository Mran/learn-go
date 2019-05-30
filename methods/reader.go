package main

import (
	"fmt"
	"io"
	"strings"
)
func main() {
	r:=strings.NewReader("Hello,Reader!")
	b:=make([]byte,8)
	for{
		//有两个返回值,一个是读取到数量,另一个是返回的错误
		//参数是读取缓存
		n,err:=r.Read(b)
		fmt.Printf("n = %v er=%v b=%v\n",n,err,b)
		fmt.Printf("b[:n]=%q\n",b[:n])
		if err==io.EOF{
			break
		}
	}
}