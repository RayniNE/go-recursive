package main

import (
	"fmt"
	"math/rand"
)

func main() {
	unsorted := getRandomArray()
	fmt.Println("Unsorted Array")
	fmt.Println(unsorted)

	sorted := mergeSort(unsorted)

	fmt.Println("Sorted Array")
	fmt.Println(sorted)
}

// getRandomArray generates an array of random numbers of length 100
func getRandomArray() []int {
	var arr []int

	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(200))
	}

	return arr
}

// mergeSort divides the provided array into two arrays and calls itself recursively
// twice with the first and second half of the original array
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	half := len(arr) / 2

	firstHalf := mergeSort(arr[:half])
	secondHalf := mergeSort(arr[half:])

	return merge(firstHalf, secondHalf)
}

// merge takes the two half of the original array, sorts them, and then append it into a final array
func merge(firstHalf, secondHalf []int) []int {
	final := []int{}
	i := 0
	j := 0

	firstLen := len(firstHalf)
	secondLen := len(secondHalf)

	for i < firstLen && j < secondLen {
		if firstHalf[i] < secondHalf[j] {
			final = append(final, firstHalf[i])
			i++
			continue
		}

		final = append(final, secondHalf[j])
		j++
	}

	for ; i < firstLen; i++ {
		final = append(final, firstHalf[i])
	}
	for ; j < secondLen; j++ {
		final = append(final, secondHalf[j])
	}
	return final
}
