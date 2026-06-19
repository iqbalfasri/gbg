package main

import "fmt"

// Type declaration: membuat tipe baru berdasarkan tipe dasar.
// Tipe baru ini disebut distinct type, bukan alias.
// Contoh: NoHP menyimpan string, tapi berbeda dengan tipe string biasa.

type NoHP string
type UserID int

type Age uint8

func main() {
	var nomor NoHP = "081234567890"
	var id UserID = 12345
	var age Age = 27
	var leonAge Age = 3

	fmt.Println("Leon is", leonAge, "years old")

	fmt.Println("NoHP:", nomor)
	fmt.Println("UserID:", id)
	fmt.Println("Age:", age)

	// Konversi eksplisit diperlukan ketika menggunakan tipe dasar lain.
	// UserID bukan int, jadi harus dikonversi secara manual.
	var x int = 42
	var convertedID UserID = UserID(x)
	fmt.Println("Converted UserID:", convertedID)
}
