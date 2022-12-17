package main

import "fmt"

type Graph [][]string

func NewGraph(sensors Sensors) *Graph {
	graph := Graph{}

	trim := sensors.Trim()
	max := sensors.Max()
	min := sensors.Min()

	fmt.Println(max + trim)
	fmt.Println(min)
	for y := min; y <= max+trim; y++ {
		row := []string{}
		for x := min; x <= max+trim; x++ {
			row = append(row, ".")
		}
		graph = append(graph, row)
	}
	return &graph
}

func (g *Graph) Set(y int, x int, val string) {
	(*g)[y][x] = val
}

func (g *Graph) Get(y int, x int) string {
	return (*g)[y][x]
}

func (g *Graph) Draw(min int, max int, trim int) {
	for y := min; y <= max; y++ {
		fmt.Println()
		fmt.Printf("%d (%d) ", y, y-trim)
		for x := min; x <= max; x++ {
			fmt.Printf("%s", g.Get(y, x))
		}
	}
}

func (g *Graph) PlotSensors(s Sensors) {
	for _, sensor := range s {
		g.Set(sensor.Location.Y, sensor.Location.X, "S")
		g.Set(sensor.Beacon.Y, sensor.Beacon.X, "B")
	}
}
