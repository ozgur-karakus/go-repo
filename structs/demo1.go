package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"bilgisyar", 22000, "huawei"})
	fmt.Println(product{"bilgisyar", 22000, ""})
}

type product struct {
	name      string
	unitprice float64
	brand     string
}
