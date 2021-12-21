package main

import (
	"fmt"
)

func main() {
	salary := 140000

	if salary < 150000 {
		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary)
		fmt.Println(err)
		return
	}

	fmt.Printf("Debe pagar impuesto\n")
}
