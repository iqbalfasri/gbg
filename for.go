package main

import "fmt"

func main() {
	// For loop: digunakan untuk melakukan iterasi atau pengulangan dalam program.
	// Struktur dasar for loop adalah:
	// for initialization; condition; post {
	//     // code yang dijalankan selama condition bernilai true
	// }

	counter := 1
	for counter <= 5 {
		fmt.Println("Counter:", counter)
		counter++
	}
	// Penjelasan: dalam contoh di atas, kita mendeklarasikan variabel counter dengan nilai awal 1.
	// Loop akan terus berjalan selama counter kurang dari atau sama dengan 5.
	// Di dalam loop, kita mencetak nilai counter dan kemudian menambahkannya dengan 1 setiap iterasi.

	// For loop dengan statement initialization dan post:
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}
	// Penjelasan: dalam contoh ini, kita menggunakan for loop dengan statement initialization (i := 0), condition (i < 5), dan post (i++).
	// Loop akan berjalan selama i kurang dari 5, dan setiap iterasi, nilai i akan bertambah 1.

	// For loop dengan mengambil isi nilai secara manual tanpa range:
	fruits := []string{"Apple", "Banana", "Cherry"}
	for i := 0; i < len(fruits); i++ {
		fmt.Println("Fruit:", fruits[i])
	}
	// Penjelasan: dalam contoh ini, kita menggunakan for loop untuk mengakses elemen-elemen dalam slice fruits.
	// Kita menggunakan len(fruits) untuk menentukan jumlah iterasi yang diperlukan, dan mengakses setiap elemen dengan indeks i.

	// For loop juga bisa digunakan untuk iterasi dengan menggunakan range, terutama saat bekerja dengan slice, array, atau map.
	// Contoh iterasi dengan range pada slice:
	fruitsRange := []string{"Mango", "Pineapple", "Grapes"}
	for index, fruit := range fruitsRange {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}
	// Penjelasan: dalam contoh di atas, kita menggunakan for loop dengan range untuk iterasi melalui slice fruitsRange.
	// Variabel index akan menyimpan indeks dari elemen saat ini, sedangkan variabel fruit akan menyimpan nilai dari elemen tersebut.

	// Contoh iterasi dengan range pada map:
	ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}
	for name, age := range ages {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
	// Penjelasan: dalam contoh ini, kita menggunakan for loop dengan range untuk iterasi melalui map ages.
	// Variabel name akan menyimpan kunci dari elemen saat ini, sedangkan variabel age akan menyimpan nilai dari elemen tersebut.

	// Contoh iterasi dengan range pada array:
	numbers := [5]int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	// Penjelasan: dalam contoh di atas, kita menggunakan for loop dengan range untuk iterasi melalui slice numbers.
	// Variabel index akan menyimpan indeks dari elemen saat ini, sedangkan variabel value akan menyimpan nilai dari elemen tersebut.
}
