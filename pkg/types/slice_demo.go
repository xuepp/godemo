package types

import (
	"bytes"
	"fmt"
	"sort"
)

func SliceDemo() {
	fmt.Println("------------以下是slice相关demo------------------")
	fmt.Println("\n1.切片的组成")
	Demo1()
	fmt.Println("\n2.创建和初始化切片")
	Demo2()
	fmt.Println("\n3.切片引用数组")
	Demo3()
	fmt.Println("\n4.从给定的切片创建切片")
	Demo4()
	fmt.Println("\n5.函数传递切片")
	//函数修改切片，原切片也被修改
	Demo5()
	//函数修改切片，原切片不变
	Demo6()
	fmt.Println("\n6.比较切片大小")
	Demo7()
	fmt.Println("\n7.切片排序")
	Demo8()
	Demo9()
	fmt.Println("\n8.切片分割")
	Demo10()
	Demo11()
}

func Demo1() {

	//创建一个数组
	arr := [7]string{"这", "是", "Golang", "基础", "教程", "在线", "www.cainiaojc.com"}

	//显示数组
	fmt.Println("数组:", arr)

	//创建切片
	myslice := arr[1:6]

	//显示切片
	fmt.Println("切片:", myslice)

	//显示切片的长度
	fmt.Printf("切片长度: %d", len(myslice))

	//显示切片的容量
	fmt.Printf("\n切片容量: %d", cap(myslice))
}

func Demo2() {

	//使用var关键字,创建切片
	var my_slice_1 = []string{"cainiaojcs", "for", "cainiaojcs"}

	fmt.Println("My Slice 1:", my_slice_1)

	//创建切片
	//使用简写声明
	my_slice_2 := []int{12, 45, 67, 56, 43, 34, 45}
	fmt.Println("My Slice 2:", my_slice_2)
}

func Demo3() {
	//创建一个数组
	arr := [4]string{"cainiaojcs", "for", "cainiaojcs", "GFG"}

	//从给定数组创建切片
	//从切片索引1开始到索引2结束
	var my_slice_1 = arr[1:2]
	my_slice_2 := arr[0:]
	my_slice_3 := arr[:2]
	my_slice_4 := arr[:]

	//显示结果
	fmt.Println("我的数组: ", arr)
	fmt.Println("My Slice 1: ", my_slice_1)
	fmt.Println("My Slice 2: ", my_slice_2)
	fmt.Println("My Slice 3: ", my_slice_3)
	fmt.Println("My Slice 4: ", my_slice_4)
}

func Demo4() {
	//创建切片
	oRignAl_slice := []int{90, 60, 40, 50, 34, 49, 30}

	//从给定的切片创建切片
	var my_slice_1 = oRignAl_slice[1:5]
	my_slice_2 := oRignAl_slice[0:]
	my_slice_3 := oRignAl_slice[:6]
	my_slice_4 := oRignAl_slice[:]
	my_slice_5 := my_slice_3[2:4]

	//显示结果
	fmt.Println("原始切片:", oRignAl_slice)
	fmt.Println("新切片 1:", my_slice_1)
	fmt.Println("新切片 2:", my_slice_2)
	fmt.Println("新切片 3:", my_slice_3)
	fmt.Println("新切片 4:", my_slice_4)
	fmt.Println("新切片 5:", my_slice_5)
}

func Demo5() {
	myfun := func(ys []string) {
		ys[2] = "python"
		fmt.Println("修改切片:", ys)
	}
	myslice := []string{"C", "C++", "Java", "go", "php"}

	fmt.Println("原切片:", myslice)
	myfun(myslice)
	fmt.Println("最终切片:", myslice)
}

func Demo6() {
	myfun := func(ys []string) {
		ys = append(ys, "python")
		fmt.Println("修改切片:", ys)
	}
	myslice := []string{"C", "C++", "Java", "go", "php"}

	fmt.Println("原切片:", myslice)
	myfun(myslice)
	fmt.Println("最终切片:", myslice)

}

func Demo7() {
	//使用简写声明创建和初始化字节片

	slice_1 := []byte{'G', 'E', 'E', 'K', 'S'}
	slice_2 := []byte{'G', 'E', 'e', 'K', 'S'}

	//比较切片

	//使用Compare函数
	//如果结果为0，则slice_1 == slice_2
	//如果结果为-1，则slice_1 <slice_2
	//如果结果为+1，则slice_1> slice_2
	res := bytes.Compare(slice_1, slice_2)

	if res == 0 {
		fmt.Println("!..切片相等..!")
	} else {
		fmt.Println("!..切片不相等..!")
	}
}

func Demo8() {
	//使用简写声明，创建和初始化切片
	scl1 := []int{400, 600, 100, 300, 500, 200, 900}
	scl2 := []int{-23, 567, -34, 67, 0, 12, -5}

	//显示切片
	fmt.Println("Slices(Before):")
	fmt.Println("Slice 1: ", scl1)
	fmt.Println("Slice 2: ", scl2)

	//对整数切片进行排序

	//使用Ints函数
	//此函数仅用于对整数切片进行排序，并按升序对切片中的元素进行排序。
	sort.Ints(scl1)
	sort.Ints(scl2)

	//显示结果
	fmt.Println("\nSlices(After):")
	fmt.Println("Slice 1 : ", scl1)
	fmt.Println("Slice 2 : ", scl2)
}

func Demo9() {

	//创建和初始化切片

	//使用简写声明
	scl1 := []int{100, 200, 300, 400, 500, 600, 700}
	scl2 := []int{-23, 567, -34, 67, 0, 12, -5}

	//显示切片
	fmt.Println("Slices:")
	fmt.Println("切片1: ", scl1)
	fmt.Println("切片2: ", scl2)

	//检查切片是否为排序形式

	//使用IntsAreSorted函数
	//此函数用于检查给定的int切片是否为排序形式（按升序排列）。如果切片为排序形式，则此方法返回true；否则，如果切片未为排序形式，则返回false
	res1 := sort.IntsAreSorted(scl1)
	res2 := sort.IntsAreSorted(scl2)

	//显示结果
	fmt.Println("\nResult:")
	fmt.Println("切片1是否已排序?: ", res1)
	fmt.Println("切片2是否已排序?: ", res2)
}

func Demo10() {
	//创建和初始化
	//字节片
	//使用简写声明
	slice_1 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's',
		'f', 'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}

	slice_2 := []byte{'A', 'p', 'p', 'l', 'e'}

	slice_3 := []byte{'%', 'g', '%', 'e', '%', 'e',
		'%', 'k', '%', 's', '%'}

	//显示切片
	fmt.Println("原始切片:")
	fmt.Printf("Slice 1: %s", slice_1)
	fmt.Printf("\nSlice 2: %s", slice_2)
	fmt.Printf("\nSlice 3: %s", slice_3)

	//分割字节片
	//使用分割函数
	res1 := bytes.Split(slice_1, []byte("eek"))
	res2 := bytes.Split(slice_2, []byte(""))
	res3 := bytes.Split(slice_3, []byte("%"))

	//显示结果
	fmt.Printf("\n\n分割后:")
	fmt.Printf("\nSlice 1: %s", res1)
	fmt.Printf("\nSlice 2: %s", res2)
	fmt.Printf("\nSlice 3: %s", res3)
}

func Demo11() {

	//创建和分割
	//字节片
	//使用分割函数
	res1 := bytes.Split([]byte("****Welcome, to, cainiaojc****"), []byte(","))

	res2 := bytes.Split([]byte("Learning x how x to x trim x a x slice of bytes"), []byte("x"))

	res3 := bytes.Split([]byte("cainiaojc, Geek"), []byte(""))

	res4 := bytes.Split([]byte(""), []byte(","))

	//显示结果
	fmt.Printf("最终结果值:\n")
	fmt.Printf("\nSlice 1: %s", res1)
	fmt.Printf("\nSlice 2: %s", res2)
	fmt.Printf("\nSlice 3: %s", res3)
	fmt.Printf("\nSlice 4: %s", res4)
}
