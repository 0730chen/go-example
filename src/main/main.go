package main

import (
	"fmt"
	"go-example/src/cccc"
	"go-example/src/helloworld"
	"go-example/src/poninter"
	"go-example/src/prints"
	"go-example/src/condition"
)


func main()  {
	helloworld.Hello()
	poninter.Example()
	condition.Run()
	prints.Print()
	cccc.Add()
	fmt.Println("哈哈")
}

