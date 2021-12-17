package main

import (
	"errors"
	"fmt"
)

func main() {
	promedio, e := promedioCalificaciones(6, 7, 5, 6, 8, 7, 8, 9, 10, 9, 8, 4, 3, 9, 9, 9, 9, 10, 10, 10)

	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Printf("El promedio de calificaciones del colegio es: %v\n", promedio)

}

func promedioCalificaciones(calificaciones ...int) (int, error) {
	promedio := 0
	cantidadAlumnos := 0
	for _, c := range calificaciones {
		cantidadAlumnos++
		if c < 0 {
			return 0, errors.New("error: Hay una calificacion negativa")
		}
		promedio += c
	}

	promedio /= cantidadAlumnos

	return promedio, nil
}
