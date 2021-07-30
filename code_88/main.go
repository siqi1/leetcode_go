package main

//给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。
//func merge(nums1 []int, m int, nums2 []int, n int)  {
//	//index := 0
//	for k := range nums1{
//		if m == 0 {
//			nums1[k] = nums2[k]
//		}
//		if k > m {
//			nums1[k] = nums2[k-n]
//		}
//	}
//	sort.Ints(nums1)
//	fmt.Println("nums1:",nums1)
//}

func merge(nums1 []int, m int, nums2 []int, n int) {
	k, j := m-1, n-1

	for i := len(nums1) - 1; k >= 0 || j >= 0 && i >= 0; i-- {
		n := 0

		if k >= 0 && j >= 0 {
			if nums1[k] > nums2[j] {
				n = nums1[k]
				k--
			} else {
				n = nums2[j]
				j--
			}
		} else if k < 0 {
			n = nums2[j]
			j--
		} else {
			n = nums1[k]
			k--
		}

		nums1[i] = n
	}
}

func main() {
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1
	merge(nums1, m, nums2, n)
}
