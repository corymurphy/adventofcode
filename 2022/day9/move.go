package main

import "fmt"

// const (
// 	up    int = 1
// 	down  int = -1
// 	right int = 1
// 	left  int = -1
// )

// I could probably map the enum to -/+ on the grid, but i'm lazy
func Move(grid *[][]int, instruction Instruction, head *Coordinates, tail *Coordinates) {

	// fmt.Printf("move %s %d\n", instruction.Direction.String(), instruction.Distance)

	if instruction.Direction == Right || instruction.Direction == Left {
		MoveHorizontal(grid, instruction, head, tail)
	}
	// if instruction.Direction == Right {
	// 	MoveRight(grid, instruction, head, tail)
	// 	// move(grid, instruction, head, tail, right)
	// 	return
	// }

	// if instruction.Direction == Left {
	// 	MoveLeft(grid, instruction, head, tail)
	// 	return
	// }

	if instruction.Direction == Up {
		MoveHorizontal(grid, instruction, head, tail)
		return
	}

	if instruction.Direction == Down {
		MoveHorizontal(grid, instruction, head, tail)
		return
	}

	if instruction.Direction == Unknown {
		fmt.Println("unknown")
	}
}

func MoveUp(grid *[][]int, instruction Instruction, head *Coordinates, tail *Coordinates) {

	for i := 0; i < instruction.Distance; i++ {

		// fmt.Printf("head x: %d, head y: %d tail x: %d, tail y: %d\n", head.X, head.Y, tail.X, tail.Y)

		if isTopEdge(grid, head) {
			fmt.Println("exceeding top boundaries, invalid instruction")
			continue
		}

		if isSamePosition(head, tail) {
			(*grid)[tail.Y][tail.X] = 1
			head.Y = head.Y + 1
			continue
		}

		head.Y = head.Y + 1

		if isSamePosition(head, tail) {
			(*grid)[tail.Y][tail.X] = 1
			continue
		}

		if !isAdjacent(head, tail) {
			tail.Y = head.Y - 1
			tail.X = head.X
		}
		(*grid)[tail.Y][tail.X] = 1
	}
	// fmt.Printf("head x: %d, head y: %d tail x: %d, tail y: %d\n", head.X, head.Y, tail.X, tail.Y)
}

func MoveDown(grid *[][]int, instruction Instruction, head *Coordinates, tail *Coordinates) {
	// fmt.Printf("move %s %d\n", instruction.Direction.String(), instruction.Distance)

	for i := 0; i < instruction.Distance; i++ {

		// fmt.Printf("head x: %d, head y: %d tail x: %d, tail y: %d\n", head.X, head.Y, tail.X, tail.Y)

		if isBottomEdge(grid, head) {
			fmt.Println("exceeding bottom boundaries, invalid instruction")
			// fmt.Printf("head x: %d, head y: %d tail x: %d, tail y: %d\n", head.X, head.Y, tail.X, tail.Y)
			continue
		}

		if isSamePosition(head, tail) {
			(*grid)[tail.Y][tail.X] = 1
			head.Y = head.Y - 1
			continue
		}

		head.Y = head.Y - 1

		if isSamePosition(head, tail) {
			(*grid)[tail.Y][tail.X] = 1
			continue
		}

		if !isAdjacent(head, tail) {
			// fmt.Println("not adjacent head ")
			tail.Y = head.Y + 1
			tail.X = head.X
		}

		(*grid)[tail.Y][tail.X] = 1
	}
}

func MoveLeft(grid *[][]int, instruction Instruction, head *Coordinates, tail *Coordinates) {

	for i := 0; i < instruction.Distance; i++ {

		// fmt.Printf("head x: %d, head y: %d tail x: %d, tail y: %d\n", head.X, head.Y, tail.X, tail.Y)

		if isLeftEdge(grid, head) {

			fmt.Println("exceeding left boundaries, invalid instruction")
			continue
		}

		if isSamePosition(head, tail) {
			(*grid)[tail.Y][tail.X] = 1
			head.X = head.X - 1
			continue
		}

		head.X = head.X - 1

		if isSamePosition(head, tail) {
			(*grid)[tail.Y][tail.X] = 1
			continue
		}

		if !isAdjacent(head, tail) {
			// fmt.Printf("not adjacent head x: %d, head y: %d tail x: %d, tail y: %d\n", head.X, head.Y, tail.X, tail.Y)
			tail.X = head.X + 1
			tail.Y = head.Y
		}
		(*grid)[tail.Y][tail.X] = 1
	}
}

func MoveRight(grid *[][]int, instruction Instruction, head *Coordinates, tail *Coordinates) {

	// fmt.Printf("move %s %d\n", instruction.Direction.String(), instruction.Distance)

	for i := 0; i < instruction.Distance; i++ {

		// fmt.Printf("head x: %d, head y: %d tail x: %d, tail y: %d\n", head.X, head.Y, tail.X, tail.Y)

		if isRightEdge(grid, head) {
			fmt.Println("exceeding right boundaries, invalid instruction")
			// (*grid)[head.Y] = append((*grid)[head.Y], 1)
			continue
		}

		if isSamePosition(head, tail) {
			(*grid)[tail.Y][tail.X] = 1
			head.X = head.X + 1
			continue
		}

		head.X = head.X + 1

		if isSamePosition(head, tail) {
			(*grid)[tail.Y][tail.X] = 1
			continue
		}

		if !isAdjacent(head, tail) {
			tail.X = head.X - 1
			tail.Y = head.Y
		}
		(*grid)[tail.Y][tail.X] = 1

	}
}

func MoveHorizontal(grid *[][]int, instruction Instruction, head *Coordinates, tail *Coordinates) {

	for i := 0; i < instruction.Distance; i++ {

		if isSamePosition(head, tail) {
			(*grid)[tail.Y][tail.X] = 1
			head.Move(instruction.Direction)
			continue
		}

		head.Move(instruction.Direction)

		// move tails

		if isSamePosition(head, tail) {
			(*grid)[tail.Y][tail.X] = 1
			continue
		}

		if !isAdjacent(head, tail) {
			tail.Follow(head, instruction.Direction)
		}
		(*grid)[tail.Y][tail.X] = 1

	}
}

func MoveDynamic(grid *[][]int, instruction Instruction, head *Coordinates, tails []*Coordinates) {

	// if instruction.Direction == Right {
	// 	MoveRight(grid, instruction, head, tail)
	// 	move(grid, instruction, head, tail, right)
	// 	return
	// }

	// if instruction.Direction == Left {
	// 	MoveLeft(grid, instruction, head, tail)
	// 	return
	// }

	// if instruction.Direction == Up {
	// 	MoveUp(grid, instruction, head, tail)
	// 	return
	// }

	// if instruction.Direction == Down {
	// 	MoveDown(grid, instruction, head, tail)
	// 	return
	// }
}
