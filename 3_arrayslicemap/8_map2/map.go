package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José Joao":      123123.32,
		"Gabriela Silva": 4345334.54,
		"Pedro Paulo":    9898.87,
	}

	funcsESalarios["Antonio Nunes"] = 123.45
	delete(funcsESalarios, "Inxistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
