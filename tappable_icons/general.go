package tappable_icons

func intRangeCheck(num int, max int, min int) int {
	if num > max {
		num = max
	} else if num < min {
		num = min
	}
	return num
}
