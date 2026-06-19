package main

import "fmt"

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, item := range data {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	// Function parameter adalah sebuah variabel yang dideklarasikan dalam tanda kurung setelah nama fungsi, yang digunakan untuk menerima nilai saat fungsi dipanggil.
	// Function parameter memungkinkan kita untuk membuat fungsi yang lebih fleksibel dan dapat digunakan kembali dengan berbagai nilai.

	// Contoh sederhana dari function parameter adalah fungsi filter yang menerima sebuah slice data dan sebuah fungsi callback sebagai parameter.
	// Fungsi filter akan mengiterasi setiap item dalam slice data dan memanggil fungsi callback untuk menentukan apakah item tersebut harus dimasukkan dalam hasil atau tidak.

	data := []string{"apple", "banana", "cherry", "date", "elderberry"}
	// Kita memiliki sebuah slice data yang berisi nama-nama buah.

	filteredData := filter(data, func(fruit string) bool {
		if fruit == "banana" {
			return true
		}
		return false
	})

	fmt.Println(filteredData)
}
