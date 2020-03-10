package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := adder()
	for i := 0; i < len(array); i++ {
		fmt.Println(sum(array[i]))
	}

}
