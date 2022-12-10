package main

type Coordinates struct {
	X int
	Y int
}

func NewCoordinates() *Coordinates {
	return &Coordinates{
		X: 20,
		Y: 20,
	}
}

func (c *Coordinates) Follow(leader *Coordinates, direction Direction) {
	if direction == Right {
		c.X = leader.X - 1
		c.Y = leader.Y
	}

	if direction == Left {
		c.X = leader.X + 1
		c.Y = leader.Y
	}

	if direction == Up {
		c.Y = leader.Y - 1
		c.X = leader.X
	}

	if direction == Down {
		c.Y = leader.Y + 1
		c.X = leader.X
	}
}

func (c *Coordinates) Move(direction Direction) {

	if direction == Right {
		c.Right()
	}

	if direction == Left {
		c.Left()
	}

	if direction == Up {
		c.Up()
	}

	if direction == Down {
		c.Down()
	}

}

func (c *Coordinates) Right() {
	c.X = c.X + 1
}

func (c *Coordinates) Left() {
	c.X = c.X - 1
}

func (c *Coordinates) Up() {
	c.Y = c.Y + 1
}

func (c *Coordinates) Down() {
	c.Y = c.Y - 1
}

func NewCoordinatesList(count int) []*Coordinates {
	cords := []*Coordinates{}
	for i := 0; i < count; i++ {
		cords = append(cords, NewCoordinates())
	}
	return cords
}
