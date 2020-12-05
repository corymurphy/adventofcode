package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func part1() {

	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	tobogganHill := initHill(lines)

	fmt.Println("row len: " + fmt.Sprint(len(tobogganHill[0])))
	fmt.Println("hill height: " + fmt.Sprint(len(tobogganHill)))

	startToboggan(tobogganHill)
}

func startToboggan(tobogganHill hill) hill {

	horizontal := 0
	vertical := 0
	trees := 0
	for vertical < len(tobogganHill)-1 {

		vertical++
		horizontal = horizontal + 3

		if tobogganHill[vertical][horizontal] {
			trees++
		}

	}

	fmt.Printf("trees encountered: %s\n", fmt.Sprint(trees))
	return tobogganHill
}

func main() {
	part1()
}

func printHill(h hill) {

	for _, horiz := range h {
		line := ""
		for _, plot := range horiz {
			line = line + parseTreeToString(plot)
		}
		fmt.Println(line)
	}
}

type hill [][]bool

type horizontal []bool

func (h hill) add(item []bool) {
	h = append(h, item)
}

func (h horizontal) add(item bool) {
	h = append(h, item)
}

func parseTreeToString(plot bool) string {
	if plot {
		return "#"
	} else {
		return "."
	}
}

func parseTree(plot string) bool {
	if plot == "#" {
		return true
	} else {
		return false
	}
}

func initHill(lines []string) hill {

	tobogganHill := hill{}
	for _, line := range lines {
		i := 0
		horiz := horizontal{}
		for i < 32 {
			for _, plot := range strings.Split(line, "") {
				horiz = append(horiz, parseTree(plot))
			}
			i++
		}
		tobogganHill = append(tobogganHill, horiz)
	}
	return tobogganHill
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
