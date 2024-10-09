package channels

func CiftSayilar(ciftSayicn chan int) {
	toplam := 0
	for i := 0; i <= 100; i += 2 {
		toplam = toplam + i
	}

	ciftSayicn <- toplam
}

func TekSayilar(tekSayicn chan int) {
	toplam := 0
	for i := 1; i <= 100; i += 2 {
		toplam = toplam + i
	}

	tekSayicn <- toplam
}
