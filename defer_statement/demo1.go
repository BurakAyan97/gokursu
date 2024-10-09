package defer_statement

import "fmt"

func A() {
	fmt.Println("a fonksiyonu çalıştı")
}
func C() {
	fmt.Println("c fonksiyonu çalıştı")
}
func D() {
	fmt.Println("d fonksiyonu çalıştı")
}
func B() {
	defer A() // bu kodu nereye koyarsam koyayım en son çalışacaktır.
	defer C() // Önce C sonra A çalışır. LIFO olarak çalışır
	defer D() // deferler fonksiyon bittikten sonra çalışırlar. Fonksiyon içinde hata olsa bile deferler her zaman çalışır. Örn : loglama işlemleri
	fmt.Println("b fonksiyonu çalıştı")
}
