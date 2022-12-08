package main

func viewLeft(startColumn int, startRow int, landscape [][]int) int {
	score := 0

	tree := landscape[startColumn][startRow]

	for i := startRow - 1; i >= 0; i-- {
		height := landscape[startColumn][i]

		if tree > height {
			score++
			continue
		}

		if tree <= height {
			score++
			break
		}
	}

	return score
}

func viewRight(startColumn int, startRow int, landscape [][]int) int {
	score := 0
	tree := landscape[startColumn][startRow]
	for i := startRow + 1; i <= len(landscape)-1; i++ {
		height := landscape[startColumn][i]

		if tree > height {
			score++
			continue
		}

		if tree <= height {
			score++
			break
		}
	}
	return score
}

func viewBottom(startColumn int, startRow int, landscape [][]int) int {
	score := 0
	tree := landscape[startColumn][startRow]

	for i := startColumn + 1; i <= len(landscape)-1; i++ {
		height := landscape[i][startRow]
		if tree > height {
			score++
			continue
		}

		if tree <= height {
			score++
			break
		}
	}
	return score
}

func viewTop(startColumn int, startRow int, landscape [][]int) int {
	score := 0
	tree := landscape[startColumn][startRow]
	// 5 (3,2)
	for i := startColumn - 1; i >= 0; i-- {
		height := landscape[i][startRow]
		// score++

		if tree > height {
			score++
			continue
		}

		if tree <= height {
			score++
			break
		}
		// if height <= tree {
		// 	score++
		// 	break
		// }

		// if height <= tree {
		// 	score++
		// 	continue
		// }
	}
	return score
}
