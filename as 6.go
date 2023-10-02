package main

import (
	"fmt"
)

func main() {
	// 1. Crie um array de inteiros com 10 elementos.
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 2. Solicite ao usuário que informe um valor.
	var valor int
	fmt.Println("Informe um valor:")
	fmt.Scan(&valor)

	// 3. Verifique se o valor está presente no array.
	encontrado := false
	for _, v := range array {
		if v == valor {
			encontrado = true
			break
		}
	}

	// 4. Imprima o resultado da busca.
	if encontrado {
		fmt.Printf("O valor %d está presente no array.\n", valor)
	} else {
		fmt.Printf("O valor %d não foi encontrado no array.\n", valor)
	}
}
