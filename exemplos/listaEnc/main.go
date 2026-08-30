package main

import (
	"atvdpoli/listaEnc"
	"fmt"
)

func main() {

	acoes := []string{"InsertHead", "InsertTail", "Insert", "RemoveHead", "RemoveTail", "Remove", "Search", "GetValue", "GetNext", "Update", "Destroy", "ShowAll", "Size", "IsEmpty", "Exist", "Exit"}
	lista := listaEnc.InitListaEnc[int]()
	var opcao int

	fmt.Println("Interface de manipulação da lista encadeada simples.")
	fmt.Println("A lista já foi inicializada. comece a manipular os dados.")
	fmt.Println("==============================================================")
	fmt.Println("Você pode inserir, remover, buscar, atualizar e destruir elementos da lista.")
	for true {

		for i, acao := range acoes {
			p := fmt.Sprint(i, ":", acao)
			fmt.Println(p)
		}
		fmt.Println("Digite o número da ação que deseja realizar:")
		fmt.Scanf("%d", &opcao)
		switch opcao {
		case 0:
			var valor int
			fmt.Println("Digite o valor a ser inserido no início da lista:")
			fmt.Scanf("%d", &valor)
			lista.InsertHead(valor)
		case 1:
			var valor int
			fmt.Println("Digite o valor a ser inserido no final da lista:")
			fmt.Scanf("%d", &valor)
			lista.InsertTail(valor)
		case 2:
			var valor int
			var pos int
			fmt.Println("Digite o valor a ser inserido na lista:")
			fmt.Scanf("%d", &valor)
			fmt.Println("Digite a posição onde deseja inserir o valor:")
			fmt.Scanf("%d", &pos)
			lista.Insert(pos, valor)
		case 3:
			lista.RemoveHead()
		case 4:
			lista.RemoveTail()
		case 5:
			var pos int
			fmt.Println("Digite a posição do elemento que deseja remover:")
			fmt.Scanf("%d", &pos)
			lista.Remove(pos)
		case 6:
			var valor int
			fmt.Println("Digite o valor que deseja buscar na lista:")
			fmt.Scanf("%d", &valor)
			no := lista.Search(valor)

			if no != nil {
				fmt.Println("O valor", valor, "foi encontrado no nó: ", no)
			} else {
				fmt.Println("O valor", valor, "não foi encontrado na lista.")
			}
		case 7:
			var id int
			fmt.Println("Digite o id do elemento que deseja obter o valor:")
			fmt.Scanf("%d", &id)
			valor := lista.GetValue(id)
			if valor != nil {
				fmt.Println("O valor na posição", id, "é:", *valor)
			} else {
				fmt.Println("Não há elemento com id", id)
			}
		case 8:
			var id int
			fmt.Println("Digite o id do elemento que deseja obter o próximo valor:")
			fmt.Scanf("%d", id)
			valor := lista.GetNext(id)
			if valor != nil {
				fmt.Println("O próximo valor após a posição", id, "é:", *valor)
			} else {
				fmt.Println("Não há próximo elemento após a posição", id)
			}
		case 9:
			var id int
			var novoValor int
			fmt.Println("Digite o id do elemento que deseja atualizar:")
			fmt.Scanf("%d", &id)
			fmt.Println("Digite o novo valor:")
			fmt.Scanf("%d", &novoValor)
			lista.Update(id, novoValor)
		case 10:
			lista.Destroy(0)
		case 11:
			lista.ShowAll()
		case 12:
			size := lista.Size()
			fmt.Println("O tamanho da lista é:", size)
		case 13:
			isEmpty := lista.IsEmpty()
			if isEmpty {
				fmt.Println("A lista está vazia.")
			} else {
				fmt.Println("A lista não está vazia.")
			}
		case 14:
			var valor int
			fmt.Println("Digite o valor que deseja verificar se existe na lista:")
			fmt.Scanf("%d", &valor)
			exists := lista.Exist(valor)
			if exists {
				fmt.Println("O valor", valor, "existe na lista.")
			} else {
				fmt.Println("O valor", valor, "não existe na lista.")
			}
		case 15:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}

	}

}
