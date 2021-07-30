package main

import "fmt"

func moveZeroes(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}
	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
	fmt.Println("nums:", nums)
	fmt.Println("k:", k)
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}
