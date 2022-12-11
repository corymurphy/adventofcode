package main

func isSamePosition(head *Coordinates, tail *Coordinates) bool {
	return head.X == tail.X && head.Y == tail.Y
}

func isSameRow(head *Coordinates, tail *Coordinates) bool {
	return head.Y == tail.Y
}

func isRightEdge(grid *[][]int, head *Coordinates) bool {
	return head.X >= len(*grid)-1
}

func isLeftEdge(grid *[][]int, head *Coordinates) bool {
	return head.X == 0
}

func isTopEdge(grid *[][]int, head *Coordinates) bool {
	return head.Y >= len(*grid)-1
	// return head.Y == 0
}

func isBottomEdge(grid *[][]int, head *Coordinates) bool {
	return head.Y == 0
}

func isAdjacent(head *Coordinates, tail *Coordinates) bool {
	xDiff := head.X - tail.X
	yDiff := head.Y - tail.Y

	return isLessThanEqualOne(xDiff) && isLessThanEqualOne(yDiff)
}

func isLessThanEqualOne(value int) bool {
	return value <= 1 && value >= -1
}

func isAdjacentVertical(head *Coordinates, tail *Coordinates) bool {
	return isLessThanEqualOne(head.Y - tail.Y)
}

func isAdjacentHorizontal(head *Coordinates, tail *Coordinates) bool {
	return isLessThanEqualOne(head.X - tail.Y)
}

func isAdjacentAngular(head *Coordinates, tail *Coordinates) bool {
	return isAdjacent(head, tail) &&
		!isAdjacentHorizontal(head, tail) &&
		!isAdjacentVertical(head, tail)
}
