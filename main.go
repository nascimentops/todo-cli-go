package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type tarefa struct {
	tarefa    string
	conclusao bool
}

var reader = bufio.NewScanner(os.Stdin)
var toDo []tarefa

func main() {
	var opcao int

	fmt.Println("Bem vindo a lista de afazeres, o que você deseja fazer?")
	fmt.Println("1. Adicionar")
	fmt.Println("2. Marcar/desmarcar como concluído")
	fmt.Println("3. Remover")
	fmt.Println("4. Parar o programa")

	reader.Scan()
	opcao, _ = strconv.Atoi(reader.Text())
	for opcao != 4 {

		switch opcao {
		case 1:
			adicionar()
		case 2:
			concluir()
		case 3:
			remover()
		default:
			fmt.Println("opção inválida")
		}

		if len(toDo) >= 1 {
			fmt.Println()
			for i, v := range toDo {
				fmt.Println(i+1, "tarefa:", v.tarefa, "\tstatus de conclusão:", v.conclusao)
			}
			fmt.Println()
		}

		fmt.Println("1. Adicionar")
		fmt.Println("2. Marcar/desmarcar como concluído")
		fmt.Println("3. Remover")
		fmt.Println("4. Parar o programa")
		reader.Scan()
		opcao, _ = strconv.Atoi(reader.Text())
	}
}

func adicionar() {
	nome := ""
	fmt.Println("digite a tarefa:")
	reader.Scan()
	nome = reader.Text()

	if nome != "" {
		toDo = append(toDo, tarefa{nome, false})
	}
}

func remover() {
	var tarefa int
	fmt.Println("Qual tarefa deseja remover? ")
	reader.Scan()
	tarefa, _ = strconv.Atoi(reader.Text())
	if tarefa > len(toDo) || tarefa < 1 {
		fmt.Println("Opção inválida")
	} else {
		toDo = append(toDo[:tarefa-1], toDo[tarefa:]...)
	}
}

func concluir() {
	var tarefa int
	fmt.Println("Qual tarefa deseja mudar o status?")
	reader.Scan()
	tarefa, _ = strconv.Atoi(reader.Text())
	if tarefa > len(toDo) || tarefa < 1 {
		fmt.Println("Opção inválida")
	} else {
		if toDo[tarefa-1].conclusao == true {
			toDo[tarefa-1].conclusao = false
		} else {
			toDo[tarefa-1].conclusao = true
		}
	}
}
