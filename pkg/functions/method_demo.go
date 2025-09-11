package functions

import "fmt"

func init() {
	fmt.Println("------------以下是MethodDemo方法的相关demo------------------")
}

func MethodDemo() {
	fmt.Println("=====结构类型接收器的方法")
	//初始化值
	//Author结构体

	res := author{
		name:      "Sona",
		branch:    "CSE",
		particles: 203,
		salary:    34000,
	}

	//调用方法
	res.show()
	fmt.Println("1.非结构类型接收器的方法")
	value1 := data(23)
	value2 := data(20)
	res2 := value1.multiply(value2)
	fmt.Println("最终结果: ", res2)
	fmt.Println("2.带指针接收器的Go方法")

	//初始化author结构体
	res3 := author2{
		name:   "Sona",
		branch: "CSE",
	}

	fmt.Println("Author's name: ", res3.name)
	fmt.Println("Branch Name(Before): ", res3.branch)

	//创建一个指针
	p := &res3

	//调用show方法
	p.show2("ECE")
	fmt.Println("Author's name: ", res3.name)
	fmt.Println("Branch Name(After): ", res3.branch)

	fmt.Println("3.Go同名方法,必须具有不同的类型接收器")

	// 初始化结构体的值
	val1 := student{"Rohit", "EEE"}

	val2 := teacher{"Java", 50}

	//调用方法
	val1.show()
	val2.show()

}

// Author 结构体
type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

// 接收者的方法
func (a author) show() {
	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}

// 类型定义
type data int

// 定义一个方法
// 非结构类型的接收器
func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

/*
//如果您尝试运行此代码，

//然后编译器将抛出错误
func(d1 int)multiply(d2 int)int{
return d1 * d2
}
*/

// Author 结构体
type author2 struct {
	name   string
	branch string
}

// 方法，使用author类型的接收者
func (a *author2) show2(abranch string) {
	(*a).branch = abranch
}

// 创建结构体
type student struct {
	name   string
	branch string
}

type teacher struct {
	language string
	marks    int
}

// 名称相同的方法，但有不同类型的接收器
func (s student) show() {

	fmt.Println("学生姓名:", s.name)
	fmt.Println("Branch: ", s.branch)
}

func (t teacher) show() {

	fmt.Println("Language:", t.language)
	fmt.Println("Student Marks: ", t.marks)
}
