package main

import "fmt"

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func add(a int, b int) int {
	return a + b
}

func multipleValues() (int, int) {
	return 10, 20
}

func main() {
	fmt.Println("Hello, World!")
	sayHello("Leon")

	sum := add(5, 3)
	fmt.Println("Sum:", sum)

	value1, value2 := multipleValues()
	fmt.Println("Multiple Values:", value1, value2)

	// Skip value yang tidak digunakan
	value1Baru, _ := multipleValues()
	fmt.Println("Value 1 (baru):", value1Baru)

}
