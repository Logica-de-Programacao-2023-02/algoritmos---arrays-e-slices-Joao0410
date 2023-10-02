package main

import (
	"fmt"
)

func main() {
	// 1. Crie um array de floats com 10 elementos.
	array := [10]float64{1.5, 2.0, 5.5, 4.2, 7.8, 8.1, 3.0, 6.2, 9.5, 4.6}

	// 2. Itere sobre o array e adicione ao slice os elementos que são maiores que 5.
	var slice []float64
	for _, val := range array {
		if val > 5 {
			slice = append(slice, val)
		}
	}

	// 3. Imprima o slice resultante.
	fmt.Println(slice)
}
