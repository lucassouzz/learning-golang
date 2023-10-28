package main

import "math"

// Inicializando método ou var com letra maiúscula ele será PUBLICO (visível dentro e fora do pacote)
// Inicializando método ou var com letra minúscula ele será PRIVADO (visível apenas dentro e fora do pacote)

// Por exemplo...
// Ponto - gerará algo público
// ponto ou _Ponto - gerará algo privado

type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
