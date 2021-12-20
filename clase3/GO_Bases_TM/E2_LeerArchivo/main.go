package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var total int
	words, err := scanLines("productos.txt")

	if err != nil {
		fmt.Printf("Error reading product\n")
	}

	fmt.Println("ID\tPrecio\tCantidad")

	for _, word := range words {
		word = strings.Replace(word, ";", "\t", -1)
		precio := string(word[3:9])
		cantidad := string(word[9:])
		precioInt := 0
		cantidadInt := 0

		_, err := fmt.Sscan(precio, &precioInt)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Sscan(cantidad, &cantidadInt)

		total += precioInt * cantidadInt

		fmt.Print(word + "\n")
	}

	fmt.Printf("\t%v\n", total)

}

func scanLines(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
}
