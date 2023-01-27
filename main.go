package main

import (
	"fmt"
	"os"
)

func adicao(num1, num2 float64) float64 {
	return num1 + num2
}

func subtracao(num1, num2 float64) float64 {
	return num1 - num2
}

func multiplicacao(num1, num2 float64) float64 {
	return num1 * num2
}

func divisao(num1, num2 float64) float64 {
	if num2 == 0 {
		return 0
	} else {
		return num1 / num2
	}
}

func exibeMenu() {

	nome := "Lucas"

	fmt.Println("Olá,", nome)
	fmt.Println("1 - Somar")
	fmt.Println("2 - Subtrair")
	fmt.Println("3 - Multiplicar")
	fmt.Println("4 - Dividir")
	fmt.Println("5 - Sair do programa")
	fmt.Println("Digite a opção da operação desejada:")

}

func leOpcao() int {

	var opcaoLida int

	fmt.Scan(&opcaoLida)
	return opcaoLida

}

func main() {

	var num1, num2, resultado float64

	for {

		exibeMenu()

		opcao := leOpcao()

		if opcao == 5 {
			fmt.Println("Saindo do programa")
			os.Exit(0)
		}

		fmt.Println("Informe o primeiro número:")
		fmt.Scanln(&num1)

		fmt.Println("Informe o segundo número:")
		fmt.Scanln(&num2)

		switch opcao {
		case 1:
			resultado = adicao(num1, num2)
		case 2:
			resultado = subtracao(num1, num2)
		case 3:
			resultado = multiplicacao(num1, num2)
		case 4:
			if num2 == 0 {
				fmt.Println("Divisão por zero!")
			}
			resultado = divisao(num1, num2)
		default:
			fmt.Println("Opção inválida. Digite novamente!")
			continue
		}

		fmt.Println("O resultado da operação é", resultado)
	}
}
