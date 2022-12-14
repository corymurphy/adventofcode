package main

type Limits struct {
	maxY  int
	maxX  int
	minY  int
	minX  int
	xMins []int
	yMins []int
}

func NewLimits() *Limits {
	return &Limits{
		maxX:  0,
		maxY:  0,
		xMins: []int{},
		yMins: []int{},
	}
}

func (l *Limits) Analyze(vectors []*Vector) {

	for _, vector := range vectors {
		if vector.X > l.maxX {
			l.maxX = vector.X + 2
		}

		if vector.Y > l.maxY {
			l.maxY = vector.Y + 2
		}

		l.xMins = append(l.xMins, vector.X)
		l.yMins = append(l.yMins, vector.Y)
	}
}

func (l *Limits) MinX() int {
	min := l.xMins[0]

	for _, val := range l.xMins {
		if val < min {
			min = val
		}
	}
	return min
}

func (l *Limits) MinY() int {
	min := l.yMins[0]

	for _, val := range l.yMins {
		if val < min {
			min = val
		}
	}
	return min
}
