package main

import (
	"fmt"
)

func main() {
	var tamanho int

	// 1. Ler o tamanho do array.
	fmt.Println("Informe o tamanho do array:")
	fmt.Scan(&tamanho)

	if tamanho <= 0 {
		fmt.Println("Por favor, insira um tamanho positivo.")
		return
	}

	// 2. Ler os valores do array.
	array := make([]int, tamanho)
	fmt.Println("Informe os valores do array:")
	for i := 0; i < tamanho; i++ {
		fmt.Scan(&array[i])
	}

	// 3. Verificar se o array está ordenado em ordem crescente.
	ordenado := true
	for i := 1; i < tamanho; i++ {
		if array[i-1] > array[i] {
			ordenado = false
			break
		}
	}

	// 4. Imprimir o resultado da verificação.
	if ordenado {
		fmt.Println("O array está ordenado em ordem crescente.")
	} else {
		fmt.Println("O array não está ordenado em ordem crescente.")
	}
}
