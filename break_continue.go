package main

import "fmt"

func main() {
	// Break statement: digunakan untuk menghentikan eksekusi loop sebelum kondisi loop terpenuhi.
	// Contoh penggunaan break dalam for loop:
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("Break at i =", i)
			break // Menghentikan loop ketika i sama dengan 5
		}
		fmt.Println("i:", i)
	}
	// Penjelasan: dalam contoh di atas, kita menggunakan break statement untuk menghentikan loop ketika nilai i mencapai 5.
	// Loop akan berhenti dan tidak akan mencetak nilai i setelah 5.

	// Continue statement: digunakan untuk melewati iterasi saat ini dan melanjutkan ke iterasi berikutnya.
	// Contoh penggunaan continue dalam for loop:
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			// fmt.Println("Continue at j =", j)
			continue // Melewati iterasi saat ini jika j adalah bilangan genap
		}
		fmt.Println("j:", j)
	}
	// Penjelasan: dalam contoh di atas, kita menggunakan continue statement untuk melewati iterasi saat ini jika nilai j adalah bilangan genap.
	// Loop akan melanjutkan ke iterasi berikutnya tanpa mencetak nilai j yang genap.
}
