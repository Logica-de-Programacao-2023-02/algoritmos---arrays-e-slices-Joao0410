package main

import (
	"fmt"
)

func main() {
	// Crie um array com 4 elementos float64
	arr := [4]float64{1.2, 3.4, 5.6, 7.8}

	// Inicie o produto com 1, já que 1 é o elemento neutro da multiplicação
	produto := 1.0

	// Multiplique cada elemento do array pelo produto
	for _, valor := range arr {
		produto *= valor
	}

	// Imprima o produto resultante
	fmt.Println("O produto dos valores é:", produto)
}
