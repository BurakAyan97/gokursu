package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Engin"
	fmt.Println(s.HasPrefix(isim, "Eng")) //StartsWith muhabbeti..
	fmt.Println(s.HasSuffix(isim, "in"))  //EndsWith muhabbeti..

	harfler := []string{"e", "n", "g", "i", "n"}
	sonuc := s.Join(harfler, "")
	fmt.Println(s.Join(harfler, "*"))
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "+", -1)) //-1 verirsen hepsini değiştirir. Sayı verirsen o sayı kadar değiştirme yapar. Örneğin 3 verirsen gördüğü ilk 3 * değerini + yapar

	fmt.Println(s.Split(sonuc, ","))
	fmt.Println(s.Repeat(sonuc, 5)) // 5 tane kopya yapar yan yana.

}
