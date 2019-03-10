package main

import (
	"errors"
	"fmt"
)

const maxSortSize = 100000

func SortDescending(data []int) ([][2]int, error) {
	size := len(data)
	if size < 1 && size > maxSortSize {
		return nil, errors.New("Invalid size")
	}
	//convert array to heap
	swaps := buildHeap(data)

	for i := len(data) - 1; i > 0; i-- {
		swap(data, 0, i, nil)
		size--
		siftDown(data, 0, size, nil)
	}

	return swaps, nil
}

func buildHeap(data []int) [][2]int {

	// convert array into heap
	swaps := [][2]int{}

	halfSize := (len(data) - 1) / 2
	for i := halfSize / 2; i >= 0; i-- {
		siftDown(data, i, len(data), &swaps)
	}
	fmt.Printf("Heap: %v\n", data)
	return swaps
}

func siftDown(data []int, index, size int, swaps *[][2]int) {
	maxIndex := index

	left := leftChildIndex(data, index)
	if left < size && data[left] < data[maxIndex] {
		maxIndex = left
	}

	right := rightChildIndex(data, index)
	if right < size && data[right] < data[maxIndex] {
		maxIndex = right
	}

	if maxIndex != index {
		swap(data, index, maxIndex, swaps)
		siftDown(data, maxIndex, size, swaps)
	}
}

func leftChildIndex(data []int, index int) int {
	return (2 * index) + 1
}

func rightChildIndex(data []int, index int) int {
	return (2 * index) + 2
}

func swap(data []int, i, j int, swaps *[][2]int) {
	if swaps != nil {
		*swaps = append(*swaps, [2]int{i, j})
	}
	tmp := data[i]
	data[i] = data[j]
	data[j] = tmp
}
