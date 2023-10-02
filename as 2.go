package main

import (
	"fmt"
)

func main() {
	// Crie uma fatia com os valores de 1 a 5
	slice := []int{1, 2, 3, 4, 5}

	// Remova o terceiro elemento
	// Isso pode ser feito cortando a fatia e depois anexando os pedaços restantes juntos
	slice = append(slice[:2], slice[3:]...)

	// Imprima a fatia resultante
	fmt.Println(slice)
}
