package calc

import (
	"fmt"

	"github.com/vitormoschetta/go/calc/math"
	"github.com/vitormoschetta/go/utils"
)

func Execute() {
	fmt.Println("Bem vindo ao Go!")
	fmt.Println("Escolha uma das opções abaixo:")
	fmt.Println("1 - Soma")
	fmt.Println("2 - Subtração")
	fmt.Println("3 - Multiplicação")
	fmt.Println("4 - Divisão")
	fmt.Println("5 - Sair")
	var opcao int
	fmt.Scanln(&opcao)
	if opcao == 5 {
		return
	}
	var a, b int
	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&a)
	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&b)
	switch opcao {
	case 1:
		fmt.Println("Resultado da soma:", math.Add(a, b))
	case 2:
		fmt.Println("Resultado da subtração:", math.Subtract(a, b))
	case 3:
		fmt.Println("Resultado da multiplicação:", math.Multiply(a, b))
	case 4:
		fmt.Println("Resultado da divisão:", math.Divide(a, b))
	default:
		fmt.Println("Opção inválida!")
	}

	fmt.Println("Digite uma data no formato dd/mm/aaaa:")
	var date string
	fmt.Scanln(&date)
	fmt.Println("Data convertida:", utils.StringToDate(date))
}
