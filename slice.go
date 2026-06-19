package main

import "fmt"

func main() {
	// Slice adalah referensi ke bagian dari array.
	// Slice memiliki panjang dan kapasitas yang dapat berubah-ubah.
	// Slice lebih fleksibel daripada array dan lebih umum digunakan dalam Go.
	months := [...]string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}
	slice1 := months[0:3] // January, February, March
	// Ambil dari yang pertama sampe dengan 6
	slice2 := months[:6] // January, February, March, April, May, June
	// Ambil dari yang 6 sampe dengan terakhir
	slice3 := months[6:] // July, August, September, October, November, December
	// Ambil semua elemen
	slice4 := months[:] // Semua bulan

	fmt.Println(slice1) // Output: [January February March]
	fmt.Println(slice2) // Output: [January February March April May June]
	fmt.Println(slice3) // Output: [July August September October November December]
	fmt.Println(slice4) // Output: [January February March April May June July August September October November December]

	// Cara membuat slice dengan make()
	// Sebenernya apa itu kegunaan fungsi make() itu sendiri?
	// Fungsi make() digunakan untuk membuat slice, map, atau channel dengan ukuran dan kapasitas tertentu.
	// Dalam contoh ini, kita membuat slice string dengan panjang 2 dan kapasitas 5.
	// Artinya, slice ini dapat menampung hingga 5 elemen, tetapi saat ini hanya memiliki 2 elemen yang diinisialisasi.
	sliceWithMake := make([]string, 2, 5) // Penjelasan: make([]Tipe, panjang, kapasitas)

	// Isi slice dengan nilai, misalnya dengan nama kucing
	sliceWithMake[0] = "Leon"                        // Penjelasan: indeks 0 diisi dengan "Leon"
	sliceWithMake[1] = "Linux"                       // Penjelasan: indeks 1 diisi dengan "Luna"
	fmt.Println("Slice dengan make:", sliceWithMake) // Output: Slice dengan make: [Leon Linux]

	// Lalu cara menambahkan elemen baru ke dalam slice itu harus menggunakan fungsi append()
	sliceWithMakeNew := append(sliceWithMake, "Luna")      // Penjelasan: menambahkan "Luna" ke sliceWithMake
	fmt.Println("Slice setelah append:", sliceWithMakeNew) // Output: Slice setelah append: [Leon Linux Luna]

	// Apakah kita bisa mengubah nilai dari slice yang sudah dibuat dengan make()?
	sliceWithMake[0] = "Leo"                            // Penjelasan: mengubah nilai indeks 0 menjadi "Leo"
	fmt.Println("Slice setelah diubah:", sliceWithMake) // Output: Slice setelah diubah: [Leo Linux]

	// Lalu coba kita print slice yang sudah di append, apakah nilai yang tadi diubah juga ikut berubah?
	fmt.Println("Slice setelah diubah dan di append:", sliceWithMakeNew) // Output: Slice setelah diubah dan di append: [Leo Linux Luna]

	// Penjelasan: sliceWithMake dan sliceWithMakeNew sebenarnya merujuk ke array yang sama di belakang layar,
	// jadi ketika kita mengubah nilai di sliceWithMake, perubahan tersebut juga terlihat di sliceWithMakeNew karena keduanya berbagi data yang sama.

	// Lalu bagaimana jika kita menambahkan elemen sampe melebihi kapasitas yang sudah di tentukan sebelumnya?
	sliceWithMakeNew = append(sliceWithMakeNew, "Luna2")        // Penjelasan: menambahkan "Luna2" ke sliceWithMakeNew
	fmt.Println("Slice setelah append lagi:", sliceWithMakeNew) // Output: Slice setelah append lagi: [Leo Linux Luna Luna2]

	// Penjelasan: ketika kita menambahkan elemen ke slice yang sudah mencapai kapasitasnya,
	// Go akan membuat array baru dengan kapasitas yang lebih besar (biasanya dua kali lipat) dan menyalin elemen-elemen lama ke array baru tersebut.
	// Oleh karena itu, setelah menambahkan "Luna2", sliceWithMakeNew sekarang merujuk ke array baru, sementara sliceWithMake masih merujuk ke array lama.
	// Jadi perubahan pada sliceWithMake tidak akan mempengaruhi sliceWithMakeNew setelah kapasitas awal terlampaui.

	// Cara copy slice ke slice lain menggunakan fungsi copy()
	// Fungsi copy() digunakan untuk menyalin elemen dari satu slice ke slice lain. Fungsi ini mengembalikan jumlah elemen yang berhasil disalin.
	// Dalam contoh ini, kita membuat slice baru dengan panjang yang sama dengan slice yang akan disalin, lalu menggunakan copy() untuk menyalin elemen-elemen dari
	// sliceWithMakeNew ke sliceCopy.
	sliceCopy := make([]string, len(sliceWithMakeNew), cap(sliceWithMakeNew)) // Membuat slice baru dengan panjang yang sama
	copy(sliceCopy, sliceWithMakeNew)                                         // Menyalin elemen dari sliceWithMakeNew ke sliceCopy
	fmt.Println("Slice setelah copy:", sliceCopy)                             // Output: Slice setelah copy: [Leo Linux Luna Luna2]
}
