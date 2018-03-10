/*

Basic bubble sort algorithm with out any optimization in Go, Golang.
link : https://en.wikipedia.org/wiki/Bubble_sort
Muhammad F. Musad 2018 | m.musad@outlook.com

*/

package main

import (
	"fmt"
	"log"
	"math/big"
	"time"
)

func main() {
	// ------------------------------ //
	// Note : this lines did not belong to Bubble Sort Algorithm, it's just for calculating algorithm running time
	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))
	// ------------------------------ //

	// Bubble Sort Start From here
	// put here any list of numbers that you need to sort it
	var unsortedNumbers = []int{7, 4, 3, 2, 6, 3, 7, 8, 53453764, 23213123, 312312, 54546, 112, 43242, 234324, 5435345, 23234, 2324234, 324, 66, 9999, 967869, 98979, 97897665, 56756976786, 7567567, 3, 1, 2, 3, 7, 8, 9, 5, 3, 8, 2, 8, 9, 11, 0, 8, 45}
	fmt.Println("Our List of unsorted Numbers is:", unsortedNumbers)

	// put list in the algorithm function
	bubbleSort(unsortedNumbers)

	// defined slice to handle output NOTE : its Optional becuse you can view it diracty from function.
	var sorted []int

	// put data in the slice
	sorted = unsortedNumbers

	// print new sorted data
	fmt.Println("The list After bubbleSorting ", sorted)
	// bubble sort end here

	// ------------------------------ //
	// Note This not belonge to Algorithm itself, its just to calculte running time
	elapsed := time.Since(start)
	log.Printf("running took %s", elapsed)
	// ------------------------------ //

}

func bubbleSort(numbers []int) {
	var N = len(numbers)
	var i int
	for i = 0; i < N; i++ {
		bubbleSwap(numbers)
	}
}

func bubbleSwap(numbers []int) {
	var N = len(numbers)
	var firstIndex = 0
	var secondIndex = 1
	for secondIndex < N {
		var firstNumber = numbers[firstIndex]
		var secondNumber = numbers[secondIndex]

		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}
		firstIndex++
		secondIndex++
	}
}
