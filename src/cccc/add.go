package cccc

import "fmt"

var name = "哈哈哈"
func Add() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	for i := 0; i < len(runeS2); i++ {
		fmt.Println(string(runeS2[i]))
	}
	fmt.Println(string(runeS2))
}

