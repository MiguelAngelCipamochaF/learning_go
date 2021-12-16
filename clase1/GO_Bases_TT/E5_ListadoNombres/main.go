package main

import "fmt"

func main() {
	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro",
		"Axel", "Alez", "Dolores", "Federico", "Hernan", "Leandro", "Eduardo", "Duvraschka"}

	for _, e := range estudiantes {
		fmt.Println(e)
	}

	estudiantes = append(estudiantes, "Gabriela")
	fmt.Println("\nNueva estudiante")
	for _, e := range estudiantes {
		fmt.Println(e)
	}
}
