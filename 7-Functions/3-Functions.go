package main

import "fmt"

// Fonksiyonlar: değişken sayıda parametre alan foksiyonlar (variadic functions)
func main() {
	sayHello("Hello", "Go", "Developers", ".")
	result := add(1, 3, 5, 7, 9)
	println(result)
}
func sayHello(messages ...string) {
	for _, message := range messages {
		fmt.Println(message)
	}
}
func add(terms ...int) int {
	sum := 0
	for _, term := range terms {
		sum += term
	}
	return sum
}
