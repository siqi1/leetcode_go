package main

import "fmt"

//func removeDuplicates(nums []int) int {
//	k := 0
//	for i:=0 ; i< len(nums) ; i++ {
//		if i != len(nums) -1 && nums[i] != nums[i+1] {
//			nums[k] = nums[i]
//			k ++
//		}else if i == len(nums) -1 {
//			nums[k] = nums[i]
//			k ++
//		}
//	}
//	fmt.Println("nums:",nums)
//	return k
//}

func removeDuplicates(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	fmt.Println("nums:", nums)
	return k
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := removeDuplicates(nums)
	fmt.Println("result:", result)
}
