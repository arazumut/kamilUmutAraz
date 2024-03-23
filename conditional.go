package main

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	var durum bool
	durum = cekilmekIstenen > hesap

	if durum {
		fmt.Println("Hesaptaki para yetersiz.")
	}

	if cekilmekIstenen <= hesap {
		fmt.Println("Paraniz hazirlaniyor...")
		hesap = hesap - cekilmekIstenen
	}

	fmt.Printf("Bitti. Hesaptaki para: %v", hesap)
}

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 1000

	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki para yetersiz.")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Paraniz hazirlaniyor")
		fmt.Println("Dikkat, hesapta para kalmadi")
	} else {
		fmt.Println("Paraniz hazirlaniyor")
	}
}

func Demo3() {
	//üç adet int değişken tanımla
	//ekrana en büyük olanı yazdır.

	var sayi1, sayi2, sayi3 int = 10, 5, 18

	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2
	}
	if sayi3 > enBuyuk {
		enBuyuk = sayi3
	}

	fmt.Printf("En Büyük sayi : %v", enBuyuk)
}

func main() {
	Demo1()
	Demo2()
	Demo3()
}
