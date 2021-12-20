package main

import (
	"fmt"
	"os"
)

func main() {
	productos := map[string][]string{
		"Manzana": {"001", "25000", "2"},
		"Durazno": {"002", "30000", "5"},
		"Naranja": {"003", "27000", "4"},
		"Pera":    {"004", "20000", "3"},
	}

	file, err := os.Create("productos.txt")

	if err != nil {
		fmt.Printf("Error creating product.txt")
	}

	defer file.Close()

	for _, product := range productos {
		for _, v := range product {
			file.WriteString(v + ";")
		}
		file.WriteString("\n")
	}
}
