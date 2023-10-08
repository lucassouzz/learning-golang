package main

import (
	"fmt"
	m "math" // alias para um import
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo float64 inferido pelo compilador

	area := PI * m.Pow(raio, 2) // forma reduzida de criar uma variável

	fmt.Println("A área da circunferência é", area)

	const ( // blocos de construção de constantes
		x = 1
		y = 2
	)

	var ( // blocos de construção de variáveis
		w = 3
		z = 4
	)

	fmt.Println("O valor de w é:", w, ". O valor de z é:", z)

	var e, f bool = true, false // declarar variáveis em sequencia

	fmt.Println(e, f)

	g, h, i := 2, false, "oba!" // declarando conjunto de variáveis por inferencia

	fmt.Println(g, h, i)
}
