package main

import "fmt"

//给了初始值1,2,3，默认值为0
var arr0 [5]int = [5]int{1, 2, 3}

var arr1 = [5]int{1, 2, 3, 4, 5}

var arr2 = [...]int{1, 2, 3, 4, 5, 6}

var str = [5]string{3: "hello world", 4: "tom"}

func arrayInit() {

	//局部注册
	a := [5]int{1, 2}

	b := [5]string{"a", "b"}

	c := [...]struct {
		name string
		age  int
	}{
		{"user1", 10},
		{"user2", 20},
	}

	fmt.Println(a, b, c)

}

func main2() {
	arrayInit()
	fmt.Println(arr0, arr1, arr2, str[0])
}
