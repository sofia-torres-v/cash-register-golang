package invoice

import (
	"cash-register/product"
	"fmt"
)

// ClearScreen limpia la pantalla usando códigos de escape ANSI.
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func CreateInvoice() {
	var customerName string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&customerName)

	var invoiceProducts []product.Product
	var total float64

	fmt.Printf("\nProductos de %s:\n", customerName) // Mostrar encabezado una vez

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

		// Verifica si el producto ya está en la factura
		var productExist bool
		for i := range invoiceProducts {
			if invoiceProducts[i].ID == selectedProduct.ID {
				invoiceProducts[i].Quantity++
				total += selectedProduct.Price
				productExist = true
				break
			}
		}

		// Si el producto no existe, lo agregamos a la lista
		if !productExist {
			selectedProduct.Quantity = 1 // Asigna una cantidad inicial de 1
			invoiceProducts = append(invoiceProducts, selectedProduct)
			total += selectedProduct.Price
		}

		// Limpia la pantalla y muestra la lista actualizada
		ClearScreen()
		fmt.Printf("Productos de %s (actualizado):\n", customerName)

		for _, product := range invoiceProducts {
			fmt.Printf("%d   %s   $%.2f each\n", product.Quantity, product.Name, product.Price)
		}
		// fmt.Printf("Current total: $%.2f\n", total)
	}

	// Mostrar el total final de la factura al salir
	fmt.Printf("\nFactura final para %s:\n", customerName)
	for _, product := range invoiceProducts {
		fmt.Printf("%d   %s   $%.2f each\n", product.Quantity, product.Name, product.Price)
	}
	fmt.Printf("Total: $%.2f\n", total)
}
