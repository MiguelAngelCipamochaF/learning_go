package main

import "fmt"

func main() {
	palabra := "Palabra"

	fmt.Printf("Cantidad de letras en %v es: %v\n", palabra, len(palabra))

	for i := 0; i < len(palabra); i++ {
		fmt.Printf("%c\n", palabra[i])
	}
}
