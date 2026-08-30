package main

import (
	"atvdpoli/polinomio"
)

func simplificar(p polinomio.Polinomio) *polinomio.Polinomio {
	result := polinomio.InitPolinomio()

	current := p.Termos.Head

	for current != nil {

		coeficiente := current.Value.Coeficiente
		expoente := current.Value.Expoente

		// se o expoente já existe no polinomio resultante, pula para o próximo termo
		if result.SearchExpoente(current.Value.Expoente) != nil || coeficiente == 0 {
			current = current.Next
			continue
		}

		simplificando := current.Next

		for simplificando != nil {
			if simplificando.Value.Expoente == expoente {
				coeficiente += simplificando.Value.Coeficiente
			}
			simplificando = simplificando.Next
		}

		termoResult := polinomio.Termo{
			Coeficiente: coeficiente,
			Expoente:    expoente,
		}

		result.InsertTermo(termoResult)

		current = current.Next
	}

	return result
}
