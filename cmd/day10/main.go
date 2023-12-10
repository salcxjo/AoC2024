package main

import (
	"aoc2023/pkg/files"
)

const Data = "data/day10"

const (
	None  = 0
	Any   = -1
	Up    = 1
	Down  = 2
	Left  = 4
	Right = 8
)

var tile_to_dir = map[rune]int{
	0:   None,
	'|': Up | Down,
	'-': Left | Right,
	'L': Up | Right,
	'J': Up | Left,
	'7': Down | Left,
	'F': Down | Right,
	'.': None,
	'S': Any,
}

type Position struct {
	X int
	Y int
}

type Location struct {
	Type     rune
	Dir      int
	Distance int
}

func solve1(file string) int {
	var data = make(map[int]map[int]rune)
	var start = Position{X: 0, Y: 0}
	y := 0
	if err := files.ReadLines(file, func(line string) {
		for x, c := range line {
			if _, ok := data[y]; !ok {
				data[y] = make(map[int]rune)
			}
			data[y][x] = c
			if c == 'S' {
				start.X = x
				start.Y = y
			}
		}
		y++
	}); err != nil {
		panic(err)
	}

	possible_start_dirs := None

	// Try up
	if up, ok := data[start.Y-1]; ok {
		if c, ok := up[start.X]; ok && tile_to_dir[c]&Down != 0 {
			possible_start_dirs |= Up
		}
	}

	// Try down
	if down, ok := data[start.Y+1]; ok {
		if c, ok := down[start.X]; ok && tile_to_dir[c]&Up != 0 {
			possible_start_dirs |= Down
		}
	}

	// Try left
	if c, ok := data[start.Y][start.X-1]; ok && tile_to_dir[c]&Right != 0 {
		possible_start_dirs |= Left
	}

	// Try right
	if c, ok := data[start.Y][start.X+1]; ok && tile_to_dir[c]&Left != 0 {
		possible_start_dirs |= Right
	}

	pos1 := start
	pos2 := start
	dir1 := None
	dir2 := None
	distance := 1
	first := true

	if possible_start_dirs&Up != 0 {
		pos1.Y--
		dir1 = Up
		first = false
	}
	if possible_start_dirs&Down != 0 {
		if first {
			pos1.Y++
			dir1 = Down
			first = false
		} else {
			pos2.Y++
			dir2 = Down
		}
	}
	if possible_start_dirs&Left != 0 {
		if first {
			pos1.X--
			dir1 = Left
			first = false
		} else {
			pos2.X--
			dir2 = Left
		}
	}
	if possible_start_dirs&Right != 0 {
		pos2.X++
		dir2 = Right
	}

	for pos1 != pos2 {
		distance++
		c1 := data[pos1.Y][pos1.X]
		c2 := data[pos2.Y][pos2.X]
		d1 := tile_to_dir[c1]
		d2 := tile_to_dir[c2]

		if dir1 != Right && d1&Left != 0 {
			pos1.X--
			dir1 = Left
		} else if dir1 != Left && d1&Right != 0 {
			pos1.X++
			dir1 = Right
		} else if dir1 != Down && d1&Up != 0 {
			pos1.Y--
			dir1 = Up
		} else if dir1 != Up && d1&Down != 0 {
			pos1.Y++
			dir1 = Down
		}

		if dir2 != Right && d2&Left != 0 {
			pos2.X--
			dir2 = Left
		} else if dir2 != Left && d2&Right != 0 {
			pos2.X++
			dir2 = Right
		} else if dir2 != Down && d2&Up != 0 {
			pos2.Y--
			dir2 = Up
		} else if dir2 != Up && d2&Down != 0 {
			pos2.Y++
			dir2 = Down
		}
	}
	return distance
}

var visited map[int]map[int]bool

// Very hacky, but whatever...
func flood_fill(pos Position, size Position) {
	if _, ok := visited[pos.Y]; !ok {
		visited[pos.Y] = make(map[int]bool)
	}
	if visited[pos.Y][pos.X] {
		return
	}
	visited[pos.Y][pos.X] = true
	if pos.X > 0 && !visited[pos.Y][pos.X-1] {
		flood_fill(Position{X: pos.X - 1, Y: pos.Y}, size)
	}
	if pos.X < size.X && !visited[pos.Y][pos.X+1] {
		flood_fill(Position{X: pos.X + 1, Y: pos.Y}, size)
	}
	if pos.Y > 0 {
		if up, ok := visited[pos.Y-1]; !ok || !up[pos.X] {
			flood_fill(Position{X: pos.X, Y: pos.Y - 1}, size)
		}
	}
	if pos.Y < size.Y {
		if down, ok := visited[pos.Y+1]; !ok || !down[pos.X] {
			flood_fill(Position{X: pos.X, Y: pos.Y + 1}, size)
		}
	}
}

func solve2(file string) int {
	var data = make(map[int]map[int]rune)
	var start = Position{X: 0, Y: 0}
	y := 0
	width := 0
	height := 0
	if err := files.ReadLines(file, func(line string) {
		for x, c := range line {
			if _, ok := data[y]; !ok {
				data[y] = make(map[int]rune)
			}
			data[y][x] = c
			if c == 'S' {
				start.X = x
				start.Y = y
			}
			width = len(line)
		}
		y++
	}); err != nil {
		panic(err)
	}
	height = y

	possible_start_dirs := None

	// This is essentially a flood fill task. But since the outside "moves" even between
	// parallel pipes, we can simply upscale the map by factor 2 to depict connections
	// and then flood fill everything from the outer borders.
	visited = make(map[int]map[int]bool)

	// Try up
	if up, ok := data[start.Y-1]; ok {
		if c, ok := up[start.X]; ok && tile_to_dir[c]&Down != 0 {
			possible_start_dirs |= Up
			visited[start.Y*2-1] = make(map[int]bool)
			visited[start.Y*2-1][start.X*2] = true
		}
	}

	// Try down
	if down, ok := data[start.Y+1]; ok {
		if c, ok := down[start.X]; ok && tile_to_dir[c]&Up != 0 {
			possible_start_dirs |= Down
			visited[start.Y*2+1] = make(map[int]bool)
			visited[start.Y*2+1][start.X*2] = true
		}
	}

	// Try left
	if c, ok := data[start.Y][start.X-1]; ok && tile_to_dir[c]&Right != 0 {
		possible_start_dirs |= Left
		visited[start.Y*2][start.X*2-1] = true
	}

	// Try right
	if c, ok := data[start.Y][start.X+1]; ok && tile_to_dir[c]&Left != 0 {
		possible_start_dirs |= Right
		if _, ok := visited[start.Y*2]; !ok {
			visited[start.Y*2] = make(map[int]bool)
		}
		visited[start.Y*2][start.X*2+1] = true
	}

	pos := start
	dir := None
	distance := 1

	if possible_start_dirs&Up != 0 {
		pos.Y--
		dir = Up
	} else if possible_start_dirs&Down != 0 {
		pos.Y++
		dir = Down
	} else {
		pos.X--
		dir = Left
	}

	//STOP

	if _, ok := visited[pos.Y*2]; !ok {
		visited[pos.Y*2] = make(map[int]bool)
	}
	visited[pos.Y*2][pos.X*2] = true

	for pos != start {
		distance++
		c1 := data[pos.Y][pos.X]
		d1 := tile_to_dir[c1]

		if dir != Right && d1&Left != 0 {
			if _, ok := visited[pos.Y*2]; !ok {
				visited[pos.Y*2] = make(map[int]bool)
			}
			visited[pos.Y*2][pos.X*2-1] = true
			pos.X--
			dir = Left
		} else if dir != Left && d1&Right != 0 {
			if _, ok := visited[pos.Y*2]; !ok {
				visited[pos.Y*2] = make(map[int]bool)
			}
			visited[pos.Y*2][pos.X*2+1] = true
			pos.X++
			dir = Right
		} else if dir != Down && d1&Up != 0 {
			if _, ok := visited[pos.Y*2-1]; !ok {
				visited[pos.Y*2-1] = make(map[int]bool)
			}
			visited[pos.Y*2-1][pos.X*2] = true
			pos.Y--
			dir = Up
		} else if dir != Up && d1&Down != 0 {
			if _, ok := visited[pos.Y*2+1]; !ok {
				visited[pos.Y*2+1] = make(map[int]bool)
			}
			visited[pos.Y*2+1][pos.X*2] = true
			pos.Y++
			dir = Down
		}

		if _, ok := visited[pos.Y*2]; !ok {
			visited[pos.Y*2] = make(map[int]bool)
		}
		visited[pos.Y*2][pos.X*2] = true
	}

	// Flood fill first row, last row, first column, and last column
	size := Position{X: width * 2, Y: height * 2}
	for x := 0; x < width*2; x++ {
		flood_fill(Position{X: x, Y: 0}, size)
		flood_fill(Position{X: x, Y: height*2 - 1}, size)
	}
	for y := 1; y < height*2-1; y++ {
		flood_fill(Position{X: 0, Y: y}, size)
		flood_fill(Position{X: width*2 - 1, Y: y}, size)
	}

	enclosed := 0
	for y := 0; y < height*2; y++ {
		if _, ok := visited[y]; ok {
			for x := 0; x < width*2; x++ {
				if _, ok := visited[y][x]; !ok {
					if x&1 == 0 && y&1 == 0 {
						enclosed++
					}
					// print(".")
				} else {
					// print("#")
				}
			}
		}
		// println()
	}

	return enclosed
}

func main() {
	print("Example 1a: ", solve1(Data+"/example1.txt"), "\n")
	print("Example 1b: ", solve1(Data+"/example2.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 2a: ", solve2(Data+"/example3.txt"), "\n")
	print("Example 2b: ", solve2(Data+"/example4.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
