package main

import (
	"bufio"
	"fmt"
	"os"
)

//! Produced by Kamil Umut Araz instagram : K.umutarazz  github : arazumut

var stack []string
var queue []string

func main() {
	for {
		fmt.Println("\nAna Menü:")
		fmt.Println("1. Stack İşlemleri")
		fmt.Println("2. Queue İşlemleri")
		fmt.Println("3. Çıkış")

		secim := input("Seçiminizi giriniz: ")

		switch secim {
		case "1":
			stackMenu()
		case "2":
			queueMenu()
		case "3":
			fmt.Println("Çıkış yapılıyor... Bir dahaki sefere görüşmek üzere :)")
			return
		default:
			fmt.Println("Geçersiz seçim!")
		}
	}
}

func stackMenu() {
	for {
		fmt.Println("\nStack İşlemleri Menüsü:")
		fmt.Println("1. Ekle")
		fmt.Println("2. Bul ve Sil")
		fmt.Println("3. Bul ve Göster")
		fmt.Println("4. Tümünü Listele")
		fmt.Println("5. Ana Menüye Dön")

		secim := input("Seçiminizi giriniz: ")

		switch secim {
		case "1":
			stackEkle()
		case "2":
			stackBulSil()
		case "3":
			stackBulGoster()
		case "4":
			stackListele()
		case "5":
			return
		default:
			fmt.Println("Geçersiz seçim!")
		}
	}
}

func stackEkle() {
	deger := input("Eklemek istediğiniz değeri giriniz: ")
	stack = append(stack, deger)
	fmt.Printf("%s değeri stack'e eklendi.\n", deger)
}

func stackBulSil() {
	bulunacak := input("Bulmak istediğiniz değeri giriniz: ")
	index := find(stack, bulunacak)
	if index != -1 {
		stack = append(stack[:index], stack[index+1:]...)
		fmt.Printf("%s değeri stack'ten silindi.\n", bulunacak)
	} else {
		fmt.Printf("%s değeri stack'te bulunamadı.\n", bulunacak)
	}
}

func stackBulGoster() {
	bulunacak := input("Bulmak istediğiniz değeri giriniz: ")
	if contains(stack, bulunacak) {
		fmt.Printf("%s değeri stack'te bulunmaktadır.\n", bulunacak)
	} else {
		fmt.Printf("%s değeri stack'te bulunamadı.\n", bulunacak)
	}
}

func stackListele() {
	if len(stack) > 0 {
		fmt.Println("Stack'teki elemanlar:")
		for _, deger := range stack {
			fmt.Println(deger)
		}
	} else {
		fmt.Println("Stack boş.")
	}
}

func queueMenu() {
	for {
		fmt.Println("\nQueue İşlemleri Menüsü:")
		fmt.Println("1. Ekle")
		fmt.Println("2. Kaldır")
		fmt.Println("3. En Önde Ne Var?")
		fmt.Println("4. Tümünü Listele")
		fmt.Println("5. Ana Menüye Dön")

		secim := input("Seçiminizi giriniz: ")

		switch secim {
		case "1":
			queueEkle()
		case "2":
			queueKaldir()
		case "3":
			queueOnEleman()
		case "4":
			queueListele()
		case "5":
			return
		default:
			fmt.Println("Geçersiz seçim!")
		}
	}
}

func queueEkle() {
	deger := input("Eklemek istediğiniz değeri giriniz: ")
	queue = append(queue, deger)
	fmt.Printf("%s değeri queue'ya eklendi.\n", deger)
}

func queueKaldir() {
	if len(queue) > 0 {
		kaldirilan := queue[0]
		queue = queue[1:]
		fmt.Printf("%s değeri queue'dan silindi.\n", kaldirilan)
	} else {
		fmt.Println("Queue boş.")
	}
}

func queueOnEleman() {
	if len(queue) > 0 {
		fmt.Printf("Queue'nun en önündeki eleman: %s\n", queue[0])
	} else {
		fmt.Println("Queue boş.")
	}
}

func queueListele() {
	if len(queue) > 0 {
		fmt.Println("Queue'daki elemanlar:")
		for _, deger := range queue {
			fmt.Println(deger)
		}
	} else {
		fmt.Println("Queue boş.")
	}
}

func input(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func find(slice []string, item string) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}
