package loops

import "fmt"

func Demo4() {
	sayi := 0
	fmt.Println("lütfen bir sayı giriniz.")
	fmt.Scanln(&sayi)
	asal := true
	for i := 2; i < sayi; i++ {
		if sayi%i == 0 {
			asal = false
		}
	}
	if asal == true {
		fmt.Println("asaldır")
	} else {
		fmt.Println("asal değildir")
	}
}
