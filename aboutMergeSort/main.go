package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	testTime := 100000
	maxsize := 10
	maxvalue := 200
	for i := 0; i < testTime; i++ {
		s1 := generateRandomArray(maxsize, maxvalue)
		var s2 []int
		for _, v := range s1 {
			s2 = append(s2, v)
		}

		if smallSum(s1) != comparator(s2) {
			fmt.Println(smallSum(s1), comparator(s2))
			return
		}
	}

}
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
func generateRandomArray(maxsize, maxvalue int) []int {
	rand.Seed(time.Now().UnixNano())
	var s []int
	for i := 0; i < maxsize; i++ {
		s = append(s, rand.Int()%(maxvalue+1))
	}
	return s
}

// 小和问题
func smallSum(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return mergeSort(arr, 0, len(arr)-1)
}
func mergeSort(arr []int, l, r int) int {
	if l >= r {
		return 0
	}
	mid := l + (r-l)>>1
	lSum := mergeSort(arr, l, mid)
	rSum := mergeSort(arr, mid+1, r)
	ans := lSum + rSum
	tmp := make([]int, 0)
	ptrL, ptrR := l, mid+1
	for ptrL <= mid && ptrR <= r {
		if arr[ptrL] < arr[ptrR] {
			ans += r - ptrR + 1
			tmp = append(tmp, arr[ptrL])
			ptrL++
		} else {
			tmp = append(tmp, arr[ptrR])
			ptrR++
		}
	}
	for ptrL <= mid {
		tmp = append(tmp, arr[ptrL])
		ptrL++
	}
	for ptrR <= r {
		tmp = append(tmp, arr[ptrR])
		ptrR++
	}
	for k, v := range tmp {
		arr[l+k] = v
	}
	return ans
}

func comparator(arr []int) int {
	ans := 0
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				ans++
			}
		}
	}
	return ans
}
