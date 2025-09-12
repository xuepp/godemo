package types

import "fmt"

func InterfaceDemo() {
	fmt.Println("------------以下是Interface相关demo------------------")
	fmt.Println("\n1.如何创建接口")
	// 访问使用桶的接口
	t := myvalue{10, 14}
	fmt.Println("桶的面积 :", t.Tarea())
	fmt.Println("桶的容量:", t.Volume())

	fmt.Println("\n2.创建一个空接口")
	var t2 tank2
	fmt.Println("tank interface值为: ", t2)
	fmt.Printf("tank 的类型是: %T ", t2)

	fmt.Println("\n3.不允许在两个或多个接口中创建相同的名称方法")

	//结构体赋值
	values := author{
		a_name:    "Mickey",
		branch:    "Computer science",
		college:   "XYZ",
		year:      2012,
		salary:    50000,
		particles: 209,
		tarticles: 309,
	}

	// 访问使用接口1的方法
	var i1 AuthorDetails = values
	i1.details()

	//访问使用接口2的方法
	var i2 AuthorArticles = values
	i2.articles()

	fmt.Println("\n4.接口嵌套")

	// 结构体赋值
	value2s := author{
		a_name:    "Mickey",
		branch:    "Computer science",
		college:   "XYZ",
		year:      2012,
		salary:    50000,
		particles: 209,
		tarticles: 309,
	}

	// 使用FinalDetails接口访问接口1,2的方法
	var f FinalDetails = value2s
	f.details()
	f.articles()

}

type tank interface {

	// 方法
	Tarea() float64
	Volume() float64
}

type myvalue struct {
	radius float64
	height float64
}

// 实现方法
// 桶的（Tank）接口
func (m myvalue) Tarea() float64 {

	return 2*m.radius*m.height + 2*3.14*m.radius*m.radius
}

func (m myvalue) Volume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}

// 创建接口
type tank2 interface {

	// 方法
	Tarea2() float64
	Volume2() float64
}

// 接口 1
type AuthorDetails interface {
	details()
}

// 接口 2
type AuthorArticles interface {
	articles()
}

// 结构体
type author struct {
	a_name    string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
}

// 实现接口方法1
func (a author) details() {

	fmt.Printf("作者: %s", a.a_name)
	fmt.Printf("\n部分: %s 通过日期: %d", a.branch, a.year)
	fmt.Printf("\n学校名称: %s", a.college)
	fmt.Printf("\n薪水: %d", a.salary)
	fmt.Printf("\n出版文章数: %d", a.particles)

}

// 实现接口方法 2
func (a author) articles() {

	pendingarticles := a.tarticles - a.particles
	fmt.Printf("\n待定文章: %d", pendingarticles)
}

// 接口3嵌套接口1和接口2
type FinalDetails interface {
	AuthorDetails
	AuthorArticles
}

// 结构体
type author2 struct {
	a_name    string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
}
