package main

import "fmt"

func Demo1() {
	// key - value
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman sayisi : ", len(sozluk))
	fmt.Println("Sozluk : ", sozluk)
	delete(sozluk, "book")
	fmt.Println("Eleman sayisi : ", len(sozluk))
	fmt.Println("Sozluk : ", sozluk)

	deger, varMi := sozluk["table"]
	fmt.Println(deger)
	fmt.Println("Listede olma durumu : ", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(sozluk2)
}

func main() {
	Demo1()
}
