package main

import (
	"fmt"
)

func main() {
	// 1. Crie um array de inteiros com 10 elementos.
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 2. Itere sobre o array e adicione ao slice os elementos que são pares.
	var slice []int
	for _, val := range array {
		if val%2 == 0 {
			slice = append(slice, val)
		}
	}

	// 3. Imprima o slice resultante.
	fmt.Println(slice)
}
