package main

import "fmt"

func main() {

}

func printOddTimesNum(arr []int) {
	eor := 0 //a^b
	for _, v := range arr {
		eor ^= v
	}
	rightone := eor & (^eor + 1)
	onlyone := 0 //a or b
	for _, v := range arr {
		if (rightone & v) == 0 {
			onlyone ^= v
		}
	}
	fmt.Println(onlyone, onlyone^eor)
}
