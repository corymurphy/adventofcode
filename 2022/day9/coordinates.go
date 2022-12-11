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

func (follower *Coordinates) Follow(leader *Coordinates, direction Direction) {
	if direction == Right {
		follower.X = leader.X - 1
		// follower.X = follower.X - 1

		// if is leader above
		if leader.Y > follower.Y {
			follower.Y = follower.Y + 1
			// return Up
		}

		// if is leader below
		if leader.Y < follower.Y {
			follower.Y = follower.Y - 1
			// return Down
		}
		// return Right
	}

	if direction == Left {
		follower.X = leader.X + 1
		// follower.X = follower.X + 1

		// if is leader above
		if leader.Y > follower.Y {
			follower.Y = follower.Y + 1
			// return Up
		}

		// if is leader below
		if leader.Y < follower.Y {
			follower.Y = follower.Y - 1
			// return Down
		}
	}

	if direction == Up {
		follower.Y = leader.Y - 1
		// follower.Y = follower.Y - 1

		// if is leader left
		if leader.X > follower.X {
			follower.X = follower.X + 1
			// return Left
		}

		// if leader is right
		if leader.X < follower.X {
			follower.X = follower.X - 1
			// return Right
		}

	}

	if direction == Down {
		follower.Y = leader.Y + 1
		// follower.Y = follower.Y + 1
		// follower.X = leader.X

		// if is leader left
		if leader.X > follower.X {
			follower.X = follower.X + 1
			// return Left
		}

		// if leader is right
		if leader.X < follower.X {
			follower.X = follower.X - 1
			// return Right
		}

	}

	// return direction
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
