package slicer

import (
	"fmt"
)

func Slice(a int) (int, int, int, float64) {
	var (
		min_num int
		max_num int
		sum     int
		avg     float64
		num     int
		nums    []int
	)

	for i := 0; i < a; i++ {
		fmt.Print("Input a number: ")
		fmt.Scan(&num)
		nums = append(nums, num)
	}

	sum = 0
	for _, v := range nums {
		sum += v
	}

	avg = float64(sum) / float64(a)

	min_num = nums[0]
	max_num = nums[0]
	for _, val := range nums {
		if val > max_num {
			max_num = val
		}
		if val < min_num {
			min_num = val
		}
	}

	return sum, min_num, max_num, avg
}
