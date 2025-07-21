package main

import (
	"fmt"
)

func insertionSort(A []int) {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1

		// Sposta gli elementi di A[0..j-1] che sono > key
		// di una posizione in avanti
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
}

func main() {
	A := []int{5, 2, 4, 6, 1, 3}
	fmt.Println("Prima dell'ordinamento:", A)
	insertionSort(A)
	fmt.Println("Dopo l'ordinamento:    ", A)
}
