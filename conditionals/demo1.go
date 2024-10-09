package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Print("Hesaptaki para yetersizdir.")
	}

	if cekilmekIstenen <= hesap {
		fmt.Print("Paranız hazırlanıyor.")
		hesap = hesap - cekilmekIstenen
	}
	fmt.Println("Bitti. Hesaptaki para : " + fmt.Sprintf("%v", hesap)) //Eğer sayı ondalıksız ise => %f ondalıklı olarak gösterir floatı. %v ise direkt tam sayı gösterir. Değil ise %f 6 basamak gösteriyor %v ise 15 basamak.
}
