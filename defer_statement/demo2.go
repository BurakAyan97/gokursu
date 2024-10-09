package defer_statement

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc() // En başa yazılması lazım
	if sayi%2 == 0 {
		return "Çift sayıdır"
	}
	if sayi%2 != 0 {
		return "Tek sayıdır"
	}

	return "Belli değil yaw"
}

func Test() {
	sonuc := Demo2(9)
	fmt.Println(sonuc)
}

func DeferFunc() {
	fmt.Println("Bitti ")
}
