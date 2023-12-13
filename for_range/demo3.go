package for_range

import "fmt"

func Demo3() {
	sozluk := map[string]string{"book": "kitap", "table": "masa"}
	for k, v := range sozluk { // burada ilk değişken map de tanımlanan ":" ,fadesinden önceki değeri yazdırırken 2. olan v değeri de ": ifadesinden sonraki değeri yazdırarak ekrana veriyor"
		fmt.Println(k)
		fmt.Println(v)
	}
}
