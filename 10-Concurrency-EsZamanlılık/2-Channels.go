package main

import (
	"runtime"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func main() {
	// Ornek1
	s := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sum(s[:len(s)], c)
	go sum(s[:len(s)/2], c)
	x := <-c
	y := <-c
	// yada
	// x, y := <-c, <-c
	//println(x, y, x+y)
	println(x)
	println(y)

	// Ornek2
	runtime.GOMAXPROCS(8)
	ch := make(chan string)
	go xFunc(ch)
	go printer(ch)
	time.Sleep(100 * time.Millisecond)
}
func xFunc(ch chan string) {
	for l := byte('a'); l <= byte('z'); l++ {
		ch <- string(l)
	}
}
func printer(ch chan string) {
	for {
		println(<-ch)
	}
}
