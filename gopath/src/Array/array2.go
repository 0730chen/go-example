package main

import "fmt"

//声明一个切片
func xxx() {
	//全局定义
	var s1 []int

	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}

	//函数内部定义
	s2 := []int{}
	//make方式定义
	var s3 []int = make([]int, 0)

	//初始化赋值
	var s4 []int = make([]int, 0, 0)

	s5 := []int{1, 2, 3}

	//从数组中切片

	s6 := [5]int{1, 2, 3, 4, 5}

	var s7 []int

	//包含第一个但是不包含末尾的
	s7 = s6[1:4]

	a := [5]int{1, 2}
	fmt.Println("哈哈", a, s1, s2, s3, s4, s5, s7)
}

func run() {
	xxx()
}
