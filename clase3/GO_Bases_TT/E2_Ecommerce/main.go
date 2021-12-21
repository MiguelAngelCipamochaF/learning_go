package main

import "fmt"

type Usuario struct {
	Nombre   string
	Apellido string
	Correo   string
	Producto []producto
}

type producto struct {
	Nombre   string
	Precio   int
	Cantidad int
}

func main() {
	u1 := &Usuario{Nombre: "Miguel"}
	p1 := NuevoProducto("Manzana", 4500)
	p2 := NuevoProducto("Naranja", 6000)

	AgregarProducto(u1, p1, 5)
	AgregarProducto(u1, p2, 10)
	fmt.Printf("Producto usuario %v: %v\n", u1.Nombre, u1.Producto)
	BorrarProducto(u1)
	fmt.Printf("Producto usuario %v: %v\n", u1.Nombre, u1.Producto)
}

func NuevoProducto(nombre string, precio int) *producto {
	p1 := producto{Nombre: nombre, Precio: precio}

	return &p1
}

func AgregarProducto(usuario *Usuario, producto *producto, cantidad int) {
	producto.Cantidad = cantidad

	usuario.Producto = append(usuario.Producto, *producto)
}

func BorrarProducto(usuario *Usuario) {
	usuario.Producto = nil
}
