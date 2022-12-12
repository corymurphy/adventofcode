package main

func Move(grid *[][]int, instruction Instruction, rope *[]*Coordinates) {

	// fmt.Printf("move %s %d\n", instruction.Direction.String(), instruction.Distance)

	for moveCount := 0; moveCount < instruction.Distance; moveCount++ {

		// fmt.Printf("movecount: %d\n", moveCount)
		// direction := instruction.Direction

		direction := instruction.Direction
		for i := 1; i < len((*rope)); i++ {
			head := (*rope)[i-1]
			tail := (*rope)[i]

			if i == 1 {
				if isSamePosition(head, tail) {
					head.Move(instruction.Direction)
					// if (*grid)[tail.Y][tail.X] != 1 {
					// 	(*grid)[head.Y][head.X] = -1
					// }
					continue
				}

				head.Move(instruction.Direction)
				// if (*grid)[tail.Y][tail.X] != 1 {
				// 	(*grid)[head.Y][head.X] = -1
				// }
			}

			if isSamePosition(head, tail) {
				if i == len((*rope))-1 {
					(*grid)[tail.Y][tail.X] = 1
				}
				continue
			}

			if !isAdjacent(head, tail) {
				// fmt.Printf("move %s %d\n", instruction.Direction.String(), instruction.Distance)

				// fmt.Printf("move %s, head x: %d, head y: %d, tail # %d x: %d, tail y: %d\n", instruction.Direction.String(), head.X, head.Y, i, tail.X, tail.Y)
				tail.Follow(head, direction, i)
			}

			if i == len((*rope))-1 {
				(*grid)[tail.Y][tail.X] = 1
			}
		}
	}
}
