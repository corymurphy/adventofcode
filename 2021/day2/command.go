package main

import (
	"strconv"
	"strings"
)

type Command struct {
	Direction string
	Units     int32
}

func NewCommand(raw string) Command {
	result := strings.Split(raw, " ")
	value, _ := strconv.ParseInt(result[1], 10, 32)

	return Command{
		Direction: result[0],
		Units:     int32(value),
	}
}
