package types

import "fmt"

func StructDemo() {
	fmt.Println("------------以下是Struct相关demo------------------")
	fmt.Println("1.Struct的定义和使用")
	p := Person{Name: "Alice", Age: 30}
	fmt.Println("Person Name:", p.Name)
	fmt.Println("Person Age:", p.Age)

	var p2 Person
	fmt.Println("Default Person Name:", p2)

	fmt.Println("\n2.指向结构的指针")
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)

	// emp8.firstName用于访问
	//字段firstName
	fmt.Println("First Name: ", emp8.firstName)
	fmt.Println("Age: ", emp8.age)

	fmt.Println("\n3.嵌套结构体")

	//初始化结构字段
	result := Teacher{
		name:    "Suman",
		subject: "Java",
		exp:     5,
		details: Student{"Bongo", "CSE", 2},
	}
	fmt.Println("result", result)

	fmt.Println("老师详细情况")
	fmt.Println("老师的名字: ", result.name)
	fmt.Println("学科: ", result.subject)
	fmt.Println("经历: ", result.exp)

	fmt.Println("\n学生详细资料")
	fmt.Println("学生的名字: ", result.details.name)
	fmt.Println("学生的部门名称: ", result.details.branch)
	fmt.Println("年龄: ", result.details.year)

	fmt.Println("\n4.创建匿名结构体")
	demo := struct {
		name string
		age  int
	}{
		name: "Bob",
		age:  25,
	}
	fmt.Println("匿名结构体:", demo)

	fmt.Println("\n5.匿名字段")

	value := student{123, "Bud", 8900.23}

	fmt.Println("入学人数 : ", value.int)
	fmt.Println("学生姓名 : ", value.string)
	fmt.Println("套餐价格 : ", value.float64)

	fmt.Print("\n6.函数用作结构体字段\n")

	// 初始化字段结构
	result1 := Author{
		name:      "Sonia",
		language:  "Java",
		Marticles: 120,
		Pay:       500,
		salary: func(Ma int, pay int) int {
			return Ma * pay
		},
	}

	fmt.Println("作者姓名: ", result1.name)
	fmt.Println("语言: ", result1.language)
	fmt.Println("五月份发表的文章总数: ", result1.Marticles)
	fmt.Println("每篇报酬: ", result1.Pay)
	fmt.Println("总工资: ", result1.salary(result1.Marticles, result1.Pay))

	fmt.Println("\n7.匿名函数用作结构体字段")
	//初始化结构字段
	result3 := Author2{
		name:      "Sonia",
		language:  "Java",
		Tarticles: 340,
		Particles: 259,
		Pending: func(Ta int, Pa int) int {
			return Ta - Pa
		},
	}

	fmt.Println("作者姓名: ", result3.name)
	fmt.Println("语言: ", result3.language)
	fmt.Println("文章总数: ", result3.Tarticles)

	fmt.Println("发表文章总数: ", result3.Particles)
	fmt.Println("待处理文章: ", result3.Pending(result3.Tarticles, result3.Particles))

}

type Person struct {
	Name string
	Age  int
}

// 定义一个结构
type Employee struct {
	firstName, lastName string
	age, salary         int
}

// 创建结构
type Student struct {
	name   string
	branch string
	year   int
}

// 创建嵌套结构
type Teacher struct {
	name    string
	subject string
	exp     int
	details Student
}

// 创建一个结构匿名字段
type student struct {
	int
	string
	float64
}

type Finalsalary func(int, int) int

// 创建结构
type Author struct {
	name      string
	language  string
	Marticles int
	Pay       int

	//函数作为字段
	salary Finalsalary
}

// 创建结构
type Author2 struct {
	name      string
	language  string
	Tarticles int
	Particles int
	Pending   func(int, int) int
}
