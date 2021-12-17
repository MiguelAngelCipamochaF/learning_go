package main

import "fmt"

func main() {
	sueldo := 150000.0

	impuesto, descuento := impuestoSalario(sueldo)
	fmt.Printf("Impuestos: %v\n", descuento)
	fmt.Printf("Tu sueldo total es: %v\n", impuesto)
}

func impuestoSalario(sueldo float64) (float64, float64) {
	descuento := 0.0
	if sueldo >= 50000.0 {
		descuento = sueldo * 17 / 100.0
		if sueldo >= 150000.0 {
			masDescuento := sueldo * 10 / 100.0
			descuento += masDescuento
		}
	}

	return sueldo - descuento, descuento
}
