package main

import "fmt"

func main() {
	var name [4]string
	name[0] = "jui"
	name[1] = "joi"
	name[2] = "jvi"
	name[3] = "jou"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])
	fmt.Println(name[3])

	var values = [3]int{
		60,
		70,
		80,
	}
	fmt.Println(values)
}
