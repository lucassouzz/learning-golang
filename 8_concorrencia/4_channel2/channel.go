package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines
// channel é um tipo

func doisTresQuartoVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviado dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)

	go doisTresQuartoVezes(1, c)

	a, b := <-c, <-c
	fmt.Println(a, b)

	fmt.Println(<-c)
}
