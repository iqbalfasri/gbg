package main

import "fmt"

func main() {
	// if-else statement: digunakan untuk membuat percabangan dalam program.
	// Struktur dasar if-else adalah:
	// if condition {
	//     // code yang dijalankan jika condition bernilai true
	// } else {
	//     // code yang dijalankan jika condition bernilai false
	// }
	var age int = 20

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// Kita juga bisa menggunakan else if untuk membuat beberapa kondisi.
	if age < 13 {
		fmt.Println("You are a child.")
	} else if age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are an adult.")
	}

	// If dengan short statement: kita bisa menulis statement pendek sebelum kondisi if.
	if age := 25; age >= 18 {
		fmt.Println("Short statement: You are an adult.")
	} else {
		fmt.Println("Short statement: You are a minor.")
	}

	// Penjelasan: dalam contoh di atas, kita mendeklarasikan variabel age dengan nilai 25 di dalam statement if.
	// Variabel ini hanya berlaku di dalam blok if-else tersebut dan tidak dapat diakses di luar blok tersebut.
}
