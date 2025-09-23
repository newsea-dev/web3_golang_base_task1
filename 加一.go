package main

import "fmt"

//	func plusOne(digits []int) []int {
//		for i := len(digits) - 1; i >= 0; i-- {
//			if digits[i] != 9 {
//				digits[i]++
//				for j := i + 1; j < len(digits); j++ {
//					digits[j] = 0
//				}
//				return digits
//			}
//		}
//		digits = make([]int, len(digits)+1)
//		digits[0] = 1
//		return digits
//	}
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func main() {
	var num int
	fmt.Print("请输入一个整数: ")
	fmt.Scan(&num)
	digits := make([]int, 0)
	for num > 0 {
		digits = append([]int{num % 10}, digits...) // 直接在前面插入，避免后续反转
		num /= 10
	}
	fmt.Print("加1后的结果是: ", plusOne(digits))
}
