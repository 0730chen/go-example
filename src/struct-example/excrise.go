package struct_example

import "fmt"

type User struct {
	Name string
	Age int
}

func run ()  {
	user := User{
		"我是一个结构体",
		1,
	}

	fmt.Println(user)
}