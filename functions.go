package main

import "fmt"

func Topla(sayi1 int, sayi2 int) int {
	var toplam = sayi1 + sayi2
	fmt.Println("Sonuç:", toplam)
	return toplam
}

func SelamVer(kullaniciAdi string) {

	fmt.Println("Merhaba Kamil Umut Araz")
}

func DortIslem(sayi1 int, sayi2 int) (int, int, int, float32) {
	toplam := sayi1 + sayi2
	cikarim := sayi1 - sayi2
	carpim := sayi1 * sayi2
	bolum := float32(sayi1) / float32(sayi2)

	return toplam, cikarim, carpim, bolum
}

func ToplaVariadic(sayilar ...int) int {
	toplam := 0
	for i := 0; i < len(sayilar); i++ {
		toplam = toplam + sayilar[i]
	}
	return toplam
}

func main() {
	Topla(3, 4)
	SelamVer("Ahmet")
	fmt.Println(DortIslem(10, 5))
	fmt.Println(ToplaVariadic(1, 2, 3, 4, 5))
}
