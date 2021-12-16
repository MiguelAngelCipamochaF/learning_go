package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	mayoresA21 := 0

	fmt.Println("Lista de empleados")
	for employee, edad := range employees {
		fmt.Println(employee, edad)
	}

	fmt.Printf("Edad de Benjamin es: %v\n", employees["Benjamin"])

	employees["Federico"] = 25
	fmt.Println("Se agrego a Federico a la lista de empleados")
	delete(employees, "Pedro")
	fmt.Println("Se elimino a Pedro de la lista de empleados")

	for _, value := range employees {
		if value > 21 {
			mayoresA21++
		}
	}

	fmt.Printf("La cantidad de empleados mayores a 21 años son: %v\n", mayoresA21)

}
