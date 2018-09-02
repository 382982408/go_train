package main

import "fmt"

func main() {
	const a int = 10
	//	如果修改a的值会报错
	//	a = 20

	//	自动推导类型,不能使用:=
	const b = 10.2
	fmt.Println("b =", b)
	fmt.Printf("b的类型是：%T", b)
}
