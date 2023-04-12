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
	testTime := 100000
	maxsize := 1000
	maxvalue := 200
	for i := 0; i < testTime; i++ {
		s1 := generateRandomArray(maxsize, maxvalue)
		var s2 []int
		for _, v := range s1 {
			s2 = append(s2, v)
		}
		selectionSort(s1)
		mergeSortNoRecursive(s2, 0, maxsize-1)
		if !loopCompare(s1, s2) {
			fmt.Println(s1, s2)
		}
	}
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

func merge(arr []int, l, mid, r int) {
	tmp := make([]int, r-l+1)
	p1 := l
	p2 := mid + 1
	i := 0
	for p1 <= mid && p2 <= r {
		if arr[p1] <= arr[p2] {
			tmp[i] = arr[p1]
			p1++
		} else {
			tmp[i] = arr[p2]
			p2++
		}
		i++
	}
	for ; p1 <= mid; p1++ {
		tmp[i] = arr[p1]
		i++
	}
	for ; p2 <= r; p2++ {
		tmp[i] = arr[p2]
		i++
	}
	//fmt.Println(tmp)
	for i = 0; i < r-l+1; i++ {
		arr[l+i] = tmp[i]
	}
	//fmt.Println(arr)
}
func mergeSort(arr []int, l, r int) {
	if l == r {
		return
	}
	mid := l + ((r - l) >> 1)
	//fmt.Println(mid)
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}
func mergeSortNoRecursive(arr []int, l, r int) {
	N := len(arr)
	mergeSize := 1
	for mergeSize < N {
		l := 0
		for l < N {
			m := l + mergeSize - 1
			if m >= N {
				break
			}
			var r int
			if m+mergeSize > N-1 {
				r = N - 1
			} else {
				r = m + mergeSize
			}
			merge(arr, l, m, r)
			l = r + 1
		}
		if mergeSize > N/2 {
			break
		}
		mergeSize *= 2
	}
}
func quickSort(arr []int, L int, R int) {
	if L < R {
		random := L + rand.Int()%(R-L+1)
		arr[R], arr[random] = arr[random], arr[R]
		tmp := partition(arr, L, R)
		quickSort(arr, L, tmp[0]-1)
		quickSort(arr, tmp[1]+1, R)
	}
}
func partition(arr []int, L int, R int) []int {
	less, more := L-1, R
	for L < more {
		if arr[L] < arr[R] {
			less++
			arr[less], arr[L] = arr[L], arr[less]
			L++
		} else if arr[L] > arr[R] {
			more--
			arr[more], arr[L] = arr[L], arr[more]
		} else {
			L++
		}
	}
	arr[L], arr[R] = arr[R], arr[L]

	return []int{less + 1, more - 1}
}

func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}
func heapify(arr []int, index, heapSize int) {
	left := index * 2
	for left < heapSize {
		var largest int
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		} else {
			largest = left
		}
		if arr[largest] > arr[index] {
			arr[largest], arr[index] = arr[index], arr[largest]
		}
		index = largest
		left = index*2 + 1
	}
}
