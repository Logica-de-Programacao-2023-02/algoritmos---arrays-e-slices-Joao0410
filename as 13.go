package main

import (
	"fmt"
)

func main() {
	// 1. Crie um array de inteiros com 7 elementos.
	array := [7]int{1, 2, 3, 4, 5, 6, 7}

	// 2. Solicite ao usuário que informe um valor.
	var valor int
	fmt.Println("Informe um número inteiro para adicionar:")
	fmt.Scan(&valor)

	// 3. Adicione esse número ao primeiro e ao último elemento do array.
	array[0] += valor
	array[len(array)-1] += valor

	// 4. Imprima o array resultante.
	fmt.Println(array)
}
