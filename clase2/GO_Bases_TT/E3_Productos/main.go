package main

import "fmt"

type tienda struct {
	Productos []producto
}

type producto struct {
	Tipo   string
	Nombre string
	Precio float64
}

type Producto interface {
	CalcularCosto()
}

type Ecommerce interface {
	Total() float64
	Agregar(prod producto)
}

func main() {

	p1 := nuevoProducto("mediano", "Taladro", 24000.0)
	p1.CalcularCosto()

	p2 := nuevoProducto("grande", "Nevera", 82000.0)
	p2.CalcularCosto()

	p3 := nuevoProducto("pequeño", "Fifa", 2000.0)
	p3.CalcularCosto()

	tien := nuevaTienda()
	tien.Agregar(p1)
	tien.Agregar(p2)
	tien.Agregar(p3)
	precioTotal := tien.Total()

	fmt.Printf("El precio total de los productos es: %v\n", precioTotal)
}

func nuevoProducto(tipoProducto string, nombre string, precio float64) producto {
	nuevoProducto := producto{
		Tipo:   tipoProducto,
		Nombre: nombre,
		Precio: precio,
	}

	return nuevoProducto
}

func nuevaTienda() Ecommerce {
	t1 := &tienda{}

	return t1
}

func (p *producto) CalcularCosto() {
	nuevoPrecio := 0.0
	switch {
	case p.Tipo == "pequeño":
		nuevoPrecio = p.Precio
		p.Precio = nuevoPrecio
	case p.Tipo == "mediano":
		nuevoPrecio = p.Precio + 3*p.Precio/100
		p.Precio = nuevoPrecio
	case p.Tipo == "grande":
		nuevoPrecio = p.Precio + 6*p.Precio/100
		nuevoPrecio += 2500
		p.Precio = nuevoPrecio
	default:
		fmt.Println("No existe ese tipo de producto")
	}
}

func (t *tienda) Total() float64 {
	precioTotal := 0.0
	for _, p := range t.Productos {
		precioTotal += p.Precio
	}

	return precioTotal
}

func (t *tienda) Agregar(p producto) {
	t.Productos = append(t.Productos, p)
}
