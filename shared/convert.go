package main

import "strconv"

func toInt(input string) int {
	value, _ := strconv.ParseInt(input, 10, 32)
	return int(value)
}
