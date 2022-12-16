package main

type Sensor struct {
	Location Coordinate
	Beacon   Coordinate
}

func NewSensor(loc Coordinate, beacon Coordinate) Sensor {
	return Sensor{
		Location: loc,
		Beacon:   beacon,
	}
}

func (s *Sensor) Delta() int {
	dY := s.Location.Y - s.Beacon.Y
	dX := s.Location.X - s.Beacon.X
	return Abs(dY) + Abs(dX)
}

// func sensor
func (s *Sensor) Area() []Coordinate {
	area := []Coordinate{}

	// delta := s.Delta()

	for delta := s.Delta(); delta >= 0; delta-- {
		for i := 0; i <= delta; i++ {
			c := NewCoordinates(i, delta-i)
			area = append(area, c)
		}

		for i := 0; i <= delta; i++ {
			c := NewCoordinates(i*-1, delta-i)
			area = append(area, c)
		}

		for i := 0; i <= delta; i++ {
			c := NewCoordinates(i, (delta-i)*-1)
			area = append(area, c)
		}

		for i := 0; i <= delta; i++ {
			c := NewCoordinates(i*-1, (delta-i)*-1)
			area = append(area, c)
		}
	}

	return area
}
