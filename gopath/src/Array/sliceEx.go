package main

import "fmt"

var arr = [...]int{0, 1, 2, 3, 4, 5, 6}
var slice0 []int = arr[2:4]
var slice1 []int = arr[0:6]
var slice2 []int = arr[2:5]

var slice3 []int = make([]int, 10)

//macke([]type,len,cap) type是类型，len是大小，cap是容量，容量必须大于等于长度
var slice4 []int = make([]int, 10, 10)

func sliceEm() {

	//切片定义

	fmt.Println("切片", arr, slice0, slice1, slice2)
}

func sliceOp() {
	data := [...]int{1, 2, 3, 4, 5, 8, 4, 1}

	sl := [...]int{1, 2, 3, 5, 4, 8: 100}

	s2 := make([]int, 6, 8) //使用mack指定len和cap值

	s3 := make([]int, 6) //不指定cap，默认等于len

	//s是一个切片
	s := data[0:4]
	s[0] += 100
	s[1] += 200

	s4 := []int{0, 1, 2, 3}

	p := &s[2] // *int, 获取底层数组元素指针。
	*p += 100

	//数组指针

	fmt.Println("我是一个切片", s, sl, s2, s3, s4)
}
