package main

import (
	"fmt"
)

func main() {
	var a int
	a = 10
	fmt.Println("a=", a)

	var b int = 12
	fmt.Println("b=", b)

	c := 15
	fmt.Println("c=", c)

	//	变量的交换
	i, j := 10, 20
	fmt.Println(i, j)
	j, i = i, j
	fmt.Println(i, j)

	//	_为匿名变量
	var tmp int
	tmp, _ = i, j
	fmt.Println(tmp)
}
