package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type BoardingPass struct {
	SeatCode string
}

type BoardingPasses []BoardingPass

func (b BoardingPasses) HighestSeatId() int {

	highest := 0
	for _, pass := range b {
		if pass.SeatId() > highest {
			highest = pass.SeatId()
		}
	}
	return highest
}

func NewBoardingPass(code string) BoardingPass {
	return BoardingPass{SeatCode: code}
}

func (b BoardingPass) SeatId() int {
	return (b.Row() * 8) + b.Column()
}

func (b BoardingPass) Row() int {
	rowCode := string(b.SeatCode[0 : len(b.SeatCode)-3])
	return row(rowCode, 0, 127, "F")
}

func (b BoardingPass) Column() int {
	colCode := string(b.SeatCode[7:len(b.SeatCode)])
	return row(colCode, 0, 7, "L")
}

func row(code string, lowerLimit int, upperLimit int, lowerIndicator string) int {

	section := string(code[0:1])

	if len(code) == 1 {

		if section == lowerIndicator {
			return lowerLimit
		} else {
			return upperLimit
		}
	}

	newCode := string(code[1:len(code)])

	diff := upperLimit - lowerLimit
	if section == lowerIndicator {
		newUpper := upperLimit - ((diff + (diff % 2)) / 2)
		return row(newCode, lowerLimit, newUpper, lowerIndicator)
	} else {
		newLower := lowerLimit + ((diff + (diff % 2)) / 2)
		return row(newCode, newLower, upperLimit, lowerIndicator)
	}

}

func part1() int {
	lines := readLines("input")

	passes := BoardingPasses{}

	for _, line := range lines {
		passes = append(passes, BoardingPass{SeatCode: line})
	}

	highest := passes.HighestSeatId()

	fmt.Println("Highest SeatID: " + fmt.Sprint(passes.HighestSeatId()))

	return highest
}

func part2() {

	lines := readLines("input")

	passes := BoardingPasses{}

	for _, line := range lines {
		passes = append(passes, BoardingPass{SeatCode: line})
	}

	highest := passes.HighestSeatId()

	availSeats := []int{}
	takenSeats := make(map[int]bool)

	mySeatId := 0

	i := 0
	for i <= highest {
		availSeats = append(availSeats, i)
		i++
	}

	for _, pass := range passes {
		takenSeats[pass.SeatId()] = true
	}

	for _, seat := range availSeats {
		if taken(takenSeats, seat-1) && !(taken(takenSeats, seat)) && taken(takenSeats, seat+1) {
			mySeatId = seat
			break
		}

	}

	fmt.Println("my seatid: " + fmt.Sprint(mySeatId))

}

func taken(takenSeats map[int]bool, seat int) bool {
	if _, ok := takenSeats[seat]; ok {
		return true
	} else {
		return false
	}
}

func main() {
	part1()
	part2()
}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
