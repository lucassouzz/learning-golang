package main

import (
	"fmt"
	"reflect"
)

func main() {

	// slice == fatia
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  // slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// slice não é um array! slice define um pedaço de um array.
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2]
	fmt.Println(a2, s3)

	// slice possui um tamanho e um ponteiro que indica o primeiro elemento
	// apartir do primeiro elemento é utilizado o tamanho para definir o conjunto de elementos
	s4 := s2[:1]
	fmt.Println(s2, s4)
}
