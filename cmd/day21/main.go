package main

import (
	"aoc2023/pkg/files"
)

const Data = "data/day21"

type Position struct {
	X, Y int
}

type Positions []Position

func solve1(file string, steps int) int {
	garden := make(map[int]map[int]bool)
	positions := make(Positions, 1)
	width := 0
	height := 0
	if err := files.ReadLines(file, func(line string) {
		width = len(line)
		garden[height] = make(map[int]bool)
		for i, c := range line {
			if c == '#' {
				garden[height][i] = true
			} else if c == 'S' {
				positions[0] = Position{i, height}
			}
		}
		height++
	}); err != nil {
		panic(err)
	}

	for i := 0; i < steps; i++ {
		uniquePositions := make(map[Position]bool)
		for _, pos := range positions {
			if pos.Y > 0 {
				if up, ok := garden[pos.Y-1][pos.X]; !ok || !up {
					uniquePositions[Position{pos.X, pos.Y - 1}] = true
				}
			}
			if pos.Y < height-1 {
				if down, ok := garden[pos.Y+1][pos.X]; !ok || !down {
					uniquePositions[Position{pos.X, pos.Y + 1}] = true
				}
			}
			if pos.X > 0 {
				if left, ok := garden[pos.Y][pos.X-1]; !ok || !left {
					uniquePositions[Position{pos.X - 1, pos.Y}] = true
				}
			}
			if pos.X < width-1 {
				if right, ok := garden[pos.Y][pos.X+1]; !ok || !right {
					uniquePositions[Position{pos.X + 1, pos.Y}] = true
				}
			}
		}

		newPositions := make(Positions, 0, len(uniquePositions))
		for pos := range uniquePositions {
			newPositions = append(newPositions, pos)
		}

		positions = newPositions
	}
	return len(positions)
}

func pymod(a, b int) int {
	r := a % b
	if r < 0 {
		return r + b
	}
	return r
}

func solve2(file string, steps int) int {
	garden := make(map[int]map[int]bool)
	positions := make(map[int]map[int]bool)
	width := 0
	height := 0
	if err := files.ReadLines(file, func(line string) {
		width = len(line)
		garden[height] = make(map[int]bool)
		for i, c := range line {
			if c == '#' {
				garden[height][i] = true
			} else if c == 'S' {
				positions[height] = make(map[int]bool)
				positions[height][i] = true
			}
		}
		height++
	}); err != nil {
		panic(err)
	}

	// TODO: Optimize to no longer flip/flop oder positions that will just alternate every second step
	for i := 0; i < steps; i++ {
		newPositions := make(map[int]map[int]bool)
		for y, row := range positions {
			for x := range row {
				if v, ok := garden[pymod(y-1, height)][pymod(x, width)]; !ok || !v {
					if _, ok := newPositions[y-1]; !ok {
						newPositions[y-1] = make(map[int]bool)
					}
					newPositions[y-1][x] = true
				}
				if v, ok := garden[pymod(y+1, height)][pymod(x, width)]; !ok || !v {
					if _, ok := newPositions[y+1]; !ok {
						newPositions[y+1] = make(map[int]bool)
					}
					newPositions[y+1][x] = true
				}
				if v, ok := garden[pymod(y, height)][pymod(x-1, width)]; !ok || !v {
					if _, ok := newPositions[y]; !ok {
						newPositions[y] = make(map[int]bool)
					}
					newPositions[y][x-1] = true
				}
				if v, ok := garden[pymod(y, height)][pymod(x+1, width)]; !ok || !v {
					if _, ok := newPositions[y]; !ok {
						newPositions[y] = make(map[int]bool)
					}
					newPositions[y][x+1] = true
				}
			}
		}
		positions = newPositions
	}

	res := 0
	for _, row := range positions {
		res += len(row)
	}

	return res
}

func main() {
	print("Example 1: ", solve1(Data+"/example.txt", 6), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt", 64), "\n")

	print("Example 2a: ", solve2(Data+"/example.txt", 6), "\n")
	print("Example 2b: ", solve2(Data+"/example.txt", 10), "\n")
	print("Example 2c: ", solve2(Data+"/example.txt", 50), "\n")
	print("Example 2d: ", solve2(Data+"/example.txt", 100), "\n")
	print("Example 2e: ", solve2(Data+"/example.txt", 500), "\n")
	print("Example 2f: ", solve2(Data+"/example.txt", 1000), "\n")
	print("Example 2g: ", solve2(Data+"/example.txt", 5000), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt", 26501365), "\n")
}
