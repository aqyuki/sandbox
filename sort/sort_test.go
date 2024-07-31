package main

import "testing"

func TestBubbleSort(t *testing.T) {
	t.Parallel()

	data := []int{5, 12, 2, 15, 4, 3, 11, 1, 6, 10, 8, 14, 7, 9, 13}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	BubbleSort(data)
	for i := range data {
		if data[i] != expected[i] {
			t.Errorf("Expected %d, but got %d", expected[i], data[i])
		}
	}
}

func TestSelectionSort(t *testing.T) {
	t.Parallel()

	data := []int{5, 12, 2, 15, 4, 3, 11, 1, 6, 10, 8, 14, 7, 9, 13}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	SelectionSort(data)
	for i := range data {
		if data[i] != expected[i] {
			t.Errorf("Expected %d, but got %d", expected[i], data[i])
		}
	}
}
