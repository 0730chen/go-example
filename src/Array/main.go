package main

import "fmt"

func main() {

	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b)
	c := *b
	fmt.Printf("value of c:%v\n", c)
	// b:0xc00001a078 type:*int
	fmt.Println(&b) // 0xc00000e018
	fmt.Println("哈哈")
	// sliceEm()
	// sliceOp()
	run()
	// appendEx()
	// copyEx()
	// forSlice()
	address()
}
