package main

import (
	"fmt"
)

func main() {
	// 1. Crie um array de floats com 6 elementos.
	array := [6]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}

	// 2. Solicite ao usuário que informe um valor.
	var valor float64
	fmt.Print("Informe um número float para adicionar: ")
	if _, err := fmt.Scan(&valor); err != nil {
		fmt.Println("Erro ao ler o valor:", err)
		return
	}

	// 3. Adicione esse número a todas as posições do array.
	for i, num := range array {
		array[i] = num + valor
	}

	// 4. Imprima o array resultante.
	fmt.Println(array)
}
