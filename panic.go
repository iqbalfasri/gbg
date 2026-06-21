package main

import "fmt"

func log() {
	fmt.Println("Log...")
}

func runApp(error bool) {
	defer log()

	if error {
		panic("Error run app")
	}

}

func main() {
	// Panic biasanya di pake untuk menghentikan program, namun `defer` akan tetep di eksekusi
	runApp(true) // Jika error = true, maka function `defer log()` akan tetep di eksekusi
}
