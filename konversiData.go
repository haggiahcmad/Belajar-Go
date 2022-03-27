package main

import "fmt"

func main() {
	var nilai8 int32 = 20
	var nilai64 int64 = int64(nilai8)
	var nilai120 int8 = int8(nilai64)

	fmt.Println(nilai8)
	fmt.Println(nilai64)
	fmt.Println(nilai120)

	var nama = "haggi"
	var g = nama[3]
	var eString = string(g)

	fmt.Println(nama)
	fmt.Println(eString)
}
