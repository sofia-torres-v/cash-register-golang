package main

import (
	"fmt"
	"os"

	"cash-register/invoice"
	"cash-register/product"
	"encoding/json"
)

func readProductJson() ([]product.Product, error) {
	data, err := os.ReadFile("product.json")
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)

		return []product.Product{}, err
	}

	var products []product.Product

	err = json.Unmarshal(data, &products)
	if err != nil {
		fmt.Println("Error al decodificar JSON:", err)
		return []product.Product{}, err
	}

	// fmt.Println(string(data))
	// castear-convertir byte a string o viceversa
	return products, nil
}

func main() {
	if len(os.Args) < 2 { //true
		fmt.Println("Usar: go run main.go [commando]")
		fmt.Println("Comandos:")
		fmt.Println("  list   - Listar productos disponibles")
		fmt.Println("  invoice - Crear una nueva factura")
		return
	}

	command := os.Args[1]

	result, err := readProductJson()
	if err != nil {
		fmt.Println("Error al cargar los productos")
		return
	}

	switch command {
	case "list":
		product.ListProducts(result)
	case "invoice":
		invoice.CreateInvoice(result)
	default:
		fmt.Println("Comando desconocido:", command)
		fmt.Println("Comandos disponibles: lista, factura")
	}
}
