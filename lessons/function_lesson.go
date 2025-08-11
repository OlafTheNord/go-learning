package lessons

func IsEven(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else if a < b {
		return b
	} else {
		return a
	}
}
