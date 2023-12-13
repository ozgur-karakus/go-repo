package loops

import "fmt"

func Demo3() {

	as := 95
	tahmin := 0
	tahminsayisi := 0

	fmt.Println("Lütfen 0-100 dahil olacak şekilde verilen aralıkta bir sayı tahmin ediniz.")
	fmt.Scanln(&tahmin)
	tahminsayisi = tahminsayisi + 1
	for as != tahmin {
		if tahmin < as {
			fmt.Println("yukarı")
			fmt.Scanln(&tahmin)
			tahminsayisi = tahminsayisi + 1
		}
		if tahmin > as {
			fmt.Println("aşağı")
			fmt.Scanln(&tahmin)
			tahminsayisi = tahminsayisi + 1
		}
		if tahmin == as {
			fmt.Println("Bravo")
		}
	}
	basari := ""
	if tahminsayisi > 0 && tahminsayisi <= 3 {
		basari = "SÜPER"
	} else if tahminsayisi > 3 && tahminsayisi <= 6 {
		basari = "FENA DEĞİL"
	} else {
		basari = "ZAYOTEM HAYIRLI OLSUN"
	}
	fmt.Printf("tahmin sayısı : %v : %v", tahminsayisi, basari)
}
