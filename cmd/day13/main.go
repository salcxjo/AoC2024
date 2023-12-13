package main

import (
	"aoc2023/pkg/files"
)

const Data = "data/day13"

type Pattern struct {
	data   map[int]map[int]bool
	width  int
	height int
}

func (p *Pattern) get(x, y int) bool {
	if p.data == nil {
		return false
	}
	if _, ok := p.data[y]; !ok {
		return false
	}
	return p.data[y][x]
}

func (p *Pattern) set(x, y int, v bool) {
	if p.data == nil {
		p.data = make(map[int]map[int]bool)
	}
	if _, ok := p.data[y]; !ok {
		p.data[y] = make(map[int]bool)
	}
	p.data[y][x] = v
	if x+1 > p.width {
		p.width = x + 1
	}
	if y+1 > p.height {
		p.height = y + 1
	}
}

func (p *Pattern) findMirrors() int {
	res := 0
	// Find vertical mirror
	for x := 0; x < p.width-1; x++ {
		mirrors_here := true
		for d := 0; x+d+1 < p.width && x-d >= 0 && mirrors_here; d++ {
			for y := 0; y < p.height; y++ {
				if p.get(x-d, y) != p.get(x+d+1, y) {
					mirrors_here = false
					break
				}
			}
		}
		if mirrors_here {
			res += x + 1
			break
		}
	}
	// Find horizontal mirror
	for y := 0; y < p.height-1; y++ {
		mirrors_here := true
		for d := 0; y+d+1 < p.height && y-d >= 0 && mirrors_here; d++ {
			for x := 0; x < p.width; x++ {
				if p.get(x, y-d) != p.get(x, y+d+1) {
					mirrors_here = false
					break
				}
			}
		}
		if mirrors_here {
			res += (y + 1) * 100
			break
		}
	}
	return res
}

func (p *Pattern) findMirrorsSmudged() int {
	res := 0
	// Find vertical mirror
	for x := 0; x < p.width-1; x++ {
		mirrors_here := true
		smudge := false
		for d := 0; x+d+1 < p.width && x-d >= 0 && mirrors_here; d++ {
			for y := 0; y < p.height; y++ {
				if p.get(x-d, y) != p.get(x+d+1, y) {
					if smudge {
						mirrors_here = false
						break
					}
					smudge = true
				}
			}
		}
		if mirrors_here && smudge {
			res += x + 1
			break
		}
	}
	// Find horizontal mirror
	for y := 0; y < p.height-1; y++ {
		mirrors_here := true
		smudge := false
		for d := 0; y+d+1 < p.height && y-d >= 0 && mirrors_here; d++ {
			for x := 0; x < p.width; x++ {
				if p.get(x, y-d) != p.get(x, y+d+1) {
					if smudge {
						mirrors_here = false
						break
					}
					smudge = true
				}
			}
		}
		if mirrors_here && smudge {
			res += (y + 1) * 100
			break
		}
	}
	return res
}

type Patterns []Pattern

func solve1(file string) int {
	res := 0
	patterns := make(Patterns, 0)
	var pattern Pattern
	y := 0

	if err := files.ReadLines(file, func(line string) {
		if line == "" {
			y = 0
			patterns = append(patterns, pattern)
			pattern = Pattern{}
			return
		}
		for x, c := range line {
			if c == '#' {
				pattern.set(x, y, true)
			}
		}
		y++
	}); err != nil {
		panic(err)
	}
	patterns = append(patterns, pattern)
	for _, pattern := range patterns {
		res += pattern.findMirrors()
	}
	return res
}

func solve2(file string) int {
	res := 0
	patterns := make(Patterns, 0)
	var pattern Pattern
	y := 0

	if err := files.ReadLines(file, func(line string) {
		if line == "" {
			y = 0
			patterns = append(patterns, pattern)
			pattern = Pattern{}
			return
		}
		for x, c := range line {
			if c == '#' {
				pattern.set(x, y, true)
			}
		}
		y++
	}); err != nil {
		panic(err)
	}
	patterns = append(patterns, pattern)
	for _, pattern := range patterns {
		res += pattern.findMirrorsSmudged()
	}
	return res
}

func main() {
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 2: ", solve2(Data+"/example.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
