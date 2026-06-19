package main

import "fmt"

func main() {
	myFirstName := "Iqbal"
	myLastName := "Fasri"
	_ = myFirstName + " " + myLastName

	myFirstName = "Already changed"

	fmt.Println("Hello, World!")
	fmt.Println("My name is", myFirstName, myLastName)
}
