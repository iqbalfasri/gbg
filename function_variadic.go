package main

func sumNumbers(numbers ...int) int {
	total := 0

	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}

	return total
}

func main() {
	// Variadic function adalah fungsi yang dapat menerima jumlah argumen yang tidak terbatas.
	// Variadic function menggunakan tanda "..." sebelum tipe data parameter terakhir.
	// Contoh variadic function
	sum := sumNumbers(1, 2, 3, 4, 5)
	println("Sum:", sum)
}
