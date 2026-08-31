package main

import (
	"atvdpoli/polinomio"
)

func subtrair(p1, p2 polinomio.Polinomio) polinomio.Polinomio {
	// polinomio que será resultante das subtrações
	result := polinomio.InitPolinomio()

	p1 = *simplificar(p1)
	p2 = *simplificar(p2)

	current := p1.Termos.Head

	for current != nil {

		termoP2 := p2.SearchExpoente(current.Value.Expoente)

		termoResult := current.Value

		if termoP2 != nil {
			termoResult.Coeficiente -= termoP2.Coeficiente
		}

		// Insere somente se o coeficiente não for zero
		if termoResult.Coeficiente != 0 {
			result.InsertTermo(termoResult)
		}

		current = current.Next

	}

	current = p2.Termos.Head

	for current != nil {

		// Verifica em p1, e não em result
		if p1.SearchExpoente(current.Value.Expoente) == nil {

			termoResult := current.Value
			termoResult.Coeficiente = -termoResult.Coeficiente

			if termoResult.Coeficiente != 0 {
				result.InsertTermo(termoResult)
			}
		}

		current = current.Next
	}

	return *result
}
