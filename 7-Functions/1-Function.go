package main

import "fmt"

func main() {
	message := "Hi Go!"
	sayHello(&message)
	fmt.Println(message)
	/*********************************/
	fmt.Println(add(1, 2, "3"))
	/********************************/
	s1, s2 := swap("Hello", "World")
	fmt.Println(s1, s2)

	/********************************/
	numTerms, sum := add2(1, 2, 3, 4, 5)
	println("Added: ", numTerms, "terms for a total of", sum)
}

func sayHello(message *string) {
	// pointer olarak gönderilmeseydi buradaki message değişkenin değeri yukarıya etki etmezdi
	println(*message)
	*message = "Hello Golang"
}

func add(x, y int, z string) int /* geri dönülecek olan tip */ {
	return x + y
}

// coklu geri donus

func swap(x, y string) (string, string) {
	return x, y
}

func add2(terms ...int) (int, int) {
	total := 0
	for _, term := range terms {
		total += term
	}
	return len(terms), total
}
