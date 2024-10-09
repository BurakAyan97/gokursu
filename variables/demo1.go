package variables

import "fmt"

func Demo1() {
	var metin string = "Selamlar mq"
	fmt.Println(metin)

	var kredi int = 24
	fmt.Println(kredi * 5)

	var kdv2 float32 = 24.123
	fmt.Println(kdv2 * 2)

	kd3 := 24
	fmt.Println(kd3)
	fmt.Printf("veri tipi : %T", kd3) //printf formatlı yazım yapıyor bizdeki ${} muhabbeti gibi.

	var durum bool
	fmt.Println(durum)

	var metin1 string = "Engin"
	var metin2 string = "Engin"

	durum = metin1 != metin2
	fmt.Println(durum)
	fmt.Println(!durum)
}
