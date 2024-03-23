package main

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dunya!"
	i := 1

	for i <= 5 {
		fmt.Println(metin)
		i++
	}

	fmt.Println("Bitti.")
}

func Demo2() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	Demo1()
	Demo2()
}
