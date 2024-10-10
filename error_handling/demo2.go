package error_handling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	aklimdakiSayi := 50
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1 ile 100 arasında bir sayı giriniz.")
	}
	if tahmin > aklimdakiSayi {
		return "Daha küçük bir sayı gir", nil
	}

	if tahmin < aklimdakiSayi {
		return "Daha büyük bir sayı gir", nil
	}

	return "Bildiniz", nil //ikinci değer olarak error döneceğim için hata yoksa = nil demekti. O yüzden nil dönüyorum.
}

func Demo2() {
	//Outputta nil'i görmek istemiyosam sadece mesajı göstermek için ikinci parametreyi _ olarak koydum.
	mesaj, _ := TahminEt(50)
	fmt.Println(mesaj)
	fmt.Println(TahminEt(101))
	fmt.Println(TahminEt(-100))
}
