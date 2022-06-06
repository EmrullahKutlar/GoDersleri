package main

import "fmt"

func main() {
	// 1.
	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[1] = "One"
	myMap[2] = "Two"
	fmt.Println(myMap)

	// Key value pairs

	// 2.

	states := make(map[string]string)
	states["Ist"] = "İstanbul"
	states["Ank"] = "Ankara"
	states["ANT"] = "Antalya"
	fmt.Println(states)

	// Şehir listesinde Ant anahtar adına sahip veriyi elde et
	antalya := states["ANT"]
	fmt.Println("Secili Şehir: ", antalya)

	// Ank anahtarına sahip veriyi sil
	delete(states, "Ank")
	fmt.Println(states)

	states["ERZ"] = "Erzurum"

	for k, v := range states {
		fmt.Println("Key", k, "Value", v)
	}

	// states map nesnesinin elemanı adetince kapasiteli bir key listesi oluştur

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	// key listesindeki index değerlerine göre states map nesnesindeki şehirleri ekrana yazdır

	for i := range keys {
		fmt.Println(states[keys[i]])

	}

}
