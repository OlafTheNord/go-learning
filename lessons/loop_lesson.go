package lessons

import "fmt"

func ProcessStream(count int) (sum, min, max int, avg float64) {
	var num int
	for i := 0; i < count; i++ {
		fmt.Println("Input a number: ")
		fmt.Scan(&num)
		sum += num
		if i == 0 {
			min = num
		}
		if min > num {
			min = num
		}
		if num > max {
			max = num
		}
	}
	avg = float64(sum) / float64(count)

	return sum, min, max, avg
}
