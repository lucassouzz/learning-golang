package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[11234567894] = "Maria"
	aprovados[98273927282] = "José"
	aprovados[52636726377] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF:%d)\n", nome, cpf)
	}

	fmt.Println(aprovados[52636726377])
	delete(aprovados, 52636726377)
	fmt.Println(aprovados[52636726377])
}
