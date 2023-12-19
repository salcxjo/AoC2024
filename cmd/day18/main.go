package main

import (
	"aoc2023/pkg/files"
	"strconv"
	"strings"
)

const Data = "data/day18"

func fill(lagoon *map[int]map[int]string, x int, y int, min_x int, max_x int, c string) {
	if _, ok := (*lagoon)[y]; !ok {
		return
	}
	if x < min_x || x > max_x {
		return
	}
	if (*lagoon)[y][x] == "" {
		(*lagoon)[y][x] = c
		fill(lagoon, x-1, y, min_x, max_x, c)
		fill(lagoon, x+1, y, min_x, max_x, c)
		fill(lagoon, x, y-1, min_x, max_x, c)
		fill(lagoon, x, y+1, min_x, max_x, c)
	}
}

func solve1(file string) int {
	lagoon := make(map[int]map[int]string)
	lagoon[0] = make(map[int]string)
	lagoon[0][0] = ""
	x, y := 0, 0
	min_x, min_y := 0, 0
	max_x, max_y := 0, 0
	if err := files.ReadLines(file, func(line string) {
		tmp := strings.Split(line, " ")
		d := line[0]
		v, _ := strconv.Atoi(tmp[1])
		c := tmp[2][1 : len(tmp[2])-1]
		switch d {
		case 'R':
			for ; v > 0; v-- {
				x++
				lagoon[y][x] = c
			}
		case 'L':
			for ; v > 0; v-- {
				x--
				lagoon[y][x] = c
			}
		case 'U':
			for ; v > 0; v-- {
				y--
				if _, ok := lagoon[y]; !ok {
					lagoon[y] = make(map[int]string)
				}
				lagoon[y][x] = c
			}
		case 'D':
			for ; v > 0; v-- {
				y++
				if _, ok := lagoon[y]; !ok {
					lagoon[y] = make(map[int]string)
				}
				lagoon[y][x] = c
			}
		}
		if x < min_x {
			min_x = x
		}
		if x > max_x {
			max_x = x
		}
		if y < min_y {
			min_y = y
		}
		if y > max_y {
			max_y = y
		}
	}); err != nil {
		panic(err)
	}

	for y := min_y; y <= max_y; y++ {
		fill(&lagoon, min_x, y, min_x, max_x, "O")
		fill(&lagoon, max_x, y, min_x, max_x, "O")
	}
	for x := min_x; x <= max_x; x++ {
		fill(&lagoon, x, min_y, min_x, max_x, "O")
		fill(&lagoon, x, max_y, min_x, max_x, "O")
	}

	res := 0
	for y := min_y; y <= max_y; y++ {
		for x := min_x; x <= max_x; x++ {
			if lagoon[y][x] == "O" {
				res++
			}
		}
	}
	return (max_x-min_x+1)*(max_y-min_y+1) - res
}

func solve2(file string) int {
	return 0
	/*
		lagoon := make(map[int]map[int]string)
		lagoon[0] = make(map[int]string)
		lagoon[0][0] = ""
		x, y := 0, 0
		min_x, min_y := 0, 0
		max_x, max_y := 0, 0
		if err := files.ReadLines(file, func(line string) {
			tmp := strings.Split(line, " ")

			d := tmp[2][7]
			_v := new(big.Int)
			_v.SetString(tmp[2][2:7], 16)
			v := int(_v.Int64())
			c := "set"
			switch d {
			case '0':
				for ; v > 0; v-- {
					x++
					lagoon[y][x] = c
				}
			case '2':
				for ; v > 0; v-- {
					x--
					lagoon[y][x] = c
				}
			case '3':
				for ; v > 0; v-- {
					y--
					if _, ok := lagoon[y]; !ok {
						lagoon[y] = make(map[int]string)
					}
					lagoon[y][x] = c
				}
			case '1':
				for ; v > 0; v-- {
					y++
					if _, ok := lagoon[y]; !ok {
						lagoon[y] = make(map[int]string)
					}
					lagoon[y][x] = c
				}
			}
			if x < min_x {
				min_x = x
			}
			if x > max_x {
				max_x = x
			}
			if y < min_y {
				min_y = y
			}
			if y > max_y {
				max_y = y
			}
		}); err != nil {
			panic(err)
		}

		for y := min_y; y <= max_y; y++ {
			fill(&lagoon, min_x, y, min_x, max_x, "O")
			fill(&lagoon, max_x, y, min_x, max_x, "O")
		}
		for x := min_x; x <= max_x; x++ {
			fill(&lagoon, x, min_y, min_x, max_x, "O")
			fill(&lagoon, x, max_y, min_x, max_x, "O")
		}

		res := 0
		for y := min_y; y <= max_y; y++ {
			for x := min_x; x <= max_x; x++ {
				if lagoon[y][x] == "O" {
					res++
				}
			}
		}
		return (max_x-min_x+1)*(max_y-min_y+1) - res
	*/
}

func main() {
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 2: ", solve2(Data+"/example.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
