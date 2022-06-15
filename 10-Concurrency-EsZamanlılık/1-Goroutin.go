package main

import (
	"runtime"
	"time"
)

// başına "go" yazılıp çalıştırılan işlem arka planda çalışmaya devam eder
// diğer işlemlerin çalıştırılması için onun bitmesi beklenmez.
func main() {
	// basit gorouting kullanımı
	go xFunc()
	time.Sleep(time.Second * 1)
	println("*************************")

	runtime.GOMAXPROCS(1) //kaç thread de çalışacağını belirtir
	go xFunc()
	time.Sleep(time.Second * 1)
}

func xFunc() {
	for l := byte('a'); l <= byte('z'); l++ {
		println(string(l))
	}
}
