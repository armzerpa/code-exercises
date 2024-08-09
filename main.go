package main

import (
	"fmt"

	"github.com/armzerpa/code-exercises/exercises"
)

func main() {
	fmt.Println("Hello world")
	nums := []int{8, 2, 1, 5, 6, 4}
	k := 3
	kthlargest := exercises.FindKthLargest(nums, k)
	fmt.Println(kthlargest)
}
