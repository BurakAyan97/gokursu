package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)

	sozluk["table"] = "masa"
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman Sayısı : ", len(sozluk))
	delete(sozluk, "book")
	fmt.Println("Eleman Sayısı : ", len(sozluk))

	deger, varMi := sozluk["table"] // 2 sonuç dönebiliyor
	fmt.Println("Listede olma durumu : ", varMi)
	fmt.Println("Türkçesi: ", deger)

	sozluk2 := map[string]string{"glass": "bardak", "woman": "kadın"}

	fmt.Println(sozluk2)

}
