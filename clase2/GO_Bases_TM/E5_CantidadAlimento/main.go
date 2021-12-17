package main

import (
	"errors"
	"fmt"
)

func main() {

	const (
		perro     = "perro"
		gato      = "gato"
		hamster   = "hamster"
		tarantula = "tarantula"
	)

	var cantidadTotal float64 = 0.0

	funPerro, err := Animal(perro)
	if err != nil {
		fmt.Println(err)
	}
	funGato, err := Animal(gato)
	if err != nil {
		fmt.Println(err)
	}
	funHamster, err := Animal(hamster)
	if err != nil {
		fmt.Println(err)
	}
	funTarantula, err := Animal(tarantula)
	if err != nil {
		fmt.Println(err)
	}

	alimentoPerro := funPerro(10)
	alimentoGato := funGato(6)
	alimentoHamster := funHamster(10)
	alimentoTarantula := funTarantula(20)

	cantidadTotal = alimentoPerro + alimentoGato + alimentoHamster + alimentoTarantula

	fmt.Printf("Alimento necesario para los perros: %vkg\n", alimentoPerro)
	fmt.Printf("Alimento necesario para los gatos: %vkg\n", alimentoGato)
	fmt.Printf("Alimento necesario para los hamsters: %vkg\n", alimentoHamster)
	fmt.Printf("Alimento necesario para las tarantulas: %vkg\n", alimentoTarantula)
	fmt.Printf("Alimento total necesario: %vkg\n", cantidadTotal)

}

func Animal(tipoAnimal string) (func(cantidadAnimal int) float64, error) {
	switch tipoAnimal {
	case "perro":
		return func(cantidadAnimal int) float64 {
			kgAlimento := 10.0
			cantidadAlimento := kgAlimento * float64(cantidadAnimal)
			return cantidadAlimento
		}, nil
	case "gato":
		return func(cantidadAnimal int) float64 {
			kgAlimento := 5.0
			cantidadAlimento := kgAlimento * float64(cantidadAnimal)
			return cantidadAlimento
		}, nil
	case "hamster":
		return func(cantidadAnimal int) float64 {
			kgAlimento := 0.25
			cantidadAlimento := kgAlimento * float64(cantidadAnimal)
			return cantidadAlimento
		}, nil
	case "tarantula":
		return func(cantidadAnimal int) float64 {
			kgAlimento := 0.15
			cantidadAlimento := kgAlimento * float64(cantidadAnimal)
			return cantidadAlimento
		}, nil
	default:
		return nil, errors.New("error: no se encuentra el tipo de animal")
	}
}
