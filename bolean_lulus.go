package main

import "fmt"

func main() {
	var nilaiAwal = 80
	var nilaiAkhir = 90

	var nilaiAwalLulus bool = nilaiAwal >= 80
	var nilaiAkhirLulus bool = nilaiAkhir >= 80

	var lulus = nilaiAwalLulus && nilaiAkhirLulus

	fmt.Println(lulus)

}
