package main

import (
	"atvdpoli/polinomio"
	"math"
)

func avaliar(p polinomio.Polinomio, x float64) float64 {
	resultado := 0.0

	current := p.Termos.Head

	for current != nil {
		coeficiente := current.Value.Coeficiente
		expoente := current.Value.Expoente

		if x == 0 && expoente == 0 {
			resultado += coeficiente
			current = current.Next
			continue
		}

		if x == 0 {
			current = current.Next
			continue
		}

		resultado += coeficiente * math.Pow(x, float64(expoente))

		current = current.Next
	}
	return resultado
}
