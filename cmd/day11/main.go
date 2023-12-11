package main

import (
	"aoc2023/pkg/files"
)

const Data = "data/day11"

type Galaxy struct {
	X, Y int
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func ManhattanDistance(a, b Galaxy) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}

func solve(file string, distance_scale int) int {
	_galaxies := make(map[int]map[int]bool)
	galaxies := make([]Galaxy, 0)
	width := 0
	y := 0
	if err := files.ReadLines(file, func(line string) {
		count := 0
		width = len(line)
		for x, c := range line {
			if c == '#' {
				count++
				if _, ok := _galaxies[x]; !ok {
					_galaxies[x] = make(map[int]bool)
				}
				_galaxies[x][y] = true
			}
		}
		if count == 0 {
			y += distance_scale
		} else {
			y++
		}
	}); err != nil {
		panic(err)
	}
	offset := 0
	for x := 0; x < width; x++ {
		if col, ok := _galaxies[x]; !ok {
			offset += distance_scale - 1
			continue
		} else {
			for y := range col {
				galaxies = append(galaxies, Galaxy{
					X: x + offset,
					Y: y,
				})
			}
		}

	}
	res := 0

	for i, g1 := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			res += ManhattanDistance(g1, galaxies[j])
		}
	}

	return res
}

func main() {
	print("Example 1: ", solve(Data+"/example.txt", 2), "\n")
	print("Solution 1: ", solve(Data+"/input.txt", 2), "\n")

	print("Example 2a: ", solve(Data+"/example.txt", 10), "\n")
	print("Example 2b: ", solve(Data+"/example.txt", 100), "\n")
	print("Solution 2: ", solve(Data+"/input.txt", 1000000), "\n")
}
