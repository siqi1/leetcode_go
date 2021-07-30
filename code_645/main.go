package main

import "fmt"

//集合 s 包含从 1 到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合 丢失了一个数字 并且 有一个数字重复 。
//
//给定一个数组 nums 代表了集合 S 发生错误后的结果。
//
//请你找出重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。
//func findErrorNums(nums []int) []int {
//	value := make(map[int]int)
//	repet := 0
//	defi := 0
//	for i := 0; i< len(nums); i++  {
//		if v,has := value[nums[i]] ; has {
//			repet = v
//		}
//		value[nums[i]] = nums[i]
//		if i+1 != nums[i] {
//			defi = i+1
//		}
//	}
//	fmt.Println("repet:",repet,"defi:",defi)
//	return []int{repet,defi}
//
//}

func findErrorNums(nums []int) []int {
	_extra := make([]int, len(nums)+1)
	repet := 0
	defi := 0
	for k := range nums {
		_extra[nums[k]]++
	}

	for i := 1; i < len(_extra); i++ {
		extra := _extra[i]
		if extra > 1 {
			repet = i
		}
		if extra == 0 {
			defi = i
		}
	}
	fmt.Println("repet:", repet, "defi:", defi)
	return []int{repet, defi}

}

func main() {
	nums := []int{3, 2, 2}
	findErrorNums(nums)
}
