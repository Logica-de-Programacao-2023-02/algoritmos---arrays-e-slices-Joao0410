package main

import (
	"fmt"
)

func main() {
	// 1. Crie uma fatia de inteiros com 8 valores.
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// 2. Solicite ao usuário que informe dois índices.
	var idx1, idx2 int
	fmt.Println("Informe o primeiro índice (0-7):")
	fmt.Scan(&idx1)

	fmt.Println("Informe o segundo índice (0-7):")
	fmt.Scan(&idx2)

	// Verifique se os índices são válidos.
	if idx1 < 0 || idx1 >= len(slice) || idx2 < 0 || idx2 >= len(slice) {
		fmt.Println("Índices inválidos!")
		return
	}

	// 3. Troque os valores nos índices fornecidos.
	slice[idx1], slice[idx2] = slice[idx2], slice[idx1]

	// 4. Imprima o slice resultante.
	fmt.Println(slice)
}
