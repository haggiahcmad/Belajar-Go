package main

import "fmt"

func main() {
	var siswa = [...]string{
		"el",
		"al",
		"il",
		"dul",
	}
	fmt.Println(siswa)

	var slice1 = siswa[1:2]
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))

}
