package main

import "fmt"

func Demo1() {
	var sayilar [5]int
	sayilar[2] = 50
	fmt.Println(sayilar[2])
}

func Demo2() {
	var sehirler [5]string
	sehirler[0] = "Ankara"
	sehirler[1] = "İstanbul"
	sehirler[2] = "İzmir"
	sehirler[3] = "Adana"
	sehirler[4] = "Diyarbakir"

	fmt.Println(sehirler)

	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}
}

func Demo3() {
	sayilar := [5]int{20, 30, 50, 10, 2}

	enBuyuk := sayilar[0]

	for i := 0; i < len(sayilar); i++ {
		if sayilar[i] > enBuyuk {
			enBuyuk = sayilar[i]
		}
	}
	fmt.Println("En büyük: ", enBuyuk)
}

func Demo4() {
	var sayilar [2][3]int

	sayilar[0][0] = 1
	sayilar[0][1] = 3
	sayilar[0][2] = 5
	sayilar[1][0] = 2
	sayilar[1][1] = 4
	sayilar[1][2] = 6

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(sayilar[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func main() {
	Demo1()
	Demo2()
	Demo3()
	Demo4()
}
