package main

import (
	"fmt"
	"time"
)

func CiftSayilar() {
	for i := 0; i < 10; i += 2 {
		fmt.Println("Çift sayi: ", i)
		time.Sleep(1 * time.Second)
	}
}

func TekSayilar() {
	for i := 1; i < 10; i += 2 {
		fmt.Println("Tek sayi: ", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go CiftSayilar()
	go TekSayilar()

	// Ana işlem bitmeden programın sonlanmaması için beklemek gerekir.
	time.Sleep(11 * time.Second)
}
