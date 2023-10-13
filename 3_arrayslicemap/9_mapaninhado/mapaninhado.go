package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel Pereira": 123.45,
			"Gustavo Silva":   6765.56,
		},
		"J": {
			"Joao Pedro": 6454.54,
		},
		"K": {
			"Karlos": 1213.123,
		},
	}

	fmt.Println(funcsPorLetra)
	delete(funcsPorLetra, "K")
	fmt.Println(funcsPorLetra)

	for letra, registros := range funcsPorLetra {
		fmt.Println(letra, registros)
	}
}
