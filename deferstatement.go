package main

import "fmt"

func A() {
	fmt.Println("a fonksiyonu çalisti")
}

func C() {
	fmt.Println("c fonksiyonu çalisti")
}

func D() {
	fmt.Println("d fonksiyonu çalisti")
}

func B() {
	defer A()
	defer C()
	defer D()
	fmt.Println("b fonksiyonu çalisti")
}

func Demo2(sayi int) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		fmt.Println("Çift sayi çalisti.")
		return "Çift sayidir."
	}

	if sayi%2 != 0 {
		fmt.Println("Tek sayi çalisti.")
		return "Tek sayidir."
	}

	return "Belli değil"
}

func Test() {
	sonuc := Demo2(9)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti")
}

type Product struct {
	productName string
	unitPrice   int
}

func (p Product) save() {
	fmt.Println("Ürün kaydedildi : ", p.productName)
	defer Log() //Log'un her zaman çalışmasını istiyorum
}

func Demo3() {
	p := Product{productName: "Laptop", unitPrice: 5000}
	defer p.save()
	p = Product{productName: "Mouse", unitPrice: 45}

	fmt.Println("İşlem başarili")
	fmt.Println(p.productName)
}

func Log() {
	fmt.Println("Loglandi.")
}

func main() {
	B()
	Test()
	Demo3()
}
