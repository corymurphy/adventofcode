package main

import "fmt"

type Graph [][]string

func NewGraph(sensors Sensors) *Graph {
	graph := Graph{}

	for y := 0; y <= sensors.Max()+sensors.Trim(); y++ {
		row := []string{}
		for x := 0; x <= sensors.Max()+sensors.Trim(); x++ {
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

// func (c *Cave) Draw() {

// 	for y := c.Limits.MinY(); y <= c.Limits.maxY+2; y++ {
// 		// for y := c.Limits.MinY(); y <= 510; y++ {
// 		fmt.Println("")
// 		for x := c.Limits.MinX() - 10; x <= c.Limits.maxX+10; x++ {
// 			fmt.Printf("%s", (*c.Graph)[y][x])
// 		}
// 	}
// }

// func NewGraph(limits *Limits) *Graph {

// 	graph := Graph{}

// 	for y := 0; y <= limits.maxY+2; y++ {
// 		row := []string{}
// 		for x := 0; x <= limits.maxX+10000; x++ {
// 			row = append(row, ".")
// 		}
// 		graph = append((graph), row)
// 	}
// 	return &graph
// }
