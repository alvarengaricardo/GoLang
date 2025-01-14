package main

import "fmt"

func main() {
	var num1, num2 int

	// Solicita ao usuário para inserir os números
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	// Calcula a soma
	soma := num1 + num2

	// Exibe o resultado
	fmt.Printf("A soma de %d e %d é %d\n", num1, num2, soma)
}
