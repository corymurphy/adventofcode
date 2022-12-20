package shared

func Abs(input int) int {
	if input < 0 {
		return input * -1
	} else {
		return input
	}
}
