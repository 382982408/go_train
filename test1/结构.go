// 程序所属包
package test1
//导入依赖包
import "fmt"
//常量定义
const Name string = "imook"
//全局变量的定义与赋值
var a string = "慕课"
//一般类型声明
type imookInt int
//结构的声明
type Learn struct {
}
// 声明接口
func Ilearn interface () {
}
//函数定义
func learnImook() {
	fmt.Print("learnImook")
}
//main()函数
func main() {
	learnImook()
	fmt.Print("hello world")
}