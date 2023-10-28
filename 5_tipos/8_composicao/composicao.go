package main

import "fmt"

type esportivo interface {
	ligarTrubo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	//é possível adicionar mais métodos
}

type bmw7 struct{}

func (bmw bmw7) ligarTrubo() {
	fmt.Println("Turbo...")
}

func (bmw bmw7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var bmw esportivoLuxuoso = bmw7{}
	bmw.ligarTrubo()
	bmw.fazerBaliza()
}
