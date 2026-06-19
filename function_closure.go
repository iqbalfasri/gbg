package main

import "fmt"

func main() {
	// Function closure adalah sebuah fungsi yang dapat mengakses variabel dari lingkup luar (enclosing scope) meskipun fungsi tersebut sudah selesai dieksekusi.
	// Ini memungkinkan kita untuk membuat fungsi yang memiliki "state" atau data yang tersimpan di dalamnya.

	// Contoh sederhana dari function closure adalah fungsi yang mengembalikan fungsi lain yang dapat mengakses variabel dari fungsi induknya.
	add := func(x int) func(int) int {
		return func(y int) int {
			return x + y
		}
	}
	// Fungsi add menerima sebuah integer x dan mengembalikan sebuah fungsi yang menerima integer y dan mengembalikan hasil penjumlahan x dan y.
	// Fungsi yang dikembalikan ini adalah sebuah closure karena dapat mengakses variabel x dari fungsi induknya.

	// Create a closure with x = 5
	addFive := add(5)
	// addFive sekarang adalah sebuah fungsi yang dapat menambahkan 5 ke angka yang diberikan sebagai argumen.

	// Use the closure to add 5 to different numbers
	fmt.Println(addFive(10)) // Output: 15
	fmt.Println(addFive(20)) // Output: 25
}
