package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s shape) {
	fmt.Println("Şeklib alani: ", s.area())
}

func Demo1() {
	r := rectangle{width: 10, height: 6}
	school(r)

	c := circle{radius: 10}
	school(c)
}

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}

	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, address: "Ankara"}
	credit2 := Mortgage{rate: 12, creditPaymentTotal: 50000, address: "İstanbul"}
	credit3 := Car{rate: 15, creditPaymentTotal: 60000, carInfo: "Polo"}

	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)

	fmt.Println("Toplam odeme: ", total)
}

func dogrula(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo3() {
	var sayi1 interface{} = "Nurettin"
	dogrula(sayi1)

	var sayi2 interface{} = 5
	dogrula(sayi2)
}

func main() {
	Demo1()
	Demo2()
	Demo3()
}
