package main

import "fmt"

var a = []int{1, 2, 3}

var b = []int{4, 5, 6}

var s1 = make([]int, 10)

// var s2 = append(s1, 5)

func appendEx() {

	c := append(a, b...)

	d := append(c, 7)
	fmt.Printf("原本数组", a, c)
	fmt.Printf("append", c)
	fmt.Printf("append", d)

	fmt.Printf("%p\n", &s1)
	// fmt.Printf("%p\n", &s2)

}

func copyEx() {
	fmt.Print("copy的数据", s1, a)
	copy(s1, a)
	fmt.Print("copy的数据", s1, a)
}

func forSlice() {

	data := []int{0, 2, 3, 5, 4, 6, 7, 9, 4, 5}
	slice := data[:]

	for index, value := range slice {
		fmt.Printf("inde : %v , value : %v\n", index, value)
	}

	var a = []int{1, 4, 5, 6}
	fmt.Print(a, len(a))

	b := a[1:2]

	fmt.Print(b, len(b))

	c := b[0:3]

	fmt.Print(c, len(c))

	str := "哈哈哈"

	fmt.Print(str)

	str = "cccc"

	fmt.Print(str)
}
