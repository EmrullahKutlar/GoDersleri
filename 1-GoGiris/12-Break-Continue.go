package main

func main() {
	i := 1
	// break
	/*for {
		// i 3'e eşit veya büyükse döngüden çıkıyor. Tekrar girmez
		if i >= 5 {
			break
		}
		println("i nin değeri:", i)
		i++
	}
	println("döngü bitti")*/

	for i = 0; i < 7; i++ {
		// continue döngüyü kırmaz altındaki işlemleri atlayıp başa doner
		if i == 3 {
			continue
		} else if i == 5 {
			break
		}
		println("i nin değeri:", i)
	}
	println("döngü bitti")
}
