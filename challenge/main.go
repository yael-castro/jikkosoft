package main

import (
	"fmt"
	"math"
)

// Dado un arreglo de números positivos enteros como por Ej: [10, 15, 3, 7, 2]
// y un resultado k Ej: 17. Halle la suma de 2 números que den k y
// retorna todas las posiciones del arreglo

func doSomething(k int, num ...int) []int {
	memory := make(map[int]int)
	results := make([]int, 0, len(num))

	for i, n := range num {
		_, exists := memory[n]
		if exists {
			results = append(results, i) // O(1)
			continue
		}

		remaining := int(math.Abs(float64(k - n)))
		memory[remaining] = i
	}

	return results
}

func main() {
	results := doSomething(17, 10, 15, 3, 7, 2)
	fmt.Println(results)

}
