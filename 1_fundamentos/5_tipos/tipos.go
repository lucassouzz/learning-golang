package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só valores positivos)... uint8 uint16 uint32 uint64 (u -> unsignal)

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // rune armazena o valor atribuito correspondente na tabela Unicode
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println("O valor da variável i2 é", i2)

	// números reais float32, float64
	var x float32 = 49.99 // ou var x = float32(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49 é", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Minha string"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	s2 := `Minha
	string
	!`
	fmt.Println("O tamanho da string é", len(s2))

	// char???
	char := 'a'
	fmt.Println("O tipo char é", reflect.TypeOf(char))
}
