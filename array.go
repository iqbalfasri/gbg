package main

func main() {
	// Array adalah urutan bernomor dari elemen dengan panjang tertentu.
	// Dalam kode Go biasa, `slice` jauh lebih umum;
	// array berguna dalam beberapa skenario khusus.

	// Deklarasi array dengan panjang tetap.
	var myArray [5]int

	// Inisialisasi array dengan nilai.
	myArray = [5]int{1, 2, 3, 4, 5}

	// Inisialisasi tanpa menyebutkan panjang, Go akan menghitungnya secara otomatis.
	myArrayWithoutLength := [...]int{1, 2, 3, 4, 5}

	// Array juga dapat diinisialisasi dengan cara yang lebih singkat.
	anotherArray := [5]int{10, 20, 30, 40, 50}

	// Akses elemen array menggunakan indeks (mulai dari 0).
	firstElement := myArray[0]  // 1
	secondElement := myArray[1] // 2

	// Panjang array dapat diperoleh dengan fungsi len().
	arrayLength := len(myArray) // 5

	println("First element:", firstElement)
	println("Second element:", secondElement)
	println("Array length:", arrayLength)
	println("Original array first element:", myArray[0])                    // Tetap 1
	println("Another array first element:", anotherArray[0])                // 10
	println("Array without length first element:", myArrayWithoutLength[0]) // 1
}
