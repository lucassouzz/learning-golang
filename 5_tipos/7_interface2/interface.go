package main

import "fmt"

type esportivo interface {
	ligarTrubo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTrubo() {
	f.turboLigado = true
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTrubo()

	var carro2 = &ferrari{"F41", false, 0}
	carro2.ligarTrubo()

	fmt.Println(carro1, carro2)
}
