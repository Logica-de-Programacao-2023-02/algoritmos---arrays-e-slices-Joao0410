package main

import (
	"fmt"
)

func main() {
	// 1. Crie um array de inteiros com 10 elementos.
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 2. Itere sobre o array nas posições pares e 3. Calcule a soma dos elementos nessas posições.
	soma := 0
	for i := 0; i < len(array); i += 2 {
		soma += array[i]
	}

	// 4. Imprima a soma resultante.
	fmt.Println("A soma dos elementos nas posições pares é:", soma)
}
