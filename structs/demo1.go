package structs

import "fmt"

func Demo1() {

	fmt.Println(product{"Bilgisayar", 245, "Xyz", 2})
	fmt.Println(product{"Bilgisayar", 245, "", 2})
	fmt.Println(product{name: "Bilgisayar", unitPrice: 245})
}

type product struct {
	name         string
	unitPrice    float64
	brand        string
	discountRate int
}
