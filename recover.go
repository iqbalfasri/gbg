package main

import "fmt"

// log adalah fungsi yang dipanggil via defer untuk menangkap panic menggunakan recover.
// recover() hanya berguna jika dipanggil dalam fungsi yang dijalankan dengan defer.
func log() {
	fmt.Println("Log...")
	message := recover() // menangkap nilai panic, mengembalikan nil jika tidak ada panic
	fmt.Println("Telat terjadi error, message: ", message)
}

// runApp akan memicu panic jika parameter error bernilai true.
// defer log() memastikan log tetap dieksekusi meskipun terjadi panic.
func runApp(error bool) {
	defer log()

	if error {
		panic("Error run app") // menghentikan eksekusi normal dan memicu panic
	}

}

func main() {
	runApp(true) // memicu panic, tetapi akan ditangkap oleh recover di log()
}
