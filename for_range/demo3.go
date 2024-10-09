package for_range

import "fmt"

func Demo3() {
	sozluk := map[string]string{"book": "Kitap", "table": "masa"}

	for k, v := range sozluk {
		fmt.Print(k)
		fmt.Println(v)
	}
}
