package main

import (
	"fmt"
	"strings"
)

// add:函数名
// Parameter_list：函数参数和类型
// Return_type：这是可选的，它包含函数返回的值的类型。如果在函数中使用return_type，则必须在函数中使用return语句。
//
//func:go语言关键字
func add(a int, b int) int {
	return a + b
}

func test() int {
	return 66
}

func test1() {
	fmt.Println(77)
}

// 可变参数函数联接字符串
func joinstr(values ...string) string {
	return strings.Join(values, "-")
}

// 定义一个全局匿名函数
var value = func() {
	fmt.Println("Welcome! to (aaaaaa)")
}

var value2 = func(ele string) {
	fmt.Println(ele)
}

// 匿名函数作为参数传递给函数
func GFG(i func(p, q string) string) {
	fmt.Println(i("Geeks", "for"))

}

// 修改值的函数
func modifydata(Z *int) {
	*Z = 70
}

// 交换值的函数，将指针指向整数
func swap(x, y *int) int {

	//临时存储变量
	tmp := *x
	*x = *y
	*y = tmp

	return tmp
}

// myfunc返回3个int类型的值
func myfunc(p, q int) (int, int, int) {
	return p - q, p * q, p + q
}

// myfunc2返回2个int类型的值
// 这里是返回值名称
// 是rectangle and square
func myfunc2(p, q int) (rectangle int, square int) {
	rectangle = p * q
	square = p * p
	return
}

func yanchi() {
	fmt.Println("------------延迟调用函数实例------------------")
}
