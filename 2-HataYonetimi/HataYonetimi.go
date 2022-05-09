package main

import (
	"fmt"
	"os"
)

func main() {
	// Eger error gostermek istemiyorsak err yerine " _ " koymamız gerekiyor
	/*file, err := os.Open("dosyam.txt")
	if err != nil {
		println("Error-", err.Error())
	}
	println("File-", file.Name())
	*/
	// Kendi hatamızı tanımlamak için
	/*myError := errors.New("Bu bir hata!")
	println("MyError:", myError.Error())*/

	_, err := os.Open("abc.rar")
	checkError(err)
}
func checkError(err error) {
	if err != nil {

		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
}
