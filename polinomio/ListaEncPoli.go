package polinomio

import (
	"atvdpoli/listaEnc"
)

type Termo struct {
	Coeficiente float64
	Expoente    int
}

type Polinomio struct {
	Termos *listaEnc.ListaEnc[Termo]
}

func InitPolinomio() *Polinomio {
	return &Polinomio{
		Termos: listaEnc.InitListaEnc[Termo](),
	}
}

func (p *Polinomio) InsertTermo(t Termo) {
	if t.Coeficiente == 0 {
		return
	}
	p.Termos.InsertTail(t)
}

func (p Polinomio) SearchExpoente(expoente int) *Termo {
	current := p.Termos.Head

	for current != nil {
		if current.Value.Expoente == expoente {
			return &current.Value
		}

		current = current.Next
	}

	return nil
}
