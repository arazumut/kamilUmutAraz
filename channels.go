package main

import "fmt"

func CiftSayilar(ciftSayiCn chan int) {
	toplam := 0
	for i := 0; i <= 10; i += 2 {
		toplam = toplam + i
	}
	ciftSayiCn <- toplam
}

func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i <= 10; i += 2 {
		toplam = toplam + i
	}
	tekSayiCn <- toplam
}

func main() {
	ciftSayiCn := make(chan int)
	tekSayiCn := make(chan int)

	go CiftSayilar(ciftSayiCn)
	go TekSayilar(tekSayiCn)

	ciftToplam, tekToplam := <-ciftSayiCn, <-tekSayiCn

	fmt.Println("Çift sayıların toplamı:", ciftToplam)
	fmt.Println("Tek sayıların toplamı:", tekToplam)
}
