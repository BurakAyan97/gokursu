package arrays

import "fmt"

func Demo3() {
	sayilar := [5]int{1, 2, 3, 4, 5}

	fmt.Println(sayilar)

	enBuyuk := sayilar[0]

	for i := 0; i < len(sayilar); i++ {
		if sayilar[i] > enBuyuk {
			enBuyuk = sayilar[i]
		}
	}

	fmt.Println("En büyük : ", enBuyuk)
}
