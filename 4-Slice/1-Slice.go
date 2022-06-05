package main

import "fmt"

func main() {
	myArray1 := [...]int{45, 23, 54}
	mySlice1 := myArray1[:]
	fmt.Println(mySlice1)
	mySlice1[0] = 100
	fmt.Println(mySlice1)
	fmt.Println(myArray1)

	var colors = []string{"red", "green", "blue"}
	fmt.Println(colors)
	colors = append(colors, "yellow")
	fmt.Println(colors)
	colors = append(colors[1:len(colors)]) //colors = append(colors[1:4])
	fmt.Println(colors)
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println(numbers)
	numbers = append(numbers, 6)
	fmt.Println(numbers)
	fmt.Println("Numbers of Numbers", len(numbers))
	fmt.Println("Cap of Numbers", cap(numbers))

}
