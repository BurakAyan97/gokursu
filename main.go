package main

import (
	"fmt"
	"golesson/project"
)

func main() {
	// variables.Demo1()
	// fmt.Print()
	// conditionals.Demo2()
	// conditionals.DemoWorkShop()
	// loops.Demo2()
	// loops.Workshop1()
	// arrays.Demo3()
	// slices.Demo2()
	// var sonuc = functions.Topla(2, 3)
	// fmt.Println(sonuc * 10)
	// fmt.Println(functions.SelamVer("Burak"))

	// data1, data2, data3, data4 := functions.DortIslem(24, 21)
	// // data1, data2, data3, _ := functions.DortIslem(24, 21) İstemiysak bir sonucu _ koyalım

	// fmt.Println("Toplam : ", data1)
	// fmt.Println("Fark : ", data2)
	// fmt.Println("Carpim : ", data3)
	// fmt.Println("Bolum : ", data4)
	// fmt.Println(functions.ToplaVariadic(1, 2, 3, 4, 5, 7, 2321))
	// fmt.Println(functions.ToplaVariadic(5, 7))
	// fmt.Println(functions.ToplaVariadic())

	// sayilar := []int{4, 6, 8, 1}
	// fmt.Println(functions.ToplaVariadic(sayilar...))
	// maps.Demo1()
	// for_range.Demo3()

	// sayi := 20
	// pointers.Demo1(&sayi) // pointerlı veriyorsan & ile karşıla fonkisyonda
	// fmt.Println("Maindeki sayı : ", sayi)

	// sayilar := []int{1, 2, 3}

	// pointers.Demo2(sayilar)
	// fmt.Println("Maindeki sayılar :", sayilar[0])
	// structs.Demo2()
	// ciftSayiCn := make(chan int)
	// tekSayiCn := make(chan int)

	// go channels.CiftSayilar(ciftSayiCn)
	// go channels.TekSayilar(tekSayiCn)

	// // carpim := ciftSayiCn * tekSayiCn -> yapamazsın çünkü bunlar sayısal değerler değil. Sadece int taşıyan bir channel

	// ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
	// carpim := ciftSayiToplam * tekSayiToplam
	// fmt.Println("Çarpım : ", carpim)

	// interfaces.Demo2()

	// defer_statement.Demo3()

	// error_handling.Demo1()

	// interfaces.Demo3()

	// fmt.Println(error_handling.TahminEt2(102))

	// string_functions.Demo2()

	// restful.Demo2()
	project.AddProduct()
	products, _ := project.GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}
}
