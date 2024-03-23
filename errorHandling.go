package main

import (
	"errors"
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo11.txt")
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadi", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadi.")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}

func TahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 100

	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasinda bir sayi giriniz.")
	}

	if tahmin > aklimdakiSayi {
		return "Daha küçük bir sayi giriniz.", nil
	}

	if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayi giriniz.", nil
	}

	return "Bildiniz..", nil
}

func Demo2() {
	mesaj, hata := TahminEt(50)
	fmt.Println(mesaj)
	fmt.Println(hata)
}

type borderException struct {
	parameter int
	message   string
}

func (b *borderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func TahminEt2(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &borderException{parameter: tahmin, message: "Sinirlarin dişinda kaldi."}
	}
	return "Bildiniz", nil
}

func main() {
	Demo1()
	Demo2()
}
