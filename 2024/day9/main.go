package main

import (
	"fmt"

	shared "github.com/corymurphy/adventofcode/shared"
)

func main() {
	input := shared.ReadInput("input_sample")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Printf("\nPart 1 answer: %d\n\n", part1)
	fmt.Printf("\nPart 2 answer: %d\n\n", part2)
}

type File struct {
	Id    int
	Size  int
	Start int
	Moved bool
}

type Free struct {
	Size  int
	Start int
}

type Block struct {
	IsFree bool
	Free   Free
	File   File
}

type Files []File
type Disk []Block

func (d *Disk) ReIndex(limit int) {

	for i := 0; i <= limit; i++ {

		block := (*d)[i]

		if !block.IsFree {
			continue
		}

		size := 1
		for {

			if i+size >= len((*d))-1 {
				break
			}

			if !(*d)[i+size].IsFree {
				break
			}

			(*d)[i+size].Free.Start = i

			size++
		}
		(*d)[i].Free.Start = i
		(*d)[i].Free.Size = size
	}
}

func (d *Disk) Compact() {

	end := len(*d) - 1

	for i := 0; i < len(*d); i++ {
		block := (*d)[i]

		if i >= end {
			return
		}

		if !block.IsFree {
			continue
		}

		freeBlock := (*d)[i]

		for {

			if (*d)[end].IsFree {
				end--
			} else {
				break
			}
		}

		fileBlock := (*d)[end]

		(*d)[i] = fileBlock
		(*d)[end] = freeBlock
	}
}

func (d *Disk) NeedsCompact() (fits bool) {

	for end := len(*d) - 1; end >= 0; end-- {
		block := (*d)[end]

		if block.IsFree {
			continue
		}

		for i := range len(*d) {

			if i >= end {
				break
			}
			if !(*d)[i].IsFree {
				continue
			}

			if (*d)[i].Free.Size >= block.File.Size && !block.File.Moved {
				return true
			}
		}
	}

	return false
}

func (d *Disk) Compact2() {

	for i := 0; i < len(*d); i++ {
		block := (*d)[i]

		if !block.IsFree {
			continue
		}

		end := len(*d) - 1
		for end >= 0 {
			if (*d)[end].IsFree {
				end--
				continue
			}

			if (*d)[end].File.Moved {
				end--
				continue
			}

		}

		// for end := len(*d) - 1; end >= 0; end-- {

		// 	freeBlock := (*d)[i]
		// 	fileBlock := (*d)[end]

		// 	if (*d)[end].IsFree {
		// 		continue
		// 	}

		// 	if freeBlock.Free.Size >= fileBlock.File.Size && !fileBlock.File.Moved {

		// 		for x := range fileBlock.File.Size {

		// 			if end-x <= i+x {
		// 				break
		// 			}

		// 			fileBlock = (*d)[end-x]
		// 			freeBlock = (*d)[i+x]

		// 			fileBlock.File.Moved = true

		// 			(*d)[i+x] = fileBlock
		// 			(*d)[end-x] = freeBlock
		// 		}

		// 		i = i + block.Free.Size - 1
		// 		d.ReIndex(0)
		// 		break
		// 	} else {
		// 		(*d)[end].File.Moved = true
		// 	}
		// 	(*d)[end].File.Moved = true

		// }
	}
}

func (d *Disk) Print() {

	fmt.Println()
	for _, block := range *d {
		if block.IsFree {
			fmt.Print(".")
		} else {
			fmt.Print(block.File.Id)
		}

	}
	fmt.Println()
}

func (d *Disk) Checksum() (checksum int) {
	for i, block := range *d {
		if block.IsFree {
			continue
		}

		checksum = checksum + (i * block.File.Id)
	}
	return checksum
}

func (d *Disk) Checksum2() (checksum int) {
	for i, block := range *d {
		if block.IsFree {
			continue
		}

		// result = i * block.File.Id
		result := i * block.File.Id

		if i > 0 && block.File.Id > 0 {
			if result < i || result < block.File.Id {
				fmt.Println(result, i, block.File.Id)

				panic("something happened")
			}

		}

		// result = checksum + result
		// if result < checksum {
		// 	panic("something happened")
		// }
		checksum = checksum + i*block.File.Id
	}
	return checksum
}

func part1(input []string) (answer int) {

	id := 0
	disk := Disk{}
	diskIndex := 0
	for i, char := range input[0] {

		size := shared.ToInt(string(char))

		block := Block{
			IsFree: i%2 > 0,
		}

		// even is file
		if i%2 == 0 {
			block.File = File{
				Id:    id,
				Size:  size,
				Start: diskIndex,
			}
			id++
		}

		for range size {
			disk = append(disk, block)
			diskIndex++
		}
	}

	disk.Compact()
	answer = disk.Checksum()

	return answer
}

func part2(input []string) (answer int) {

	id := 0
	disk := Disk{}
	diskIndex := 0
	for i, char := range input[0] {

		size := shared.ToInt(string(char))

		block := Block{
			IsFree: i%2 > 0,
		}

		// even is file
		if i%2 == 0 {
			block.File = File{
				Id:    id,
				Size:  size,
				Start: diskIndex,
			}
			id++
		}

		if i%2 > 0 {
			block.Free = Free{
				Size:  size,
				Start: diskIndex,
			}
		}

		for range size {
			disk = append(disk, block)
			diskIndex++
		}
	}

	for {
		if !disk.NeedsCompact() {
			break
		}
		fmt.Println("compacting")
		disk.Compact2()
	}

	// disk.Compact2()

	answer = disk.Checksum2()
	disk.Print()

	// disk.Print()

	// fmt.Pri

	return answer
}
