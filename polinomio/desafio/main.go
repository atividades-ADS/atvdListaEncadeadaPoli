package main

import (
	"atvdpoli/leitor"
	"atvdpoli/polinomio"
	"fmt"

	//Biblioteca para conversão de string para int/float64
	"strconv"
)

func IsAction(l string) bool {
	switch string(l) {
	case "G", "g", "T", "t", "+", "-", "*", "A", "a", "P", "p", "S", "s":
		return true
	}
	return false
}

func doAction(action string, polinomios []polinomio.Polinomio) {
	switch action {

	case "G", "g":
		if len(polinomios) < 1 {
			return
		}

		g := grau(polinomios[0])
		fmt.Println("Grau:", g)

	case "T", "t":
		if len(polinomios) < 1 {
			return
		}

		t := tamanho(polinomios[0])
		fmt.Println("Tamanho:", t)

	case "+":
		if len(polinomios) < 2 {
			return
		}

		som := somar(polinomios[0], polinomios[1])
		fmt.Println("Soma:", Exibicao(som))

	case "-":
		if len(polinomios) < 2 {
			return
		}

		sub := subtrair(polinomios[0], polinomios[1])
		fmt.Println("Subtração:", Exibicao(sub))

	case "*":
		if len(polinomios) < 2 {
			return
		}

		m := multiplicar(polinomios[0], polinomios[1])
		fmt.Println("Multiplicação:", Exibicao(m))

	case "A", "a":
		if len(polinomios) < 2 {
			return
		}

		if polinomios[0].Termos.Head == nil {
			a := avaliar(polinomios[1], 0.0)
			fmt.Println("Avaliação:", a)
			return
		}

		x := polinomios[0].Termos.Head.Value.Coeficiente

		fmt.Println("x =", x)

		a := avaliar(polinomios[1], x)

		fmt.Println("Avaliação:", a)

	case "P", "p":
		if len(polinomios) < 1 {
			return
		}

		fmt.Println("Polinômio:", Exibicao(polinomios[0]))

	case "S", "s":
		if len(polinomios) < 1 {
			return
		}

		sim := simplificar(polinomios[0])
		fmt.Println("Simplificado:", Exibicao(*sim))
	}
}

func criarPolinomio(row []string) polinomio.Polinomio {
	p := polinomio.InitPolinomio()
	// Se a linha não é uma ação, é um polinômio, então processa os elementos da linha, mesmo que só tenha um elemento
	for i := 0; i < len(row); i += 2 {

		// A base recebe o numero atual, usa o strconv para converter de string para float64
		base, err := strconv.ParseFloat(row[i], 64)
		if err != nil {
			fmt.Println("Erro ao converter para float64:", err)
			break
		}
		//O expoente recebe o proximo numero, se existir, senao, expoente = 0. usa o strconv para converter de string para int
		var expoente int
		if i+1 < len(row) {
			expoente, err = strconv.Atoi(row[i+1])
			if err != nil {
				fmt.Println("Erro ao converter para int:", err)
				break
			}
		} else {
			expoente = 0
		}
		termo := polinomio.Termo{
			Coeficiente: base,
			Expoente:    expoente,
		}

		// Cria um novo polinomio e insere o termo nele, depois adiciona o polinomio na lista de polinomios
		p.InsertTermo(termo)
	}
	return *p

}

func main() {

	// Le o arquivo e retorna uma matriz de strings.
	rows, err := leitor.ArquivoParaMatriz("./polinomio.txt")

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	numrow := len(rows)

	var action string
	polinomios := make([]polinomio.Polinomio, 0)

	for rowIndex := 0; rowIndex < numrow; rowIndex++ {

		//Fiz uma lista de polinomios, para cada linha do arquivo, cria um novo polinomio e adiciona na lista

		row := rows[rowIndex]

		// Ignora linhas vazias
		if len(row) == 0 {
			continue
		}

		// Verifica se a linha contém apenas um caracter, se sim, verifica se é uma ação
		if len(row) == 1 && IsAction(row[0]) {

			// Se já havia uma ação, significa que os parâmetros da anterior acabaram, então roda ela.
			if action != "" {
				doAction(action, polinomios)
				polinomios = nil
			}

			action = row[0]
			continue
		}

		// Se a linha não é uma ação, é um polinômio, então processa os elementos da linha, mesmo que só tenha um elemento
		p := criarPolinomio(row)
		polinomios = append(polinomios, p)

	}
	if action != "" {
		doAction(action, polinomios)
	}
}
