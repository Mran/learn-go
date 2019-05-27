package main
import (
	"fmt"
)
func describe(i interface{}){
	fmt.Println("(%v,%T)\n",i,i)
}
//空接口可以保存任何类型的值。 （每种类型至少实现零方法。）
//处理未知类型值的代码使用空接口。 
//例如，fmt.Print接受任意数量的interface {}类型的参数。

func main() {
	var i interface{}
	describe(i)
	i=42

	describe(i)

	i="hello"

	describe(i)
}