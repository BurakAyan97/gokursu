package interfaces

import "fmt"

//Type asseriton muhabbeti c# daki object gibidir. Go dilinde interfaceler her türlü değişkeni karşılayabilirler. casting işlemi yapmadıkları için daha performanslıdır
func Dogrula(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo3() {
	var sayi interface{} = "Engin"
	Dogrula(sayi)

	var sayi2 interface{} = 5
	Dogrula(sayi2)
}
