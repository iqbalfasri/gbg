package main

import "fmt"

func log() {
	fmt.Println("Log print")
}

func runApp() {
	defer log() // Fungsi log() akan dieksekusi setelah runApp() selesai dieksekusi
	fmt.Println("Run App")
}

func main() {
	// Defer statement digunakan untuk menunda eksekusi sebuah fungsi sampai fungsi yang memanggilnya selesai dieksekusi.
	// Defer biasanya digunakan untuk membersihkan, seperti menutup file atau koneksi jaringan, setelah selesai digunakan.

	// Penjelasan: jadi setelah runApp() selesai dieksekusi, maka fungsi log() akan dieksekusi dan mencetak "Log print" ke layar.
	// daripada kita memanggil log() di akhir runApp(), kita bisa menggunakan defer untuk menunda eksekusi log() sampai runApp() selesai dieksekusi.

	runApp()
}
