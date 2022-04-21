// Dalam sebuah project minimal harus ada sebuah package main
package main

// import digunakan untuk mengimport package, dalam hal ini fmt yanf fungsinya untuk keperluan I/O
import "fmt"

// Dalam Sebuah Project go harus ada file program dengan fungsi bernama main()
func main() {
	// Ini adalah Komentar
	/*
		Ini adalah Komentar Juga
		Tapi Ini Multiline
	*/
	// Fungsi fmt.Println() untuk memunculkan text ke layar, berada dalam package fmt
	fmt.Println("Hello, Saya Dari Go-Lang!!!")
	// Ini tidak akan dieksekusi
}
