package main

import (
	"fmt"
)

func main() {
	// 1. Crie um array de inteiros com 5 elementos.
	array := [5]int{1, 3, 5, 6, 9}

	// 2. Itere sobre o array e adicione ao slice os elementos que são múltiplos de 3.
	var slice []int
	for _, val := range array {
		if val%3 == 0 {
			slice = append(slice, val)
		}
	}

	// 3. Imprima o slice resultante.
	fmt.Println(slice)
}
