package main

import "fmt"

func main() {
	// Deklarasi variable menggunakan `var` keyword berserta tipe data
	var myFirstCatName string = "Leon"
	var mySecondCatName string = "Linux"

	// Deklarasi variable tanpa meng-inisialisasi nilai
	var myThirdCatName string
	// Mengisi nilai ke variable yang sudah dideklarasikan
	myThirdCatName = "Loki"

	// Deklarasi variable dengan tipe data yang bisa di-infer oleh compiler
	var myFourthCatName = "Luna"

	// Deklarasi variable dengan pendeklarasian singkat (short declaration)
	myFifthCatName := "Leo"

	// Deklarasi multiple variable sekaligus
	var (
		mySixthCatName   string = "Lily"
		mySeventhCatName string = "Loki"
	)
	var myEighthCatName, myNinthCatName string
	myEighthCatName = "Luna"
	myNinthCatName = "Leo"

	myTenthCatName, myEleventhCatName := "Lily", "Loki"

	// Mencetak nilai dari variable
	fmt.Print(myFirstCatName)
	fmt.Print(mySecondCatName)
	fmt.Print(myThirdCatName)
	fmt.Print(myFourthCatName)
	fmt.Print(myFifthCatName)
	fmt.Print(mySixthCatName)
	fmt.Print(mySeventhCatName)
	fmt.Print(myEighthCatName)
	fmt.Print(myNinthCatName)
	fmt.Print(myTenthCatName)
	fmt.Print(myEleventhCatName)

	// Mencetak nilai dari variable dengan format yang lebih baik
	// Fungsi ini digunakan untuk menampilkan output dalam bentuk tertentu.
	// Kegunaannya sama seperti fungsi fmt.Println(), hanya saja struktur outputnya didefinisikan di awal.
	fmt.Printf("My first cat's name is %s\n", myFirstCatName)
	fmt.Printf("My second cat's name is %s\n", mySecondCatName)
	fmt.Printf("My third cat's name is %s\n", myThirdCatName)
	fmt.Printf("My fourth cat's name is %s\n", myFourthCatName)
	fmt.Printf("My fifth cat's name is %s\n", myFifthCatName)
	fmt.Printf("My sixth cat's name is %s\n", mySixthCatName)
	fmt.Printf("My seventh cat's name is %s\n", mySeventhCatName)
	fmt.Printf("My eighth cat's name is %s\n", myEighthCatName)
	fmt.Printf("My ninth cat's name is %s\n", myNinthCatName)
	fmt.Printf("My tenth cat's name is %s\n", myTenthCatName)
	fmt.Printf("My eleventh cat's name is %s\n", myEleventhCatName)

	// Mencetak nilai dari variable dengan format yang lebih baik menggunakan Println
	fmt.Println("My first cat's name is", myFirstCatName)
	fmt.Println("My second cat's name is", mySecondCatName)
	fmt.Println("My third cat's name is", myThirdCatName)
	fmt.Println("My fourth cat's name is", myFourthCatName)
	fmt.Println("My fifth cat's name is", myFifthCatName)
	fmt.Println("My sixth cat's name is", mySixthCatName)
	fmt.Println("My seventh cat's name is", mySeventhCatName)
	fmt.Println("My eighth cat's name is", myEighthCatName)
	fmt.Println("My ninth cat's name is", myNinthCatName)
	fmt.Println("My tenth cat's name is", myTenthCatName)
	fmt.Println("My eleventh cat's name is", myEleventhCatName)
}
