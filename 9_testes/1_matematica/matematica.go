package matematica

import (
	"fmt"
	"strconv"
)

// Média é responsavel por calcular a média dos valores
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, numero := range numeros {
		total += numero
	}

	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%2.f", media), 64)

	return mediaArredondada
}
