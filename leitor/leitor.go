package leitor

import (
	"bufio"
	"os"
	"regexp"
)

func ArquivoParaMatriz(nomeArquivo string) ([][]string, error) {
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()

	re := regexp.MustCompile(`[-+]?\d+(?:\.\d+)?|[^\s]`)

	var matriz [][]string

	scanner := bufio.NewScanner(arquivo)

	for scanner.Scan() {
		linha := scanner.Text()

		elementos := re.FindAllString(linha, -1)

		matriz = append(matriz, elementos)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matriz, nil
}
