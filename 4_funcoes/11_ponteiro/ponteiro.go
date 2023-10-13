package main

import "fmt"

func inc1(n int) {
	n++
}

// um ponteiro é representado por um *
func inc2(n *int) {
	// * é utilizado quando para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta

	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	// & é utilizado para obter o endereço de determinada variável
	inc2(&n) // por referência
	fmt.Println(n)
}
