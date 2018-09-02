package main

import "fmt"

func main() {
	a := 10
	//Println一段一段输出，自动加换行
	fmt.Println("a=", a)

	//Printf 格式化输出，将a放入%d位置
	fmt.Printf("a = %d\n", a)

	//Printf的优势, 如果用Println，反而不简便
	b, c := 20, 30
	fmt.Printf("a = %d, b = %d, c = %d", a, b, c)
	fmt.Println("a = ", a, "b =", b, "c =", c)

}
