package main

import (
	"fmt"
)

type ErrnegativeSqrt float64

func (e ErrnegativeSqrt) Error() string {
	//进行了一个类型转换
	return fmt.Sprintf("cannot Sqrt negative number %v", float64(e))
}
//返回多个参数
func Sqrt(x float64) (float64, error) {
	if x <= 0 {
		return -2, ErrnegativeSqrt(x)
	}
	return 0, nil

}
func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
