package main

import (
	"fmt"
)

func main() {
	// Defina um array bidimensional com 3 linhas e 2 colunas
	var matriz [3][2]int

	// Solicite ao usuário os valores de cada elemento da matriz
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Informe o valor para o elemento na linha %d, coluna %d: ", i+1, j+1)
			fmt.Scan(&matriz[i][j])
		}
	}

	// Imprima a matriz resultante
	fmt.Println("Matriz resultante:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Println() // para começar uma nova linha após cada linha da matriz
	}
}
