package main

import "strconv"

func main() {

	// Kullanıcı oluşturma v1
	user1 := &User{
		FirstName: "Ali",
		LastName:  "Kılıç",
		UserName:  "ali.kilic",
		Pay: &Payment{
			Salary: 1000,
			Bonus:  500,
		},
	}
	println(user1)
	println(user1.GetFullName())
	println(user1.GetUserName())
	println(strconv.FormatFloat(user1.GetPayment(), 'f', 2, 64))
	println("------------------------------")

	// Kullanıcı Oluşturma v2
	user2 := newUser()
	user2.FirstName = "Ahmet"
	user2.LastName = "Kılıç"
	user2.UserName = "ahmet.kilic"
	user2.Age = 20
	user2.Pay.Salary = 2000
	user2.Pay.Bonus = 1000
	println(user2.GetFullName())
	println(user2.GetUserName())
	println(user2.Age)
	println(strconv.FormatFloat(user2.GetPayment(), 'f', 2, 64))
}

// Kullanıcı Yapısı
type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

// Kullanıcının Yapıcı Metodu
func newUser() *User {
	user := new(User)
	user.Pay = newPayment()
	return user
}

// Payment Yapısı
type Payment struct {
	Salary float64
	Bonus  float64
}

// Odemenin Yapıcı Metodu
func newPayment() *Payment {
	payment := new(Payment)
	return payment
}

// Metotlar
func (u User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) GetUserName() string {
	return u.UserName
}
func (u *User) GetPayment() float64 {
	pay := u.Pay.Salary + u.Pay.Bonus
	return pay
}
