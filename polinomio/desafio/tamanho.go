package main

import (
	"atvdpoli/polinomio"
)

func tamanho(p polinomio.Polinomio) int {
	counter := 0
	current := p.Termos.Head

	for current != nil {
		counter++
		current = current.Next
	}

	return counter
}
