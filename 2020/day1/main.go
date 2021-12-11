package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseInt(num string) int64 {
	if val, err := strconv.ParseInt(num, 10, 64); err != nil {
		log.Fatalf("error parsing int: %s", err)
		return 0
	} else {
		return val
	}
}

func part1() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	processed := make(map[int64]int)

	for i, line := range lines {

		num := parseInt(line)

		if diffIndex, ok := processed[num]; ok {
			pair := parseInt(lines[diffIndex])
			fmt.Printf("part 1 answer: %s\n", fmt.Sprint(pair*num))
			return
		}

		processed[2020-num] = i
	}
}

func part2() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for _, int1Str := range lines {
		for _, int2Str := range lines {
			for _, int3Str := range lines {
				if (parseInt(int1Str) + parseInt(int2Str) + parseInt(int3Str)) == 2020 {
					fmt.Printf("part 2 answer: %s\n", fmt.Sprint(parseInt(int1Str)*parseInt(int2Str)*parseInt(int3Str)))
					return
				}
			}
		}
	}
}

func main() {
	part1()
	part2()
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
