package main

import (
	"fmt"

	"github.com/goeasy2022/demo/sum"
	var1 "github.com/goeasy2022/demo/var"
)

func main() {
	fmt.Println("Hello World!")
	a := 1
	b := 2
	c := sum.Add(a, b)
	fmt.Println(c)
	var1.Mul(3, 4)
}
