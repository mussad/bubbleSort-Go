package main

import "fmt"

func main() {
	var numbers = []int{8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("Our List of unsorted Numbers is:", numbers)
	bubbleSort(numbers)
	fmt.Println("The list After bubbleSorting ", numbers)
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
