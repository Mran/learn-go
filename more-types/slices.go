package main

import (
	"fmt"
)
func main() {
	primes:=[6]int{1,2,3,4,5,6}
	//类似于python中的切片操作
	var s =primes[1:4]
	fmt.Println(s)
}