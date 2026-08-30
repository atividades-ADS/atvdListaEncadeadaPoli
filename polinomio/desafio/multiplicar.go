package main

import (
	"atvdpoli/polinomio"
)

func multiplicar(p1, p2 polinomio.Polinomio) polinomio.Polinomio {
	resultado := polinomio.InitPolinomio()

	p1 = *simplificar(p1)
	p2 = *simplificar(p2)

	poli1 := p1.Termos.Head

	for poli1 != nil {

		poli2 := p2.Termos.Head

		for poli2 != nil {

			coeficiente := poli1.Value.Coeficiente * poli2.Value.Coeficiente
			expoente := poli1.Value.Expoente + poli2.Value.Expoente

			termoResult := polinomio.Termo{
				Coeficiente: coeficiente,
				Expoente:    expoente,
			}

			resultado.InsertTermo(termoResult)

			poli2 = poli2.Next
		}

		poli1 = poli1.Next
	}
	resultado = simplificar(*resultado)

	return *resultado
}
