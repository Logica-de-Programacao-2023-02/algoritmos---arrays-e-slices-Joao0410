package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 1. Inicialize o gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	// Crie uma fatia de inteiros com 10 valores aleatórios.
	slice := make([]int, 10)
	for i := range slice {
		slice[i] = rand.Intn(100) // Valores aleatórios entre 0 e 99
	}

	// Imprima a fatia gerada
	fmt.Println("Slice:", slice)

	// Verifique se a fatia está vazia antes de continuar
	if len(slice) == 0 {
		fmt.Println("O slice está vazio!")
		return
	}

	// Inicialize min e max com o primeiro valor da fatia
	min, max := slice[0], slice[0]

	// Percorra a fatia para encontrar os valores mínimo e máximo.
	for _, val := range slice[1:] { // Comece do segundo elemento, já que o primeiro já foi usado para inicializar
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	// Imprima os valores mínimo e máximo.
	fmt.Printf("Mínimo: %d\n", min)
	fmt.Printf("Máximo: %d\n", max)
}
