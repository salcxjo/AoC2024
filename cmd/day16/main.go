package main

import (
	"aoc2023/pkg/files"
)

const Data = "data/day16"

const (
	Empty     = 0
	Mirror1   = 1
	Mirror2   = 2
	SplitterV = 3
	SplitterH = 4
)

const (
	None  = 0
	Up    = 1
	Down  = 2
	Left  = 4
	Right = 8
)

type Contraption struct {
	Width  int
	Height int
	Board  map[int]map[int]int
	Energy map[int]map[int]int
}

func (c *Contraption) Energize(x, y, dir int, start bool) {
	if start {
		c.Energy = make(map[int]map[int]int)
	}

	if x < 0 || x >= c.Width || y < 0 || y >= c.Height {
		return // Beam left the contraption
	}

	row, ok := c.Energy[y]
	if !ok {
		row = make(map[int]int)
		c.Energy[y] = row
	}
	if v, ok := c.Energy[y][x]; ok && v&dir != 0 {
		return // Beam already passed through here in this direction
	}
	c.Energy[y][x] |= dir
	switch c.Board[y][x] {
	case Empty:
		switch dir {
		case Up:
			c.Energize(x, y-1, Up, false)
		case Down:
			c.Energize(x, y+1, Down, false)
		case Left:
			c.Energize(x-1, y, Left, false)
		case Right:
			c.Energize(x+1, y, Right, false)
		}
	case Mirror1: // /
		switch dir {
		case Up:
			c.Energize(x+1, y, Right, false)
		case Down:
			c.Energize(x-1, y, Left, false)
		case Left:
			c.Energize(x, y+1, Down, false)
		case Right:
			c.Energize(x, y-1, Up, false)
		}
	case Mirror2: // \
		switch dir {
		case Up:
			c.Energize(x-1, y, Left, false)
		case Down:
			c.Energize(x+1, y, Right, false)
		case Left:
			c.Energize(x, y-1, Up, false)
		case Right:
			c.Energize(x, y+1, Down, false)
		}
	case SplitterV:
		if dir == Up {
			c.Energize(x, y-1, Up, false)
		} else if dir == Down {
			c.Energize(x, y+1, Down, false)
		} else {
			c.Energize(x, y-1, Up, false)
			c.Energize(x, y+1, Down, false)
		}
	case SplitterH:
		if dir == Left {
			c.Energize(x-1, y, Left, false)
		} else if dir == Right {
			c.Energize(x+1, y, Right, false)
		} else {
			c.Energize(x-1, y, Left, false)
			c.Energize(x+1, y, Right, false)
		}
	}
}

func (c *Contraption) Measure() int {
	res := 0
	for _, row := range c.Energy {
		for range row {
			res++
		}
	}
	return res
}

func solve1(file string) int {
	contraption := Contraption{
		Width:  0,
		Height: 0,
		Board:  make(map[int]map[int]int),
		Energy: make(map[int]map[int]int),
	}
	if err := files.ReadLines(file, func(line string) {
		contraption.Board[contraption.Height] = make(map[int]int)
		for i, c := range line {
			switch c {
			case '.':
				// contraption.Board[contraption.Height][i] = Empty
			case '/':
				contraption.Board[contraption.Height][i] = Mirror1
			case '\\':
				contraption.Board[contraption.Height][i] = Mirror2
			case '|':
				contraption.Board[contraption.Height][i] = SplitterV
			case '-':
				contraption.Board[contraption.Height][i] = SplitterH
			}
		}
		contraption.Width = len(line)
		contraption.Height++
	}); err != nil {
		panic(err)
	}
	contraption.Energize(0, 0, Right, true)
	return contraption.Measure()
}

func solve2(file string) int {
	contraption := Contraption{
		Width:  0,
		Height: 0,
		Board:  make(map[int]map[int]int),
		Energy: make(map[int]map[int]int),
	}
	if err := files.ReadLines(file, func(line string) {
		contraption.Board[contraption.Height] = make(map[int]int)
		for i, c := range line {
			switch c {
			case '/':
				contraption.Board[contraption.Height][i] = Mirror1
			case '\\':
				contraption.Board[contraption.Height][i] = Mirror2
			case '|':
				contraption.Board[contraption.Height][i] = SplitterV
			case '-':
				contraption.Board[contraption.Height][i] = SplitterH
			}
		}
		contraption.Width = len(line)
		contraption.Height++
	}); err != nil {
		panic(err)
	}
	best := 0
	for x := 0; x < contraption.Width; x++ {
		contraption.Energize(x, 0, Down, true)
		v := contraption.Measure()
		if v > best {
			best = v
		}
		contraption.Energize(x, contraption.Height-1, Up, true)
		v = contraption.Measure()
		if v > best {
			best = v
		}
	}
	for y := 0; y < contraption.Height; y++ {
		contraption.Energize(0, y, Right, true)
		v := contraption.Measure()
		if v > best {
			best = v
		}
		contraption.Energize(contraption.Width-1, y, Left, true)
		v = contraption.Measure()
		if v > best {
			best = v
		}
	}
	return best
}

func main() {
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 2: ", solve2(Data+"/example.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
