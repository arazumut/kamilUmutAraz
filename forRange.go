package main

import "fmt"

func Demo1() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"}
	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}

	for _, sehir := range sehirler {
		fmt.Println(sehir)
	}

	for i, sehir := range sehirler {
		fmt.Print(i, " ")
		fmt.Println(sehir)
	}
}

func Demo2() {
	sayilar := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toplam := 0

	for _, sayi := range sayilar {
		if sayi%2 == 1 {
			toplam += sayi
			fmt.Println(sayi)
		}
	}
	fmt.Println(toplam)
}

func Demo3() {
	sozluk := map[string]string{"book": "kitap", "table": "masa"}

	for k := range sozluk {
		fmt.Println(k)
	}
}

func main() {
	Demo1()
	Demo2()
	Demo3()
}
