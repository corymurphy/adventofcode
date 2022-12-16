package main

import (
	"strings"

	shared "github.com/corymurphy/adventofcode/shared"
)

type Sensors []Sensor

func ParseNewSensors(input []string) Sensors {
	sensors := NewSensors()
	for _, line := range input {

		// 2 sensor x
		// 3 sensor y

		// 8 beacon x
		// 9 beacon y

		data := strings.Split(line, " ")

		sensorX := data[2]
		sensorX = strings.TrimSuffix(sensorX, ",")
		sensorX = strings.TrimPrefix(sensorX, "x=")

		sensorY := data[3]
		sensorY = strings.TrimSuffix(sensorY, ":")
		sensorY = strings.TrimPrefix(sensorY, "y=")

		beaconX := data[8]
		beaconX = strings.TrimSuffix(beaconX, ",")
		beaconX = strings.TrimPrefix(beaconX, "x=")

		beaconY := data[9]
		beaconY = strings.TrimSuffix(beaconY, " ")
		beaconY = strings.TrimPrefix(beaconY, "y=")

		loc := NewCoordinates(shared.ToInt(sensorY), shared.ToInt(sensorX))
		beacon := NewCoordinates(shared.ToInt(beaconY), shared.ToInt(beaconX))

		sensor := NewSensor(loc, beacon)

		sensors.Add(sensor)
	}

	return sensors
}

func NewSensors() Sensors {
	return []Sensor{}
}

func NewTrimmedSensors(sensors Sensors, trim int) Sensors {
	trimmed := NewSensors()

	for _, sensor := range sensors {
		loc := sensor.Location
		beacon := sensor.Beacon

		loc.X = loc.X + trim
		loc.Y = loc.Y + trim

		beacon.X = beacon.X + trim
		beacon.Y = beacon.Y + trim

		trimmedSensor := NewSensor(loc, beacon)
		trimmed = append(trimmed, trimmedSensor)
	}
	return trimmed
}

func (s *Sensors) Add(new Sensor) {
	*s = append(*s, new)
}

func (s *Sensors) Min() int {
	lowest := 0
	for _, sensor := range *s {
		if sensor.Beacon.X < lowest {
			lowest = sensor.Beacon.X
		}
		if sensor.Beacon.Y < lowest {
			lowest = sensor.Beacon.Y
		}
		if sensor.Location.X < lowest {
			lowest = sensor.Location.X
		}
		if sensor.Location.Y < lowest {
			lowest = sensor.Location.Y
		}
	}
	return lowest
}

func (s *Sensors) Max() int {
	largest := 0
	for _, sensor := range *s {
		if sensor.Beacon.X > largest {
			largest = sensor.Beacon.X
		}
		if sensor.Beacon.Y > largest {
			largest = sensor.Beacon.Y
		}
		if sensor.Location.X > largest {
			largest = sensor.Location.X
		}
		if sensor.Location.Y > largest {
			largest = sensor.Location.Y
		}
	}
	return largest
}

func (s *Sensors) MaxDelta() int {
	largest := 0
	for _, sensor := range *s {
		delta := sensor.Delta()
		if delta > largest {
			largest = delta
		}
	}
	return largest
}

func (s *Sensors) Trim() int {
	min := s.Min()
	return Abs(min) + Abs(s.MaxDelta())
}
