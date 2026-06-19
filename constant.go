package main

import "fmt"

func main() {
	// Constant ini digunakan untuk menyimpan nilai yang tidak berubah selama program berjalan
	const API_KEY = "blablabla"

	// Constant lebih dari satu value
	const (
		API_KEY2 = "blablabla"
		API_URL  = "https://api.example.com"
	)

	fmt.Println(API_KEY)
	fmt.Println(API_KEY2)
	fmt.Println(API_URL)
}
