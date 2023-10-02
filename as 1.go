package main

import (
	"fmt"
)

func main() {
	// Criando um array de inteiros com 3 elementos
	numeros := [3]int{10, 20, 30}

	// Calculando a soma dos valores no array
	soma := 0
	for _, valor := range numeros {
		soma += valor
	}

	// Imprimindo a soma
	fmt.Printf("A soma dos valores no array é: %d\n", soma)
}
