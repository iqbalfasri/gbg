package main

import "fmt"

func welcome(name string) string {
	return fmt.Sprintf("Welcome to Go programming, %s!", name)
}

func main() {
	// Function value adalah sebuah fungsi yang dapat disimpan dalam variabel, dilewatkan sebagai argumen ke fungsi lain, atau dikembalikan sebagai nilai dari fungsi lain.
	// Ini memungkinkan kita untuk menggunakan fungsi sebagai "first-class citizens" dalam bahasa pemrograman Go.

	// Contoh sederhana dari function value adalah fungsi yang menerima sebuah fungsi sebagai argumen dan memanggilnya.
	welcomeIqbal := welcome("Iqbal")
	fmt.Println(welcomeIqbal)
	// Fungsi welcome mengembalikan sebuah string yang berisi pesan selamat datang, dan kita menyimpan hasilnya dalam variabel welcomeIqbal.

}
