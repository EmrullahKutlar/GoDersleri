package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Konsol veri girisi

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text:")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Println("Enrter a number:")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value ıf number:", f)
	}

}
