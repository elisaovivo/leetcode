package main

import (
	"fmt"
	"math/rand"
)

func loopCompare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func main() {
	testTime := 5000
	for i := 0; i < testTime; i++ {
		s1 := generateRandomArray(1000, 200)
		var s2 []int
		for _, v := range s1 {
			s2 = append(s2, v)
		}
		selectionSort(s1)
		bubbleSort(s2)
		if !loopCompare(s1, s2) {
			fmt.Println(s1, s2)
		}
	}
}

func swap(a, b int) {
	a, b = b, a
}
func selectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for i := range arr {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func bubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}
func insertSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}
func generateRandomArray(maxsize, maxvalue int) []int {
	var s []int
	for i := 0; i < maxsize; i++ {
		s = append(s, rand.Int()%(maxvalue+1))
	}
	return s
}
