package main

import (
	"fmt"
)

func main() {
	var tamanho int

	// Pergunte ao usuário o tamanho da fatia
	fmt.Print("Informe o tamanho da fatia: ")
	fmt.Scan(&tamanho)

	// Crie uma fatia do tamanho especificado
	slice := make([]int, tamanho)

	// Solicite ao usuário os valores dos elementos
	for i := 0; i < tamanho; i++ {
		fmt.Printf("Informe o valor para o elemento %d: ", i+1)
		fmt.Scan(&slice[i])
	}

	// Calcule a soma dos valores
	soma := 0
	for _, valor := range slice {
		soma += valor
	}

	// Imprima o slice e a soma dos valores
	fmt.Println("Slice:", slice)
	fmt.Println("Soma dos valores:", soma)
}
