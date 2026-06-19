package data_type

import "fmt"

func main() {
	// Tipe data string
	var myFirstName string = "Iqbal"
	var myLastName string = "Fasri"

	// Tipe data integer
	var myAge int = 30

	// Tipe data float
	var myHeight float64 = 1.75

	// Tipe data boolean
	var isMarried bool = true

	// Mencetak nilai dari variable
	fmt.Printf("My name is %s %s\n", myFirstName, myLastName)
	fmt.Printf("I am %d years old\n", myAge)
	fmt.Printf("My height is %.2f meters\n", myHeight)
	fmt.Printf("Am I married? %t\n", isMarried)
}
