package main

import "fmt"

type Cave struct {
	Trim    int
	Sensors Sensors
	Graph   *Graph
}

func NewCave(input []string) *Cave {

	// sensors := ParseNewSensors(input)
	// trimmed := NewTrimmedSensors(sensors, sensors.Trim())
	trimmed := ParseNewSensors(input)

	fmt.Println("sensors parsed")

	graph := NewGraph(trimmed)

	fmt.Println("graph created")

	graph.PlotSensors(trimmed)

	return &Cave{
		Trim:    Abs(0),
		Sensors: trimmed,
		Graph:   graph,
	}
}

func (c *Cave) PlotArea(line int) {
	// sensor := c.Sensors[0]
	// fmt.Println(sensor)
	// fmt.Println(sensor.A)
	maxDelta := c.Sensors.MaxDelta()

	// fmt.Println(maxDelta)

	fmt.Println(len(c.Sensors))
	for _, sensor := range c.Sensors {

		if sensor.Location.Y > line+maxDelta && sensor.Location.Y < line-maxDelta {
			continue
		}

		for _, coordinate := range sensor.Area() {
			// fmt.Println(coordinate)
			if c.Graph.Get(sensor.Location.Y+coordinate.Y, sensor.Location.X+coordinate.X) == "." {
				c.Graph.Set(sensor.Location.Y+coordinate.Y, sensor.Location.X+coordinate.X, "#")
			}
		}
	}

}

func (c *Cave) BeaconExclusionZones(line int) int {

	c.PlotArea(line)
	exclusion := 0

	for _, val := range (*c.Graph)[c.Trim+line] {
		if val == "#" {
			exclusion++
		}
	}

	c.Graph.Draw(c.Sensors.Min(), c.Sensors.Max(), c.Trim)

	return exclusion
}
