package main

import "fmt"

type Matrix struct {
	x           uint16
	y           uint16
	Values      [3][3]float64
	cuadratic   bool
	valorMaximo float64
}

func main() {
	m1 := Matrix{
		x:         3,
		y:         3,
		cuadratic: false,
	}

	m1.Set(2.0, 3.0, 4.0, 1.0, 2.0, 2.0, 5.0, 6.0, 7.0)
	m1.Print()
}

func (m *Matrix) Set(valores ...float64) {
	temp := 0.0
	x := 0
	y := 0
	var matrix [3][3]float64
	for _, valor := range valores {

		if temp <= valor {
			temp = valor
		}

		matrix[x][y] = valor

		y += 1
		if y == 3 || y == 6 {
			x += 1
			y = 0
		}
	}
	m.Values = matrix
	m.valorMaximo = temp
}

func (m Matrix) Print() {
	for i := 0; uint16(i) < m.x; i++ {
		for j := 0; uint16(j) < m.y; j++ {
			fmt.Printf("%v ", m.Values[i][j])
		}
		fmt.Printf("\n")
	}

	fmt.Printf("Valor en X: %v\n", m.x)
	fmt.Printf("Valor en Y: %v\n", m.y)
	fmt.Printf("Es cuadratica: %v\n", m.cuadratic)
	fmt.Printf("Valor maximo: %v\n", m.valorMaximo)
}
