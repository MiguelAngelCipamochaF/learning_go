package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	inicio := time.Now().UnixMilli()

	go InsertionSort(variable1, c1)
	fmt.Printf("Tiempo variable1 Insertion sort: %f\n", <-c1)
	go BubbleSort(variable1, c1)
	fmt.Printf("Tiempo variable1 Bubble sort: %f\n", <-c1)
	go SelectionSort(variable1, c1)
	fmt.Printf("Tiempo variable1 Selection sort: %f\n\n", <-c1)

	go InsertionSort(variable2, c2)
	fmt.Printf("Tiempo variable2 Insertion sort: %f\n", <-c2)
	go BubbleSort(variable2, c2)
	fmt.Printf("Tiempo variable2 Bubble sort: %f\n", <-c2)
	go SelectionSort(variable2, c2)
	fmt.Printf("Tiempo variable2 Selection sort: %f\n\n", <-c2)

	go InsertionSort(variable3, c3)
	fmt.Printf("Tiempo variable3 Insertion sort: %f\n", <-c3)
	go BubbleSort(variable3, c3)
	fmt.Printf("Tiempo variable3 Bubble sort: %f\n", <-c3)
	go SelectionSort(variable3, c3)
	fmt.Printf("Tiempo variable3 Selecion sort: %f\n\n", <-c3)

	fin := time.Now().UnixMilli()
	fmt.Printf("\nToda la ejecucion duró %d milisegundos\n", fin-inicio)
}

func InsertionSort(array []int, c chan float64) {
	inicio := time.Now().UnixMilli()
	for i := 0; i < len(array); i++ {
		actual := array[i]
		for j := i; j > 0 && array[j-1] > actual; j-- {
			array[j] = array[j-1]
		}
		array[i] = actual
	}
	fin := time.Now().UnixMilli()
	fmt.Printf("Ordenamiento por insercion completo para array de: %v\n", len(array))

	c <- float64(fin - inicio)
}

func BubbleSort(array []int, c chan float64) {
	iteracion := 0
	permutation := true

	inicio := time.Now().UnixMilli()
	for permutation {
		permutation = false
		iteracion++
		for actual := 0; actual < len(array)-iteracion; actual++ {
			if array[actual] > array[actual+1] {
				permutation = true
				temp := array[actual]
				array[actual] = array[actual+1]
				array[actual+1] = temp
			}
		}
	}
	fin := time.Now().UnixMilli()
	fmt.Printf("Ordenamiento por burbuja completo para array de: %v\n", len(array))

	c <- float64(fin - inicio)
}

func SelectionSort(array []int, c chan float64) {

	inicio := time.Now().UnixMilli()
	for actual := 0; actual < len(array); actual++ {
		minor := actual
		for j := actual + 1; j < len(array); j++ {
			if array[j] < array[minor] {
				minor = j
			}
		}
		if minor != actual {
			temp := array[actual]
			array[actual] = array[minor]
			array[minor] = temp
		}
	}
	fin := time.Now().UnixMilli()

	fmt.Printf("Ordenamiento por seleccion completo para array de: %v\n", len(array))

	c <- float64(fin - inicio)
}
