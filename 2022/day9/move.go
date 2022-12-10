package main

func Move(grid *[][]int, instruction Instruction, head *Coordinates, tail *Coordinates) {

	for i := 0; i < instruction.Distance; i++ {

		if isSamePosition(head, tail) {
			head.Move(instruction.Direction)
			continue
		}

		head.Move(instruction.Direction)

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

func MoveTails(grid *[][]int, instruction Instruction, head *Coordinates, tails []*Coordinates) {

	for i := 0; i < instruction.Distance; i++ {

		if isSamePosition(head, tails[0]) {
			head.Move(instruction.Direction)
			continue
		}

		head.Move(instruction.Direction)

		for i, tail := range tails {
			if isSamePosition(head, tail) {
				if i == len(tails)-1 {
					(*grid)[tail.Y][tail.X] = 1
				}
				continue
			}

			if !isAdjacent(head, tail) {
				tail.Follow(head, instruction.Direction)
			}
			if i == len(tails)-1 {
				(*grid)[tail.Y][tail.X] = 1
			}
		}
	}
}
