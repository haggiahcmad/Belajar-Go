package main

import "fmt"

func main() {
	var a = 20
	var b = 10
	var c = a + b
	fmt.Println(c)

	c += 30
	fmt.Println(c)
	c++
	fmt.Println(c)

	c--
	fmt.Println(c)
}
