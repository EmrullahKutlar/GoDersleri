package main

import "fmt"

func main() {
	// Konsol : Veri Çıkışı
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	istrue := true
	stringLength, _ := fmt.Println(str1, str2, str3)
	// if err == nil {
	fmt.Println("string length:", stringLength)
	// }
	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", istrue)
	fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber))
	fmt.Printf("Data types: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, istrue)

	myString := fmt.Sprintf("Data types as var: %T,%T,%T,%T, and %T", str1, str2, str3, aNumber, istrue)
	fmt.Print(myString)

}
