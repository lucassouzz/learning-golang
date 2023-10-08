package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiros

	var p *int = nil

	p = &i // pega o endereço da variável
	*p++   // desrefrenciando (pegando valor)
	i++

	fmt.Println(p, *p, i, &i)
}
