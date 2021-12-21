package main

import (
	"errors"
	"fmt"
)

func main() {
	horasMes := 45
	valorHora := 234
	mejorSalario := -160000
	mesesTrabajados := 4

	salario, err := salarioMensual(horasMes, valorHora)

	if err != nil {
		fmt.Println(err)
	}

	aguinaldo, err := medioAguinaldo(mejorSalario, mesesTrabajados)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Tu salario mensual es: %v\n", salario)
	fmt.Printf("Tu medio aguinaldo es: %v\n", aguinaldo)

}

func salarioMensual(horasMes int, valorHora int) (int, error) {

	salario := horasMes * valorHora

	if salario >= 150000 {
		descuento := salario * 10 / 100
		salario -= descuento
	}

	if horasMes < 80 || horasMes < 0 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	return salario, nil
}

func medioAguinaldo(mejorSalario int, mesesTrabajados int) (int, error) {
	aguinaldo := (mejorSalario / 12) * mesesTrabajados

	if mejorSalario < 0 || mesesTrabajados < 0 {
		return 0, fmt.Errorf("error: numero negativo ingresado")
	}

	return aguinaldo, nil
}
