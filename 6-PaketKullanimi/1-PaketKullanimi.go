package main

import (
	"fmt"
	"math/rand"
	"strings"
)

/*
	Paket bir araya toplanmıs (bundeled) fonksiyonlar topluluğudur. Farklı bircok nesneyi de kendi içerisinde barındırabilir

	- Her Go programı paketlerden oluşur.
	- Programlar main paketinde calısmaya baslar.

	Faydaları?
	- Namespace, yani ilgili fonksiyonları kapsayan ortak isim uzayı sağlar ve proje ya da sistemdeki diğer fonksiyonlarla
	karışmayı onler.
	- Kod organizasyonu sağlar.
	- Derleyiciyi hızlarndırır. Ornegin "fmt" her kullanışımızda derleyici tarafından derlenmez.

*/
func main() {

	// fmt
	fmt.Println("Bu bir ornek")
	fmt.Println("My favorite number is", rand.Intn(10))

	// strings
	fmt.Println(strings.Contains("Hello World", "Hello"))
	fmt.Println(strings.Count("Hello World", "l"))
	fmt.Println(strings.HasPrefix("Hello_World", "Hello"))
	fmt.Println(strings.HasSuffix("dosya.rar", "rar"))
	fmt.Println(strings.Index("test", "e"))

	// github uzerinden package indirme islemlerine bak,
	// paket yapısı ve paket oluşturmaya bak
}
