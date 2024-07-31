package main

import "fmt"

func main() {
	data := []int{5, 12, 2, 15, 4, 3, 11, 1, 6, 10, 8, 14, 7, 9, 13}

	bubble := make([]int, len(data))
	selection := make([]int, len(data))
	copy(bubble, data)
	copy(selection, data)

	fmt.Printf("Before: %v\n", data)

	BubbleSort(bubble)
	fmt.Printf("Sorted (Bubble): %v\n", bubble)

	SelectionSort(selection)
	fmt.Printf("Sorted (Selection): %v\n", selection)
}

func BubbleSort(array []int) {
	var swap, compare int

	for range len(array) {
		for j := range len(array) - 1 {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swap++
			}
			compare++
		}
	}
	fmt.Printf("Swaps: %d, Compares: %d\n", swap, compare)
}

func SelectionSort(array []int) {
	var index, swap, compare int

	for i := range len(array) {
		index = i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[index] {
				index = j
			}
			compare++
		}
		array[i], array[index] = array[index], array[i]
		swap++
	}
	fmt.Printf("Swaps: %d, Compares: %d\n", swap, compare)
}
