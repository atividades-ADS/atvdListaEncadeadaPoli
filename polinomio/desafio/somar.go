package main

import (
	"atvdpoli/polinomio"
)

func somar(p1, p2 polinomio.Polinomio) polinomio.Polinomio {
	// polinomio que será resultante das somas
	result := polinomio.InitPolinomio()

	p1 = *simplificar(p1)
	p2 = *simplificar(p2)

	current := p1.Termos.Head

	for current != nil {

		termoP2 := p2.SearchExpoente(current.Value.Expoente)

		if termoP2 != nil {
			termoResult := polinomio.Termo{
				Coeficiente: current.Value.Coeficiente + termoP2.Coeficiente,
				Expoente:    current.Value.Expoente,
			}
			result.InsertTermo(termoResult)
		} else {
			result.InsertTermo(current.Value)
		}
		current = current.Next

	}

	current = p2.Termos.Head

	for current != nil {
		if result.SearchExpoente(current.Value.Expoente) == nil {
			result.InsertTermo(current.Value)

		}
		current = current.Next
	}

	return *result
}
