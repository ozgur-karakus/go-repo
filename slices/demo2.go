package slices

import "fmt"

func Demo2() {
	sehirler := []string{"ankara", "istanbul", "izmir"} // "string{}"" ifadesi doğrudan slices ekleme için kullanılır.
	fmt.Println(sehirler)
	sehirlerkopya := make([]string, len(sehirler))
	fmt.Println(sehirlerkopya)
	copy(sehirlerkopya, sehirler) //copy değişken değeri kopyalamak için kullanılır.
	fmt.Println(sehirlerkopya)
	sehirler = append(sehirler, "antalya") //append = ekleme metodu.
	fmt.Println(sehirler)
	fmt.Println(sehirlerkopya)

	fmt.Println(sehirler[1:3]) // 1 dahil 3 e kadar olan değerler yazılır
	fmt.Println(sehirler[1:])  // 1 dahil son değere kadar olan değerler yazdırılır
	fmt.Println(sehirler[:2])  // 2 ye kadar olan değerler yazdırılır
}
