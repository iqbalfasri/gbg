package main

import "fmt"

// User adalah struct dengan dua field: Name dan Email
type User struct {
	Name, Email string
}

// getUserEmail adalah fungsi biasa (standalone function)
// yang menerima parameter bertipe User dan mengembalikan Email-nya
func getUserEmail(user User) string {
	return user.Email
}

// sayHello adalah method milik struct User.
// (user User) disebut "receiver" — mengikat fungsi ini ke tipe User.
// Method bisa diakses lewat instance User, misal: user.sayHello("WOW")

// Catatan penting: Di Go, method TIDAK bisa didefinisikan di dalam body struct.
// Berbeda dengan OOP klasik (Java/PHP/Python), struct di Go hanya untuk DATA (field).
// Method ditulis di luar struct dengan receiver yang mengikatnya ke tipe tersebut.
// Ini memisahkan data dan perilaku secara eksplisit.
func (user User) sayHello(name string) {
	fmt.Println("Hello ", name, "your email is ", user.Email)
}

// --- Contoh real-world sederhana: E-Commerce ---
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

// Method untuk mengecek ketersediaan stok
func (p Product) IsAvailable(qty int) bool {
	return p.Stock >= qty
}

// Method untuk menghitung total harga
func (p Product) TotalPrice(qty int) float64 {
	return p.Price * float64(qty)
}

type CartItem struct {
	Product  Product
	Quantity int
}

type Order struct {
	ID     int
	Items  []CartItem
	Status string
}

// Method untuk menambahkan item ke order
func (o *Order) AddItem(p Product, qty int) {
	if !p.IsAvailable(qty) {
		fmt.Printf("Stok %s tidak cukup (tersedia: %d)\n", p.Name, p.Stock)
		return
	}
	o.Items = append(o.Items, CartItem{Product: p, Quantity: qty})
	p.Stock -= qty
}

// Method untuk menghitung total belanja
func (o Order) GrandTotal() float64 {
	var total float64
	for _, item := range o.Items {
		total += item.Product.TotalPrice(item.Quantity)
	}
	return total
}

// Method untuk menampilkan struk order
func (o Order) PrintReceipt() {
	fmt.Printf("\n=== ORDER #%d ===\n", o.ID)
	for _, item := range o.Items {
		fmt.Printf("%s x%d = %.0f\n", item.Product.Name, item.Quantity, item.Product.TotalPrice(item.Quantity))
	}
	fmt.Printf("Total: %.0f\n", o.GrandTotal())
}

func main() {
	// Contoh User dari materi sebelumnya
	var user User
	user.Email = "email@mail.com"
	user.Name = "Iqbal"

	fmt.Println("User :", user)
	fmt.Println("User email :", getUserEmail(user))
	user.sayHello("WOW")

	// --- Contoh real-world ---
	laptop := Product{ID: 1, Name: "Laptop", Price: 12000000, Stock: 5}
	mouse := Product{ID: 2, Name: "Mouse", Price: 250000, Stock: 10}

	order := Order{ID: 1001}
	order.AddItem(laptop, 2)
	order.AddItem(mouse, 3)

	order.PrintReceipt()
}
