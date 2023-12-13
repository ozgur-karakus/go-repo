package goroutines

import "fmt"

func Ciftsayilar() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("çift sayılar : ", i)
	}
}

func Teksayilar() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("tek sayılar :", i)
	}
}
