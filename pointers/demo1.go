package pointers

import "fmt"

func Demo1(sayi *int) {
	*sayi = *sayi + 1                           //Adres olan int ile değer olan int farklıdır. İkisine de vermelisin
	fmt.Println("Demo1 içindeki sayı :", *sayi) //sayi yazarsan bellekteki yerini gösterir. *sayi yaparsan değeri görürsün
}
