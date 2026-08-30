package main

import (
	"atvdpoli/polinomio"
)

func grau(p polinomio.Polinomio) int {

	current := p.Termos.Head
	poliGrau := 0
	for current != nil {
		if poliGrau > current.Value.Expoente {
			poliGrau = current.Value.Expoente
			continue
		}

		current = current.Next
	}

	return poliGrau
}
