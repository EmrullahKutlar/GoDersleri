package main

import (
	"fmt"
	"time"
)

func main() {

	// Date() metodu ile kendi tarih verimizi oluşturuyoruz.
	t := time.Date(2022, time.April, 25, 18, 15, 0, 0, time.UTC)

	// t adıyla tarih tipinde oluşturulan veriyi string tipinde ekrana yazdırıyoruz
	println("Go Launch at %s", t)
	println("----------------------")

	// time.Now() ile şuanın zaman bilgisini alıyoruz
	now := time.Now()
	println("The time now is: %s", now)
	println("----------------------")

	// İlk başta oluşturduğumuz t adındaki zaman bilgisinden Ay, Gün ve Haftanın Günü
	// bilgilerini yazdırıyoruz.
	fmt.Println("The month is", t.Month())
	fmt.Println("The day is", t.Day())
	fmt.Println("The weekday is", t.Weekday())
	fmt.Println("-------------")

	// Tarihe 1 gün ekle!
	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v\n",
		tomorrow.Weekday(),
		tomorrow.Month(),
		tomorrow.Day(),
		tomorrow.Year())
	fmt.Println("----")

	longFormat := "Monday, Junuary 2, 2006"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))
	fmt.Println("----")
	shortFormat := "1/2/06"
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))

}
