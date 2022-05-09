package main

// range
/*
 for döngüsünün range ("aralık") formu ile bir dilim ("slice") veya eşlem ("map") üzerinde dolaşılır.
 Bir dilim üzerinde dolaşılırken her seferinde iki değer dönülür. İlki index, ikincisi o index elemanın bir kopyası.

 Değer dizisinin indexlerini veya o indexe karşılık gelen değerleri ataması "_" ile atlayabilirsiniz.
 Yalnızca indexi kullanmak istiyorsanız ", value" kısmını tamamen çıkarınız.
*/

func main() {
	// array
	/*pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for index, value := range pow {
		println("2 üzeri", index, "=", value)
	}
	a := [...]string{"a", "b", "c", "d", "e"}

	for i := range a {
		println("Array item", i, "is", a[i])
	}*/

	// map
	baskentler1 := map[string]string{"Turkey": "Istanbul", "France": "Paris", "Japan": "Tokyo", "İtaly": "Roma"}
	for key := range baskentler1 {
		println("Capital of", key, "is", baskentler1[key])
	}
	println("****************************")
	// ya da aşşağıdaki gibi de olabilir
	for key, value := range baskentler1 {
		println("Capital of", key, "is", value)
	}

}
