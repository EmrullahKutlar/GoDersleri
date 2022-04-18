package main

import "strconv"

func main() {

	var myString = "1"
	var x = 10
	var f float32 = 2.0
	println(myString, x, f)

	// string to int
	number, _ := strconv.Atoi(myString)
	println("Sonuc:", number)
	result := number + 2
	println(result)
	println("*******************************")
	//int to string
	newString := strconv.Itoa(x)
	println(newString)
	println("*******************************")

	//Casting
	var i int = 55
	var a float64 = float64(i)
	var u uint = uint(a)
	println(i)
	println(a)
	println(u)
	println("*******************************")

}
