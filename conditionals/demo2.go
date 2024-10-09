package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900

	if cekilmekIstenen > hesap {
		fmt.Print("Hesaptaki para yetersizdir.")
	} else if cekilmekIstenen == hesap {
		fmt.Print("Paranız hazırlanıyor.")
	} else {
		fmt.Print("Paranız hazırlanıyor.")
	}
}
