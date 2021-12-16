package main

import "fmt"

func main() {
	var apellido string = "Gomez"
	var edad int = 35
	boolean := false
	var sueldo float64 = 45857.90
	var nombre string = "Julian"

	fmt.Printf("Variable apellido: %v tipo string correcto\n", apellido)
	fmt.Printf("Variable edad: %v tipo int correcto\n", edad)
	fmt.Printf("Variable boolean: %v tipo string incorrecto\n", boolean)
	fmt.Printf("Variable sueldo: %v tipo string incorrecto\n", sueldo)
	fmt.Printf("Variable nombre: %v tipo string correcton\n", nombre)
	fmt.Println("Correciones hechas!")
}
