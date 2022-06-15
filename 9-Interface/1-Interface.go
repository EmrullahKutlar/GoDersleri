package main

import (
	"strconv"
)

func CarExacute(c Carface) {
	println("Araç Bilgisi:" + c.Information() + "\n")
	msg := ""
	isRun := c.Run()
	if isRun {
		msg = "Araç Çalışıyor"

	} else {
		msg = "Araç Çalışmıyor"
	}
	println(msg)

	isStop := c.Stop()
	if isStop {
		msg = "Araç Durdu"
	} else {
		msg = "Araç Durmadı"
	}
	println(msg)

}
func main() {
	//Ferrari
	Ferrari := new(Ferrari)
	Ferrari.Brand = "Ferrari"
	Ferrari.Model = "488"
	Ferrari.Color = "Red"
	Ferrari.Speed = 300
	Ferrari.Price = 10000000
	Ferrari.Special = true
	//println(Ferrari.Information())

	//Lamborghini
	Lamborghini := new(Lamborghini)
	Lamborghini.Brand = "Lamborghini"
	Lamborghini.Model = "Huracan"
	Lamborghini.Color = "Red"
	Lamborghini.Speed = 300
	Lamborghini.Price = 10000000
	Lamborghini.Special = true
	//println(Lamborghini.Information())

	//Mercedes
	Mercedes := new(Mercedes)
	Mercedes.Brand = "Mercedes"
	Mercedes.Model = "C class"
	Mercedes.Color = "Red"
	Mercedes.Speed = 300
	Mercedes.Price = 10000000
	Mercedes.Special = false
	//println(Mercedes.Information())

	//CarExacute
	CarExacute(Ferrari)
	CarExacute(Lamborghini)
	CarExacute(Mercedes)
}

// Interface
type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

// Base structs
type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

type SpecialProduction struct {
	Special bool
}

//ferrari

type Ferrari struct {
	Car               // composition - embedding
	SpecialProduction // composition - embedding
}

func (_ Ferrari) Run() bool {
	return true
}
func (_ Ferrari) Stop() bool {
	return true
}
func (x *Ferrari) Information() string {
	ret := "\n" + x.Brand + " " + "\n" + x.Model + " " + "\n" + x.Color + " " + "\n" + strconv.Itoa(x.Speed) + " " + "\n" + strconv.Itoa(int(x.Price)) + "Million"
	add := "Yes"
	if x.Special {
		ret += "\n" + "Special:" + add
	}
	return ret
}

//Lamborghini

type Lamborghini struct {
	Car               // composition - embedding
	SpecialProduction // composition - embedding
}

func (_ Lamborghini) Run() bool {
	return true
}
func (_ Lamborghini) Stop() bool {
	return true
}
func (x *Lamborghini) Information() string {
	ret := "\n" + x.Brand + " " + "\n" + x.Model + " " + "\n" + x.Color + " " + "\n" + strconv.Itoa(x.Speed) + " " + "\n" + strconv.Itoa(int(x.Price)) + "Million"
	add := "Yes"
	if x.Special {
		ret += "\n" + "Special:" + add
	}
	return ret
}

//Mercedes

type Mercedes struct {
	Car               // composition - embedding
	SpecialProduction // composition - embedding
}

func (_ Mercedes) Run() bool {
	return true
}
func (_ Mercedes) Stop() bool {
	return true
}
func (x *Mercedes) Information() string {
	ret := "\n" + x.Brand + " " + "\n" + x.Model + " " + "\n" + x.Color + " " + "\n" + strconv.Itoa(x.Speed) + " " + "\n" + strconv.Itoa(int(x.Price)) + "Million"
	// add := "Yes"
	// if x.Special {
	// 	ret += "\n" + "\n" + "Special:" + add
	// }
	return ret
}
