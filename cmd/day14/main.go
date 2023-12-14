package main

import (
	"aoc2023/pkg/files"
)

const Data = "data/day14"

const (
	Empty   = 0
	Fixed   = 1
	Rolling = 2
)

type Board struct {
	data   [][]int
	width  int
	height int
}

func (b *Board) append(line string) {
	b.width = len(line)
	b.data = append(b.data, make([]int, b.width))
	for x, c := range line {
		switch c {
		case '.':
			b.data[b.height][x] = Empty
		case '#':
			b.data[b.height][x] = Fixed
		case 'O':
			b.data[b.height][x] = Rolling
		}
	}
	b.height++
}

func (b *Board) print() {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			switch b.data[y][x] {
			case Empty:
				print(".")
			case Fixed:
				print("#")
			case Rolling:
				print("O")
			}
		}
		print("\n")
	}
}

func (b *Board) tiltNorth() {
	for x := 0; x < b.width; x++ {
		for y := 1; y < b.height; y++ {
			if b.data[y][x] == Rolling {
				for o := 1; o <= y; o++ {
					if b.data[y-o][x] == Empty {
						b.data[y-o][x] = Rolling
						b.data[y+1-o][x] = Empty
					} else {
						break
					}
				}
			}
		}
	}
}

func (b *Board) tiltWest() {
	for y := 0; y < b.height; y++ {
		for x := 1; x < b.width; x++ {
			if b.data[y][x] == Rolling {
				for o := 1; o <= x; o++ {
					if b.data[y][x-o] == Empty {
						b.data[y][x-o] = Rolling
						b.data[y][x+1-o] = Empty
					} else {
						break
					}
				}
			}
		}
	}
}

func (b *Board) tiltSouth() {
	for x := 0; x < b.width; x++ {
		for y := b.height - 1; y >= 0; y-- {
			if b.data[y][x] == Rolling {
				for o := 1; o < b.height-y; o++ {
					if b.data[y+o][x] == Empty {
						b.data[y+o][x] = Rolling
						b.data[y-1+o][x] = Empty
					} else {
						break
					}
				}
			}
		}
	}
}

func (b *Board) tiltEast() {
	for y := 0; y < b.height; y++ {
		for x := b.width - 1; x >= 0; x-- {
			if b.data[y][x] == Rolling {
				for o := 1; o < b.width-x; o++ {
					if b.data[y][x+o] == Empty {
						b.data[y][x+o] = Rolling
						b.data[y][x-1+o] = Empty
					} else {
						break
					}
				}
			}
		}
	}
}

func (b *Board) tiltCycle() {
	b.tiltNorth()
	b.tiltWest()
	b.tiltSouth()
	b.tiltEast()
}

func (b *Board) loadNorth() int {
	res := 0
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			if b.data[y][x] == Rolling {
				res += b.height - y
			}
		}
	}
	return res
}

func solve1(file string) int {
	board := Board{}
	if err := files.ReadLines(file, func(line string) {
		board.append(line)
	}); err != nil {
		panic(err)
	}

	board.tiltNorth()
	return board.loadNorth()
}

func findPeriodity(loads []int) int {
	positions := make(map[int][]int)
	maxPeriod := 0
	for i, load := range loads {
		if pos, ok := positions[load]; ok {
			period := i - pos[len(pos)-1]
			if period > maxPeriod {
				maxPeriod = period
			}
		}
		positions[load] = append(positions[load], i)
	}
	if maxPeriod == len(loads)+1 {
		return -1
	}
	return maxPeriod
}

func solve2(file string) int {
	board := Board{}
	if err := files.ReadLines(file, func(line string) {
		board.append(line)
	}); err != nil {
		panic(err)
	}

	// Tilt a few times just in case there's no periodity at the start
	for i := 0; i < 1000; i++ {
		board.tiltCycle()
	}

	// Record the next 1000 cycles, these should be enough to find periodity
	loads := make([]int, 0)
	for i := 0; i < 1000; i++ {
		board.tiltCycle()
		loads = append(loads, board.loadNorth())
	}

	// Find the periodity
	periodity := findPeriodity(loads)

	// Now let's just skip the cycles that change nothing
	remaining_cycles := (1000000000 - 2000) % periodity

	for i := 0; i < remaining_cycles; i++ {
		board.tiltCycle()
	}

	/*
		min_load := -1
		for i := 0; i < 1000000000; i++ {
			board.tiltCycle()
			load := board.loadNorth()
			if min_load == -1 || load < min_load {
				min_load = load
			}
			print("\bCycle ", i, " load ", load, "...\n")
		}
	*/

	return board.loadNorth()
}

func main() {
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 2: ", solve2(Data+"/example.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
