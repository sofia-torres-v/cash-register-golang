package main

import (
	"fmt"
	"os"

	"cash-register/invoice"
	"cash-register/product"
)

func main() {
	if len(os.Args) < 2 { //true
		fmt.Println("Usage: go run main.go [command]")
		fmt.Println("Commands:")
		fmt.Println("  list   - List available products")
		fmt.Println("  invoice - Create a new invoice")
		return
	}

	command := os.Args[1]

	switch command {
	case "list":
		product.ListProducts()
	case "invoice":
		invoice.CreateInvoice()
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: list, invoice")
	}
}
