package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)
	sozluk["table"] = "masa"
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman sayısı : ", len(sozluk)) // "len" kullanıcı için gereken eleman sayısı girdi sayısı gibi olan değerleri verir
	fmt.Println("sözlük : ", sozluk)
	delete(sozluk, "book")
	fmt.Println("Eleman sayısı : ", len(sozluk))
	fmt.Println("Sözlük : ", sozluk)
	deger, varmi := sozluk["book"]
	fmt.Println(deger)
	fmt.Println("listede olma durumu : ", varmi)
	sozluk2 := map[string]string{"glass": "bardak", "tea": "çay"}
	fmt.Println(sozluk2)
}
