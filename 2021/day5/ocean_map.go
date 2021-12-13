package main

import (
	"fmt"
	"strings"
)

type OceanMap struct {
	Map [1000][1000]int
}

func start(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func end(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func NewOceanMap(input []string) OceanMap {

	var plotMap [1000][1000]int

	oceanMap := OceanMap{
		Map: plotMap,
	}

	// y - outer , x - inner

	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			oceanMap.Map[y][x] = 0
		}
	}

	oceanMap.plotVentsPart1(input)

	return oceanMap
}

func (o *OceanMap) plotVentsPart1(input []string) {
	for _, raw := range input {
		vectors := strings.Split(raw, " -> ")
		startX := toInt(strings.Split(vectors[0], ",")[0])
		startY := toInt(strings.Split(vectors[0], ",")[1])
		endX := toInt(strings.Split(vectors[1], ",")[0])
		endY := toInt(strings.Split(vectors[1], ",")[1])

		if startX != endX && startY != endY {
			continue // only support straight lines
		} else if startX == endX {
			o.plotY(startY, endY, startX)
		} else if startY == endY {

			o.plotX(startX, endX, startY)
		}

	}
}

func (o *OceanMap) plotY(startY int, endY int, staticX int) {

	for i := start(startY, endY); i <= end(startY, endY); i++ {
		o.Map[staticX][i]++
	}

}

func (o *OceanMap) plotX(startX int, endX int, staticY int) {

	for i := start(startX, endX); i <= end(startX, endX); i++ {
		o.Map[i][staticY]++
	}
}

func (o *OceanMap) Part1() int {
	return o.CountDangerousAreas()
}

func (o *OceanMap) ToString() {
	for y := 0; y < 1000; y++ {
		row := ""
		for x := 0; x < 1000; x++ {
			if o.Map[y][x] == 0 {
				row = row + " ."
			} else {
				row = fmt.Sprintf("%s %d", row, o.Map[y][x])
			}
		}
		fmt.Println(row)
	}
}

func (o *OceanMap) CountDangerousAreas() int {

	count := 0

	for y := 0; y < len(o.Map); y++ {
		for x := 0; x < len(o.Map[y]); x++ {
			if o.Map[y][x] > 1 {
				count++
			}
		}
	}

	return count
}
