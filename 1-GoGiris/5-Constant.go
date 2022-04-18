package main

type Brand string

const (
	FACEBOOK  Brand = "Facebook"
	MICROSOFT Brand = "Microsoft"
	GOOGLE    Brand = "Google"
	DIJIBIL   Brand = "Dijibil"
)

func printBrand(brand Brand) {
	println(brand)
}

func main() {

	printBrand(GOOGLE)
	printBrand(MICROSOFT)
	printBrand(DIJIBIL)

}
