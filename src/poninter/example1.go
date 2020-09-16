package main

import "fmt"

func example() {

	var a int
	fmt.Println(&a)
	var p *int

	p = &a
	fmt.Println(p)

	*p = 20

	fmt.Println(a)
}
