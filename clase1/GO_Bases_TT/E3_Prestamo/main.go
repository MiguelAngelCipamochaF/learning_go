package main

import "fmt"

func main() {
	edad := 23
	empleado := true
	antiguedad := 1
	sueldo := 90000

	if edad > 22 && empleado && antiguedad >= 1 {
		fmt.Println("Prestamo otorgado")
		if sueldo >= 100000 {
			fmt.Println("No te cobraremos intereses")
		}
	} else {
		fmt.Println("Prestamo denegado")
	}
}
