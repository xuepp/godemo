package types

import "fmt"

func ArrayDemo() {
	fmt.Println("------------以下是Array相关demo------------------")
	demo1()
	demo2()
	demo3()

}

func demo1() {
	fmt.Print("1.数组声明")
	//使用var关键字，创建一个字符串类型的数组
	var myarr [3]string
	myarr[0] = "qqq"
	myarr[1] = "666"
	myarr[2] = "cccc"

	fmt.Println("数组的元素:")
	fmt.Println("元素 1: ", myarr[0])
	fmt.Println("元素 2: ", myarr[1])
	fmt.Println("元素 3: ", myarr[2])

	arr := [3]string{"zhangsan", "177cm", "64kg"}
	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])

	}

}

func demo2() {

	fmt.Println("2.数组复制")
	//创建和初始化数组
	//使用简写声明
	my_arr1 := [5]string{"Scala", "Perl", "Java", "Python", "Go"}

	//将数组复制到新变量中
	//在这里，元素按值传递
	//my_arr2 := my_arr1
	my_arr2 := &my_arr1

	fmt.Println("Array_1: ", my_arr1)
	fmt.Println("Array_2:", my_arr2)

	my_arr1[0] = "C++"

	//在这里，当我们复制数组时
	//按值放入另一个数组
	//然后对原始内容进行更改
	//数组不反映在
	//该数组的副本
	fmt.Println("\nArray_1: ", my_arr1)
	fmt.Println("Array_2: ", my_arr2)

}

func demo3() {
	fmt.Print("3.数组指针")
	//创建和初始化数组
	//使用简写声明
	new_arr1 := [6]int{12, 456, 67, 65, 34, 34}

	//将数组复制到新变量中
	//在这里，元素通过引用传递
	new_arr2 := &new_arr1

	fmt.Println("Array_1: ", new_arr1)
	fmt.Println("Array_2:", *new_arr2)

	new_arr1[5] = 1000000

	//在这里，当我们复制数组时
	//通过引用放入另一个数组
	//然后对原始内容进行更改
	//数组将反映在
	//该数组的副本
	fmt.Println("\nArray_1: ", new_arr1)
	fmt.Println("Array_2:", *new_arr2)

}
