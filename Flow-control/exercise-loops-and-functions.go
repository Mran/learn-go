package main

import "fmt"

func Sqrt(x float64) float64 {
	z := x/2
	count:=1
	for ;z >= 1;count++{
		z -= (z*z - x) / (2 * z)
		if count>10{
			break
		}
		fmt.Println(z)
	}
	return z
}
func main() {
	fmt.Println(Sqrt(9))
}
