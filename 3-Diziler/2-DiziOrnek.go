package main

import "fmt"

func main() {
	// Basit Dizi
	myArray1 := [3]int{}
	myArray1[0] = 32
	myArray1[1] = 23
	myArray1[2] = 54
	fmt.Println(myArray1)

	// Renk dizisi
	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Println(colors)
	fmt.Println(colors[1])
	colors[1] = "yellow"
	fmt.Println(colors[1])

	// Dizilerde işlemler
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println("Numbers of Numbers", len(numbers))

	// Eleman sayısını vermeden dizi oluşturma
	myArray2 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(myArray2)
	myArray2[3] = 5
	fmt.Println(myArray2)

	// cap() dizinin kapasitesini verir
	var cars [3]string
	cars[0] = "BMW"
	cars[1] = "Mercedes"
	cars[2] = "Audi"

	i := 0
	for i < len(cars) {
		fmt.Println(cars[i])
		i++
	}
	fmt.Println("**************************")
	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}
	fmt.Println("**************************")
	for i, value := range cars {
		fmt.Println("i =", i, ", value =", value)
	}
	fmt.Println("**************************")
	for i := range cars {
		fmt.Println("i =", i, ", value =", cars[i])
	}
}
