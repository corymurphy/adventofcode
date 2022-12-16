package main

type Cave struct {
	Trim    int
	Sensors Sensors
	Graph   *Graph
}

func NewCave(input []string) *Cave {

	sensors := ParseNewSensors(input)
	trimmed := NewTrimmedSensors(sensors, sensors.Trim())

	graph := NewGraph(trimmed)
	graph.PlotSensors(trimmed)

	return &Cave{
		Trim:    Abs(sensors.Trim()),
		Sensors: trimmed,
		Graph:   graph,
	}
}

func (c *Cave) PlotArea() {
	// sensor := c.Sensors[0]
	// fmt.Println(sensor)
	// fmt.Println(sensor.A)
	for _, sensor := range c.Sensors {
		for _, coordinate := range sensor.Area() {
			// fmt.Println(coordinate)
			if c.Graph.Get(sensor.Location.Y+coordinate.Y, sensor.Location.X+coordinate.X) == "." {
				c.Graph.Set(sensor.Location.Y+coordinate.Y, sensor.Location.X+coordinate.X, "#")
			}
		}
	}

}

func (c *Cave) BeaconExclusionZones(line int) int {

	c.PlotArea()
	exclusion := 0

	for _, val := range (*c.Graph)[c.Trim+line] {
		if val == "#" {
			exclusion++
		}
	}

	return exclusion
}
