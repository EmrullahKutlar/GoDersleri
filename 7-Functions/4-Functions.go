package main

// Anonim Fonksiyonlar

func main() {
	addFunc := func(terms ...int) (numterms int, sum int) {
		for _, term := range terms {
			sum += term
		}
		numterms = len(terms)
		return
	}
	numTerms, sum := addFunc(1, 2, 3, 4, 5)
	println("Added: ", numTerms, "terms for a total of", sum)
}
