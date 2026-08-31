package main

import (
	"atvdpoli/polinomio"
)

func simplificar(p polinomio.Polinomio) *polinomio.Polinomio {
	result := polinomio.InitPolinomio()

	current := p.Termos.Head

	processados := make(map[int]bool)

	for current != nil {

		expoente := current.Value.Expoente

		if processados[expoente] {
			current = current.Next
			continue
		}

		// Marca o expoente como processado,
		// mesmo que o resultado da soma seja zero.
		processados[expoente] = true

		coeficiente := current.Value.Coeficiente

		simplificando := current.Next

		for simplificando != nil {
			if simplificando.Value.Expoente == expoente {
				coeficiente += simplificando.Value.Coeficiente
			}
			simplificando = simplificando.Next
		}

		// Só adiciona se a soma for diferente de zero
		if coeficiente != 0 {
			termoResult := polinomio.Termo{
				Coeficiente: coeficiente,
				Expoente:    expoente,
			}

			result.InsertTermo(termoResult)
		}

		current = current.Next
	}

	return result
}
