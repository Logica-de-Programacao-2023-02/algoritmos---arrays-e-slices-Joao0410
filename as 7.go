package main

import (
	"fmt"
)

func main() {
	// 1. Crie uma fatia de inteiros com tamanho inicial de 5.
	slice := make([]int, 0, 5)

	// 2. Solicite ao usuário que informe um valor.
	var valor int
	fmt.Println("Informe um número inteiro:")
	fmt.Scan(&valor)

	// 3. Verifique se o número já está presente na fatia.
	presente := false
	for _, v := range slice {
		if v == valor {
			presente = true
			break
		}
	}

	// 4. Se o número não estiver presente, adicione-o à fatia.
	if !presente {
		slice = append(slice, valor)
	} else {
		fmt.Println("Número já está presente no slice.")
	}

	// 5. Imprima a fatia resultante.
	fmt.Println(slice)
}
