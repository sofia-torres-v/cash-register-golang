package product

import (
	"fmt"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

var products = []Product{
	{ID: 1, Name: "Apple", Price: 0.50},
	{ID: 2, Name: "Grape", Price: 0.30},
	{ID: 3, Name: "Onion", Price: 1.20},
	{ID: 4, Name: "Bread", Price: 0.80},
	{ID: 5, Name: "Sugar", Price: 0.80},
}

// algoritmo insertion sort
func SortList(products []Product) []Product {
	for i := 1; i < len(products); i++ {
		elementOrd := products[i]
		j := i - 1

		for j >= 0 && elementOrd.Name < products[j].Name {
			products[j+1] = products[j]
			j = j - 1
		}
		products[j+1] = elementOrd
	}
	return products
}

func ListProducts() {
	fmt.Println("Available Products:")
	for _, product := range SortList(products) {
		fmt.Printf("ID: %d, Name: %s, Price: $%.2f\n", product.ID, product.Name, product.Price)
	}
}

func FindProductByID(id int) (Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, fmt.Errorf("product with ID %d not found", id)
}
