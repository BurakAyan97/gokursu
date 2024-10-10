package project

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryID  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Product, error) {
	response, err := http.Get("http://localhost:3000/products")

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	var products []Product

	json.Unmarshal(bodyBytes, &products)
	return products, nil
}

func AddProduct() ([]Product, error) {
	product := Product{4, "Xbox", 1, 22500}

	jsonProduct, err := json.Marshal(product)

	if err != nil {
		return nil, err
	}

	response, err := http.Post("http://localhost:3000/products", "application/json:charset=utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	var productResponse []Product

	json.Unmarshal(bodyBytes, &productResponse)

	return productResponse, nil

}
