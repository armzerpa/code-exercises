package main

import (
	"fmt"

	"github.com/armzerpa/code-exercises/exercises"
)

func main() {
	fmt.Println("Hello world")
	nums := []int{3, 2, 1, 5, 6, 4, 8}
	k := 2
	kthlargest := exercises.FindKthLargest(nums, k)
	fmt.Println("Quick Selection: ", kthlargest)
	kthlargest = exercises.FindKthLargestUsingHeap(nums, k)
	fmt.Println("Heap: ", kthlargest)
}
