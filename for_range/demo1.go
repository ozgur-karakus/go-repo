package for_range

import "fmt"

func Demo1() {
	sehirler := []string{"ankara", "istanbul", "izmir"}
	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}
	for i, sehir := range sehirler {
		fmt.Print(i)
		fmt.Println(sehir)
	}
}
