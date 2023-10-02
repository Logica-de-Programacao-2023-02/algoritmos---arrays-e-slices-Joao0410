package main

import (
	"fmt"
)

func main() {
	// 1. Crie uma matriz bidimensional de inteiros com 2 linhas e 3 colunas.
	matriz := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	// 2. Solicite ao usuário que informe os índices da linha e da coluna.
	var linha, coluna int
	fmt.Println("Informe o índice da linha (0 ou 1):")
	fmt.Scan(&linha)

	fmt.Println("Informe o índice da coluna (0, 1 ou 2):")
	fmt.Scan(&coluna)

	// 3. Valide se os índices informados são válidos para a matriz.
	if linha < 0 || linha >= 2 || coluna < 0 || coluna >= 3 {
		fmt.Println("Índices informados são inválidos!")
		return
	}

	// 4. Imprima o valor encontrado nos índices informados.
	fmt.Printf("Valor na posição [%d][%d]: %d\n", linha, coluna, matriz[linha][coluna])
}
