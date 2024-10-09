package goroutines

import "fmt"

func CiftSayilar() {
	for i := 0; i <= 10000; i += 2 {
		fmt.Println("Çift sayi : ", i)
	}
}

func TekSayilar() {
	for i := 1; i <= 10000; i += 2 {
		fmt.Println("Tek sayi : ", i)
	}
}
