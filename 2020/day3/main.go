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

	// I'm not proud of this, but I have my reasons :)

	horizontal := 0
	horizontal1 := 0
	horizontal3 := 0
	horizontal5 := 0
	horizontal7 := 0

	vertical := 0
	vertical2 := 0
	trees1 := 0
	trees3 := 0
	trees5 := 0
	trees7 := 0
	treesVert2 := 0

	for vertical < len(tobogganHill)-1 {

		vertical++
		horizontal1++
		horizontal3 = horizontal3 + 3
		horizontal5 = horizontal5 + 5
		horizontal7 = horizontal7 + 7

		if tobogganHill[vertical][horizontal1] {
			trees1++
		}

		if tobogganHill[vertical][horizontal3] {
			trees3++
		}

		if tobogganHill[vertical][horizontal5] {
			trees5++
		}

		if tobogganHill[vertical][horizontal7] {
			trees7++
		}

	}

	for vertical2 < len(tobogganHill)-1 {

		vertical2 = vertical2 + 2
		horizontal++

		if tobogganHill[vertical2][horizontal] {
			treesVert2++
		}

	}

	fmt.Printf("trees encountered horizontal1: %s\n", fmt.Sprint(trees1))
	fmt.Printf("trees encountered horizontal3: %s\n", fmt.Sprint(trees3))
	fmt.Printf("trees encountered horizontal5: %s\n", fmt.Sprint(trees5))
	fmt.Printf("trees encountered horizontal7: %s\n", fmt.Sprint(trees7))
	fmt.Printf("trees encountered vert2: %s\n", fmt.Sprint(treesVert2))

	final := trees1 * trees3 * trees5 * trees7 * treesVert2
	fmt.Printf("all trees multiplied: %s\n", fmt.Sprint(final))
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
		for i < 128 {
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
