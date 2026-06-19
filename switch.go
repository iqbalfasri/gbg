package main

func main() {
	// Switch statement: digunakan untuk membuat percabangan berdasarkan nilai dari sebuah variabel.
	// Struktur dasar switch adalah:
	// switch variable {
	//     case value1:
	//         // code yang dijalankan jika variable == value1
	//     case value2:
	//         // code yang dijalankan jika variable == value2
	//     default:
	//         // code yang dijalankan jika variable tidak cocok dengan case manapun
	// }
	var day int = 3

	switch day {
	case 1:
		println("Monday")
	case 2:
		println("Tuesday")
	case 3:
		println("Wednesday")
	case 4:
		println("Thursday")
	case 5:
		println("Friday")
	default:
		println("Weekend")
	}

	// Kita juga bisa menggunakan switch tanpa ekspresi untuk membuat percabangan berdasarkan kondisi.
	switch {
	case day < 3:
		println("Early in the week")
	case day == 3:
		println("Midweek")
	default:
		println("Late in the week")
	}

	// Switch dengan short statement: kita bisa menulis statement pendek sebelum kondisi switch.
	switch day := 4; day {
	case 1:
		println("Short statement: Monday")
	case 2:
		println("Short statement: Tuesday")
	case 3:
		println("Short statement: Wednesday")
	case 4:
		println("Short statement: Thursday")
	case 5:
		println("Short statement: Friday")
	default:
		println("Short statement: Weekend")
	}
	// Penjelasan: dalam contoh di atas, kita mendeklarasikan variabel day dengan nilai 4 di dalam statement switch.
	// Variabel ini hanya berlaku di dalam blok switch tersebut dan tidak dapat diakses di luar blok tersebut.

	// Penjelasan: switch adalah cara yang lebih bersih dan mudah dibaca untuk menangani banyak kondisi dibandingkan dengan if-else yang panjang.
	// Switch juga dapat digunakan untuk menggantikan if-else yang kompleks dengan banyak kondisi, sehingga membuat kode lebih mudah dipahami dan dipelihara.

	// Penjelasan: dalam contoh di atas, switch pertama menggunakan ekspresi day untuk menentukan hari apa yang sedang diproses.
	// Switch kedua tidak menggunakan ekspresi, sehingga setiap case adalah kondisi yang dievaluasi secara independen.
}
