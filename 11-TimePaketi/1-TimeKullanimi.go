package main

import (
	"fmt"
	"time"
)

func main() {
	println("Suanki unix zamanı", time.Now().Unix())
	//time.Sleep(time.Second * 2)
	println("Suanki unix zamanı", time.Now().Unix())
	t := time.Date(2020, time.January, 10, 20, 0, 0, 0, time.UTC)
	fmt.Println(t) // sadece println ile calışmadı.
	fmt.Println(time.Now())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Weekday())
	tomorrow := time.Now().AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v %v %v %v \n", tomorrow.Weekday(), tomorrow.Day(), tomorrow.Month(), tomorrow.Year())
	fmt.Println(tomorrow.Format("Monday January 2, 2006"))
	fmt.Println(tomorrow.Format("01/02/2006"))

}
