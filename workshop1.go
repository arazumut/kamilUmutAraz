package main

import "fmt"

func Demo3() {
	aklimdakiSayi := 80
	tahminEdilenSayi := 0
	tahminSayisi := 0

	fmt.Println("Bir sayi tahmin ediniz")
	fmt.Scanln(&tahminEdilenSayi)
	tahminSayisi = tahminSayisi + 1

	for aklimdakiSayi != tahminEdilenSayi {
		if tahminEdilenSayi < aklimdakiSayi {
			fmt.Println("Daha yüksek tahmin ediniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1

		}
		if tahminEdilenSayi > aklimdakiSayi {
			fmt.Println("Daha düşük tahmin ediniz")
			fmt.Scanln(&tahminEdilenSayi)
			tahminSayisi = tahminSayisi + 1
		}
	}

	basariDurumu := ""
	if tahminSayisi > 0 && tahminSayisi <= 3 {
		basariDurumu = "Süper"
	} else if tahminSayisi <= 6 {
		basariDurumu = "Güzel"
	} else {
		basariDurumu = "Fena değil"
	}

	fmt.Printf("Tebrikler bildiniz. %v tahmin: %v", tahminSayisi, basariDurumu)
}

func Demo4() {
	sayi := 0
	fmt.Println("Bir sayi giriniz...")
	fmt.Scanln(&sayi)

	asalMi := true
	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asalMi = false
		}
	}

	if asalMi == true {
		fmt.Println("Asaldir.")
	} else {
		fmt.Println("Asal değildir.")
	}
}

func Demo5() {
	sayi1 := 220
	sayi2 := 284
	toplam1 := 0
	toplam2 := 0

	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			toplam1 += i
		}
	}
	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			toplam2 += i
		}
	}
	if toplam1 == sayi2 && toplam2 == sayi1 {
		fmt.Printf("%v ve %v arkadaş sayilardir...", sayi1, sayi2)
	} else {
		fmt.Printf("%v ve %v arkadaş sayi değillerdir...", sayi1, sayi2)
	}
}

func main() {
	Demo3()
	Demo4()
	Demo5()
}
