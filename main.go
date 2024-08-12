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

	//-----Coin change
	coins := []int{2, 4, 5}
	amount := 12
	result := exercises.CoinChange(coins, amount)
	fmt.Printf("Minimum number of coins needed to make %d: %d\n", amount, result)

	//-----------------Longest Palindrome sub string
	s := "racecar"
	res2 := exercises.LongestPalindrome(s)
	fmt.Printf("Input: %s\nLongest palindrome: %s\n\n", s, res2)

}
