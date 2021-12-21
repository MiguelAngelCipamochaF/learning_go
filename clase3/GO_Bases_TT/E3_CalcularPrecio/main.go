package main

import "fmt"

type Productos struct {
	Nombre   string
	Precio   int
	Cantidad int
}

type Servicios struct {
	Nombre  string
	Precio  int
	Minutos int
}

type Mantenimiento struct {
	Nombre string
	Precio int
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	productosArray := []Productos{}
	serviciosArray := []Servicios{}
	mantenimientoArray := []Mantenimiento{}

	p1 := Productos{
		Nombre:   "Manzana",
		Precio:   5000,
		Cantidad: 5,
	}

	productosArray = append(productosArray, p1)

	s1 := Servicios{
		Nombre:  "Trabajo X",
		Precio:  670,
		Minutos: 2345,
	}

	serviciosArray = append(serviciosArray, s1)

	m1 := Mantenimiento{
		Nombre: "Carro",
		Precio: 234000,
	}

	mantenimientoArray = append(mantenimientoArray, m1)

	go SumarProductos(productosArray, c1)
	go SumarServicios(serviciosArray, c2)
	go SumarMantenimiento(mantenimientoArray, c3)

	totalP := <-c1
	totalS := <-c2
	totalM := <-c3

	fmt.Printf("Total Productos: %v\n", totalP)
	fmt.Printf("Total Servicios: %v\n", totalS)
	fmt.Printf("Total Mantenimiento: %v\n", totalM)
	fmt.Printf("Total: %v\n", totalP+totalS+totalM)
}

func SumarProductos(productos []Productos, c chan int) {
	total := 0
	for _, p := range productos {
		precio := p.Precio
		cantidad := p.Cantidad

		total += precio * cantidad
	}

	c <- total
}

func SumarServicios(servicios []Servicios, c chan int) {
	total := 0

	for _, s := range servicios {
		mins := s.Minutos
		if s.Minutos < 30 {
			mins = 30
		}
		precio := s.Precio

		total += precio * (mins / 30)
	}

	c <- total
}

func SumarMantenimiento(mantenimiento []Mantenimiento, c chan int) {
	total := 0
	for _, m := range mantenimiento {
		total += m.Precio
	}

	c <- total
}
