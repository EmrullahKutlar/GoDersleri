package main

//defer kullanılan yer her zaman fonksiyonun en sonunda çalışır.
//Diğer işlermlerin bitmesini bekler.
var isConnected bool = false

func main() {
	println("Starting database process.", isConnected)
	databaseProcess()
	println("Finished database process.", isConnected)
}
func databaseProcess() {
	connect()
	println("Defering disconnect.")
	defer disconnect()
	println("Connection open:", isConnected)
	println("Doing something...")
}
func connect() {
	isConnected = true
	println("Connecting...")
}
func disconnect() {
	isConnected = false
	println("Disconnecting...")
}
