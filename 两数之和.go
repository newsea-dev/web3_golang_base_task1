package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, exist := numsMap[complement]; exist {
			return []int{j, i}
		}
		numsMap[num] = i
	}
	return []int{}
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Println("测试案例1结果:", twoSum(nums1, target1))

	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Println("测试案例1结果:", twoSum(nums2, target2))

	nums3 := []int{3, 3}
	target3 := 6
	fmt.Println("测试案例1结果:", twoSum(nums3, target3))
}
