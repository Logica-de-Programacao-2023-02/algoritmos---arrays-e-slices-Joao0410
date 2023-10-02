package main

import (
	"fmt"
)

// isPrime verifica se um número é primo.
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// generatePrimes gera os n primeiros números primos.
func generatePrimes(n int) []int {
	primes := make([]int, 0, n)
	for i := 2; len(primes) < n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	var n int

	// Ler um número inteiro positivo n.
	fmt.Println("Informe um número inteiro positivo:")
	fmt.Scan(&n)

	// Verificar se o número é positivo.
	if n <= 0 {
		fmt.Println("Por favor, insira um número positivo.")
		return
	}

	// Gerar e imprimir os n primeiros números primos.
	primes := generatePrimes(n)
	fmt.Printf("Os %d primeiros números primos são: %v\n", n, primes)
}
