package invoice

import (
	"cash-register/product"
	"fmt"
)

// limpia la pantalla usando códigos de escape ANSI.
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func CreateInvoice(products []product.Product) {
	var customerName string
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanln(&customerName)

	var invoiceProducts []product.Product
	var total float64

	fmt.Printf("\nProductos de %s:\n", customerName) // Mostrar encabezado una vez

	for {
		var productID int

		fmt.Print("Ingresa ID de producto (o 0 para finalizar): ")
		fmt.Scanln(&productID)

		if productID == 0 {
			break
		}

		selectedProduct, err := product.FindProductByID(productID, products)
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
		fmt.Printf("Productos de %s:\n", customerName)

		for _, product := range invoiceProducts {
			fmt.Printf("%d   %s   $%.2f \n", product.Quantity, product.Name, product.Price)
		}
	}

	// Validación cuando la lista esta vacía
	if len(invoiceProducts) == 0 {
		fmt.Println("No se agregaron productos a la lista")
		return
	}

	// Mostrar el total final de la factura al salir
	fmt.Printf("\nFactura final para %s:\n", customerName)
	for _, product := range invoiceProducts {
		fmt.Printf("%d   %s   $%.2f each\n", product.Quantity, product.Name, product.Price)
	}
	fmt.Printf("Total: $%.2f\n", total)
}
