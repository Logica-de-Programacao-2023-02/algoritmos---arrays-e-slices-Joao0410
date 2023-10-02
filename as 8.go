package main

import (
	"fmt"
)

func main() {
	// 1. Crie um slice de strings com 8 valores.
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	// 2. Solicite ao usuário que informe um valor.
	var valor string
	fmt.Println("Informe uma string para remover:")
	fmt.Scan(&valor)

	// 3. Itere sobre o slice e remova todas as ocorrências dessa string.
	var novoSlice []string
	for _, v := range slice {
		if v != valor {
			novoSlice = append(novoSlice, v)
		}
	}
	slice = novoSlice

	// 4. Imprima o slice resultante.
	fmt.Println(slice)
}
