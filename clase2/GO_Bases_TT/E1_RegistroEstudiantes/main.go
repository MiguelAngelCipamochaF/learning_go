package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	Dni      string
	Fecha    string
}

func main() {

	e1 := Alumnos{
		Nombre:   "Miguel",
		Apellido: "Cipamocha",
		Dni:      "1001",
		Fecha:    "13-12-2021",
	}

	e2 := Alumnos{
		Nombre:   "Anonimo",
		Apellido: "Perez",
		Dni:      "9999",
		Fecha:    "10-04-2010",
	}

	e1.Detalle()
	e2.Detalle()

}

func (a Alumnos) Detalle() {
	fmt.Printf("Nombre: %v\n", a.Nombre)
	fmt.Printf("Apellido: %v\n", a.Apellido)
	fmt.Printf("Dni: %v\n", a.Dni)
	fmt.Printf("Fecha: %v\n", a.Fecha)
}
