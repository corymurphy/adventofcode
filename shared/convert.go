package shared

import "strconv"

func ToInt(input string) int {
	value, _ := strconv.ParseInt(input, 10, 32)
	return int(value)
}

func MustInt64(input string) int64 {
	val, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		panic("unable to convert")
	}
	return val
}
