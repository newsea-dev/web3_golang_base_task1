package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		current := intervals[i]
		if current[0] <= last[1] {
			if last[1] > current[1] {
				current[1] = last[1]
			}
			result[len(result)-1] = last
		} else {
			result = append(result, current)
		}
	}
	return result
}

func main() {
	// 测试案例
	test1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println("测试案例1合并结果:", merge(test1)) // 预期 [[1 6] [8 10] [15 18]]

	test2 := [][]int{{1, 4}, {4, 5}}
	fmt.Println("测试案例2合并结果:", merge(test2)) // 预期 [[1 5]]

	test3 := [][]int{{1, 4}, {2, 3}}
	fmt.Println("测试案例3合并结果:", merge(test3)) // 预期 [[1 4]]
}
