package main

import "fmt"

func main() {
	precio := 20000
	descuento := 40

	descuentoAPrecio := precio * descuento / 100
	total := precio - descuentoAPrecio

	fmt.Printf("El precio sin descuento del producto es: %v\n", precio)
	fmt.Printf("El descuento es: %v\n", descuentoAPrecio)
	fmt.Printf("El valor del producto con descuento es: %v\n", total)
}
