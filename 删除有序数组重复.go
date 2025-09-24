package main

import (
	"fmt"
)

func removeD(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	nums1 := []int{1, 1, 2}
	len1 := removeD(nums1)
	fmt.Printf("新长度: %d, 数组前%d个元素: %v\n", len1, len1, nums1[:len1])

	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	len2 := removeD(nums2)
	fmt.Printf("新长度: %d, 数组前%d个元素: %v\n", len2, len2, nums2[:len2])
}
