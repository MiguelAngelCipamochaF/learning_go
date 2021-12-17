package main

import "fmt"

func main() {
	horasMes := 172.0
	categoria := "a"

	salario := calcularSalario(horasMes, categoria)

	fmt.Printf("Su salario es: %v\n", salario)
}

func calcularSalario(horasMes float64, categoria string) float64 {
	salario := 0.0
	switch categoria {
	case "c":
		salario = 1000 * horasMes
	case "b":
		salario = 1500 * horasMes
		adicional := salario * 20 / 100
		salario += adicional
	case "a":
		salario = 3000 * horasMes
		adicional := salario * 50 / 100
		salario += adicional
	default:
		return 0
	}

	return salario
}
