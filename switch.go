package main

import "fmt"

func main() {

	switch nilai := 11; nilai%2 == 0 {
	case true:
		fmt.Println("Genap")
	default:
		fmt.Println("Ganjil")
	}

}
