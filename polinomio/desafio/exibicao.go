package main

import (
	"atvdpoli/polinomio"
	"fmt"
	"strconv"
)

func Exibicao(p polinomio.Polinomio) string {
	if p.Termos == nil || p.Termos.Head == nil {
		return "0"
	}

	var resultado string
	current := p.Termos.Head

	for current != nil {
		termo := current.Value

		coeficiente := termo.Coeficiente
		expoente := termo.Expoente

		// Sinal
		if resultado != "" {
			if coeficiente >= 0 {
				resultado += " + "
			} else {
				resultado += " - "
				coeficiente = -coeficiente
			}
		} else if coeficiente < 0 {
			resultado += "-"
			coeficiente = -coeficiente
		}

		// Termo
		switch {
		case expoente == 0:
			resultado += fmt.Sprintf("%s", strconv.FormatFloat(coeficiente, 'f', -1, 64))

		case coeficiente == 1:
			resultado += "x"

		default:
			resultado += fmt.Sprintf("%sx", strconv.FormatFloat(coeficiente, 'f', -1, 64))
		}

		// Expoente
		if expoente > 1 {
			resultado += fmt.Sprintf("^%s", strconv.Itoa(expoente))
		}

		current = current.Next
	}

	return resultado
}
