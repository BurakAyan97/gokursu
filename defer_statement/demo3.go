package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Ürün kaydedildi : ", p.productName)
	defer Log()
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 5000}

	// defer p.save() Buraya koyarsan kayıt atarken laptop olarak kayıt atar. Yine en son çalışır ama sınıfın mevcut halini aklında tutar ve onu kayıt eder. Sonradan değişse bile umrunda olmaz

	fmt.Println("İşlem başarılı")

	p = product{productName: "Mouse", unitPrice: 5000}

	fmt.Println(p.productName)

	defer p.save() // En sona koyarsan sınıfların değerleri için daha garanti olur.
}

func Log() {
	fmt.Println("Loglandı")
}
