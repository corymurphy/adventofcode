package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1() int{
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	increased := 0

	for i, line := range lines {
		if i < 1 {
			continue
		}

		depth, _ := strconv.ParseInt(line, 10, 32)
		previousDepth, _ := strconv.ParseInt(lines[i-1], 10, 32)
		if depth > previousDepth {
			fmt.Printf("%s (increased)\n", line)
			increased++
		} else if depth == previousDepth {
			fmt.Printf("%s (no change)\n", line)
		} else {
			fmt.Printf("%s (decreased)\n", line)
		}
	}

	return increased
}

func part2() int {


	fmt.Println("")

	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	increased := 0

	for i, _ := range lines {
		if i < 2 {
			continue
		}

		currentWindowDepth := windowDepth(lines[i-2:i+1])
		
		if i == 2 {
			fmt.Printf("%d (n/a - no previous sum)", currentWindowDepth)
			continue
		}

		previousWindowDepth := windowDepth(lines[i-3:i])

		if currentWindowDepth > previousWindowDepth {
			fmt.Printf("%d (increased)\n", currentWindowDepth)
			increased++
		} else if currentWindowDepth == previousWindowDepth {
			fmt.Printf("%d (no change)\n", currentWindowDepth)
		} else {
			fmt.Printf("%d (decreased)\n", currentWindowDepth)
		}

		// if thisWindowDepth > previousWindowDepth {
		// 	increased++
		// }
	}

	fmt.Println("")

	return increased
}

func windowDepth(window []string) int32 {

	windowDepth := int32(0)
	
	for _, element := range window {
		depth, _ := strconv.ParseInt(element, 10, 32)
		windowDepth = windowDepth + int32(depth)
	}

	return windowDepth
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

func main() {
	part1 := part1()
	part2 := part2()

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}
