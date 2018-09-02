package main

import "fmt"

func main() {
	//	iota为常量自动生成器，用于给常量赋值使用，没隔一行会自动累加1
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)

	//	iota如果遇到const，则重置恢复为0
	const d = iota
	fmt.Printf("d=%d\n", d)

	//	可以只写一个iota
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Printf("a1=%d, b1=%d, c1=%d\n", a1, b1, c1)

	//	如果是同一行，则值都一样
	const (
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)
	fmt.Printf("i=%d, j1=%d, j2=%d, j3= %d, k=%d\n", i, j1, j2, j3, k)
}
