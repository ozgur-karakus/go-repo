package loops

import "fmt"

func Demo1() {
	var metin string = "cumhuriyet 100 yaşında!"
	i := 1
	for i <= 5 {
		fmt.Println(metin)
		i = i + 1 // i yi vir arttırma yapar buda yukarıdaki i<=5'i sonsuz döngüden kurtarır
	}
	fmt.Println("bitti")
}
