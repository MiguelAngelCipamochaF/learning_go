package main

import "fmt"

func main() {
	var nombre string
	var estatura_persona int
	var apellido string
	var edad int
	nombre = "Miguel"
	apellido = "Hernandez"
	edad = 19
	var licenciaDeConducir = true
	apellido2 := "Perez"
	estatura_persona = 174
	cantidadDeHijos := 2

	fmt.Printf("Variable apellido: %v declarada correctamente.\n", apellido)
	fmt.Printf("Variable cantidadDeHijos: %v declarada correctamente\n", cantidadDeHijos)
	fmt.Printf("Variable licencia_de_conducir: %v declarada incorrectamente\n", licenciaDeConducir)
	fmt.Printf("Variable nombre: %v declarada incorrectamente.\n", nombre)
	fmt.Printf("Variable edad: %v declarada incorrectamente.\n", edad)
	fmt.Printf("Variable apellido2: %v declarada incorrectamente.\n", apellido2)
	fmt.Printf("Variable estatura_persona: %v declarada incorrectamente.\n", estatura_persona)

}
