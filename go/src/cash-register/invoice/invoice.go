package invoice

import (
	"fmt"

	"cash-register/product"
)

func CreateInvoice() {
	var customerName string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&customerName)

	var invoiceProducts []product.Product
	var total float64

	for {
		var productID int
		fmt.Print("Enter product ID (or 0 to finish): ")
		fmt.Scanln(&productID)

		if productID == 0 {
			break
		}

		selectedProduct, err := product.FindProductByID(productID)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		invoiceProducts = append(invoiceProducts, selectedProduct)
		// total = total + product.Price
		total += selectedProduct.Price

		// Mostrar los productos mientars se va agregando
		fmt.Printf("\nInvoice for %s:\n", customerName)
		for _, product := range invoiceProducts {
			fmt.Printf("Product: %s, Price: $%.2f\n", product.Name, product.Price)
		}
	}

	fmt.Printf("Total: $%.2f\n", total)
}
