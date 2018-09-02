package main

import "fmt"

func main() {
	var (
		a int
		b float64
	)
	a, b = 10, 3.14
	fmt.Println(a, b)

	const (
		i int     = 10
		j float64 = 3.14
	)
	fmt.Println(i, j)
}
