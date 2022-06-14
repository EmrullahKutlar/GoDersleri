package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

// varsayılan ve boş yapıcı metot
func NewHuman() *Human {
	h := new(Human)
	return h
}
func NewHuman2(firstName, LastName string, age int) *Human {
	h := new(Human)
	h.FirstName = firstName
	h.LastName = LastName
	h.Age = age
	return h
}
func main() {
	// nesne orneği oluşturma
	ali := Human{
		FirstName: "Ali",
		LastName:  "Kılıç",
		Age:       30,
	}
	fmt.Println(ali.FirstName)

	x := new(Human)
	x.FirstName = "Ahmet"
	fmt.Println(x.FirstName)

	// yapıcı metot kullanımı
	y := NewHuman()
	y.Age = 20
	println(y.Age)

	//parametreli yapıcı metot kullanımı
	z := NewHuman2("Mehmet", "Kılıç", 30)
	println(z.FirstName)

	//veriyi okuyalım
	var data = z.FirstName + " " + z.LastName + "/" + strconv.Itoa(z.Age)
	println(data)
}
