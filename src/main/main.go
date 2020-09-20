package main

import (
	"fmt"
	"go-example/src/cccc"
	"go-example/src/helloworld"
	"go-example/src/poninter"
	"go-example/src/prints"
)


func main()  {
	helloworld.Hello()
	poninter.Example()
	prints.Print()
	cccc.Add()
	fmt.Println("哈哈")
}

