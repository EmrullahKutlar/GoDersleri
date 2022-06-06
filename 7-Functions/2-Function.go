package main

import "fmt"

func main() {
	/*
		Go nun sonuc değerleri aynı bir değişken gibi adlandırılır.
		böyle yapıldığında değişken sankı fonksiyonun başında tanımlaş gibi muamele görür.

		Sonuçlar manalarına, işlevlerine göre adlandırımalıdır.
		Argümansız bir return deyimi adlandırılmış sonuc değerlerini döndürür. Bu "çıplak"
		return olarak bilinir. Çıplak return deyimleri, aynı örnekteki gibi, sadece kısa fonksiyonlarda kullanlmalıdır.
		Uzun fonksiyonlarda kullanılırsa kodun okunabilirliğini azaltırlar.
	*/
	// 1.
	fmt.Println(split(17))

	// 2.
	numTerms, sum := add(1, 2, 3, 4, 5)
	fmt.Println("Added: ", numTerms, "terms for a total of", sum)
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func add(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	numTerms = len(terms)
	return
}
