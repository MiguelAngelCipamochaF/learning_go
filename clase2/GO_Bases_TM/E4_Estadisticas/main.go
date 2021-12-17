package main

import (
	"errors"
	"fmt"
)

func main() {

	const (
		minimo   = "minimo"
		maximo   = "maximo"
		promedio = "promedio"
	)

	min, e := tipoCalculo(minimo)
	if e != nil {
		fmt.Println(e)
	}
	max, e := tipoCalculo(maximo)
	if e != nil {
		fmt.Println(e)
	}
	prom, e := tipoCalculo(promedio)
	if e != nil {
		fmt.Println(e)
	}

	valorMinimo := min(4, 5, 6, 3, 8, 7, 6, 9, 2, 10)
	valorMaximo := max(4, 5, 6, 3, 8, 7, 6, 9, 2, 10)
	valorPromedio := prom(4, 5, 6, 3, 8, 7, 6, 9, 2, 10)

	fmt.Printf("Valor minimo es: %v\n", valorMinimo)
	fmt.Printf("Valor maximo es: %v\n", valorMaximo)
	fmt.Printf("Valor promedio es: %v\n", valorPromedio)
}

func hallarMinimo(nums ...int) int {
	temp := 100
	for _, i := range nums {
		if temp >= i {
			temp = i
		}
	}
	return temp
}

func hallarMaximo(nums ...int) int {
	temp := 0
	for _, i := range nums {
		if temp <= i {
			temp = i
		}
	}
	return temp
}

func hallarPromedio(nums ...int) int {
	cantidadNums := 0
	promedio := 0
	for _, i := range nums {
		cantidadNums++
		promedio += i
	}

	return promedio / cantidadNums
}

func tipoCalculo(tipo string) (func(nums ...int) int, error) {
	switch tipo {
	case "minimo":
		return hallarMinimo, nil
	case "maximo":
		return hallarMaximo, nil
	case "promedio":
		return hallarPromedio, nil
	default:
		return nil, errors.New("error de tipo de calculo")
	}
}
