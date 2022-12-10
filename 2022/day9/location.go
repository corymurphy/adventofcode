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

	return xDiff <= 1 && xDiff >= -1 && yDiff <= 1 && yDiff >= -1
}
