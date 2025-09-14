package main

import "fmt"

func singleNumber(nums []int) int {
	// 创建map存储数字出现次数
	countMap := make(map[int]int)

	// 遍历数组，统计每个数字出现的次数
	for _, num := range nums {
		countMap[num]++
	}

	// 3. 遍历map，找到次数为1的数字并返回
	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}
	return 0
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println("只出现一次的数字是", singleNumber(nums))
}
