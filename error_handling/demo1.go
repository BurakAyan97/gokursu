package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt")
	//Eğer err değerine nil atanmıyorsa hata vardır. Eğer error nil ise bir hata yoktur demektir. Hata yakalama mantığı go da bu şekilde çalışıyor.
	if err != nil {
		//path error ise pErr'e atanıyor değilse ok true false olarak atanıyor
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı :", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadı")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}
