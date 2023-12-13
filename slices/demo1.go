package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)
	fmt.Println(isimler)
	isimler[0] = "Engin"
	isimler[1] = "dilan"
	isimler[2] = "polat"
	isimler = append(isimler, "dubai") // "append" slices oluşturarak yeni bir ekleme alanı açar. append=eklemek
	fmt.Println(isimler)
	fmt.Println(len(isimler)) // "len" isimler adlı değişkenin kaç elemandan oluştuğunu gösterir.
}
