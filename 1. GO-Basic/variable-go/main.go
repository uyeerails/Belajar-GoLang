package main

import "fmt"

func main() {
	// Keyword var untuk deklarasi variable namaDepan
	var namaDepan string = "uyee"
	// Keyword var untuk deklarasi variable namaBelakang
	var namaBelakang string = "rails"
	// fmt.Printf() untuk menampilkan output dalam bentuk tertentu Cara 1
	fmt.Printf("hai %s %s!\n", namaDepan, namaBelakang)
	// fmt.Printf() menampilkan output dalam bentuk tertentu Cara 2
	fmt.Printf("halo", namaDepan, namaBelakang+"!")
	// fmt.Printf() menampilkan output dalam bentuk tertentu Cara 3
	fmt.Printf("hai uyee rails!\n")
}
