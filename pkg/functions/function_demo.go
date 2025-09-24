package functions

import (
	"fmt"
	"strings"
)

func init() {
	fmt.Println("------------以下是函数的相关demo------------------")
}

func FunctionDemo() {
	//传int类型参
	fmt.Println(add(5, 4))
	//fmt.Println(functions.FunctionDemo.add(5, 4))
	//定义了函数返回int类型
	fmt.Println(test())
	//没有定义函数返回int类型
	test1()
	//延时调用hanshu函数,最后执行这个函数
	defer yanchi()
	//多个参数
	fmt.Println(joinstr("GEEK", "GFG"))
	fmt.Println(joinstr("Geeks", "for", "Geeks"))
	fmt.Println(joinstr("G", "E", "E", "k", "S"))
	//在可变函数中传递一个切片
	element := []string{"geeks", "FOR", "geeks"}
	fmt.Println("传递一个切片:", joinstr(element...))

	//调用全局匿名函数
	value()

	//在匿名函数传递参数
	func(ele string) {
		fmt.Println(ele)
	}("cainiaojc")

	//调用全局匿名函数，往函数里传参
	value2("cainiaojc")

	//局部定义一个匿名函数并赋值个aa
	aa := func(p, q string) string {
		return p + q + "Geeks"
	}
	//往函数里传递匿名函数
	GFG(aa)

	var Zz int = 10
	fmt.Printf("函数调用前，Zz的值为 = %d", Zz)

	//通过引用调用传递变量Z地址
	modifydata(&Zz)

	fmt.Printf("\n函数调用后，Zz的值为 = %d", Zz)

	var f int = 700
	var s int = 900

	fmt.Printf("函数调用前的值\n")
	fmt.Printf("f = %d 和 s = %d\n", f, s)

	//通过引用调用
	//传递变量地址
	swap(&f, &s)

	fmt.Printf("\n函数调用后的值\n")
	fmt.Printf("f = %d 和 s = %d", f, s)

	//将返回值分配到，三个不同的变量
	var myvar1, myvar2, myvar3 = myfunc(4, 2)

	// 显示值
	fmt.Printf("结果为: %d", myvar1)
	fmt.Printf("\n结果为: %d", myvar2)
	fmt.Printf("\n结果为: %d", myvar3)

	//将返回值分配到
	//两个不同的变量
	//var area1, area2 = myfunc2(2, 4)
	//_忽略值,go语言定义的变量没有使用会报错
	var area1, _ = myfunc2(2, 4)

	fmt.Printf("矩形面积为: %d", area1)
	//fmt.Printf("\n正方形面积为: %d", area2)

	//构造函数
	user := NewUser()
	fmt.Println("User Name: ", user)
	greeter := NewGreeter(user)
	greeter.Greet()
}

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

func init() {
	fmt.Println("Welcome to init() function")
}

// User 是一个简单的业务对象
type User struct {
	Name string
}

// NewUser 是 User 的构造函数
func NewUser() *User {
	return &User{Name: "Tom"}
}

// Greeter 依赖 User
type Greeter struct {
	User *User
}

// NewGreeter 是 Greeter 的构造函数
func NewGreeter(u *User) *Greeter {
	return &Greeter{User: u}
}

// Greet 打印问候
func (g *Greeter) Greet() {
	fmt.Printf("Hello, %s!\n", g.User.Name)
}
