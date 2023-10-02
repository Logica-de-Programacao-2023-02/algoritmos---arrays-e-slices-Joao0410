package main

import (
	"fmt"
)

func main() {
	var tamanho int

	// 1. Ler o tamanho dos arrays.
	fmt.Println("Informe o tamanho dos arrays:")
	fmt.Scan(&tamanho)

	if tamanho <= 0 {
		fmt.Println("Por favor, insira um tamanho positivo.")
		return
	}

	// 2. Ler os valores dos dois arrays.
	array1 := make([]int, tamanho)
	array2 := make([]int, tamanho)

	fmt.Println("Informe os valores do primeiro array:")
	for i := 0; i < tamanho; i++ {
		fmt.Scan(&array1[i])
	}

	fmt.Println("Informe os valores do segundo array:")
	for i := 0; i < tamanho; i++ {
		fmt.Scan(&array2[i])
	}

	// 3. Gera um terceiro array com a soma.
	arraySoma := make([]int, tamanho)
	for i := 0; i < tamanho; i++ {
		arraySoma[i] = array1[i] + array2[i]
	}

	// 4. Imprimir o terceiro array.
	fmt.Println("A soma dos arrays é:", arraySoma)
}
