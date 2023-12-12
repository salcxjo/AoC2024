package main

import (
	"aoc2023/pkg/files"
	"strconv"
	"strings"
)

const Data = "data/day12"

const (
	Undefined   = ' '
	Operational = '.'
	Damaged     = '#'
	Unknown     = '?'
)

// Naive, but should be reasonable fast for part 1
func testPattern1(l string, pattern []int) int {
	unknowns := 0
	for _, c := range l {
		if c == Unknown {
			unknowns++
		}
	}

	res := 0

	for seed := 0; seed < 1<<unknowns; seed++ {
		cpattern := make([]int, 0)
		p := Undefined
		i := 1
		for _, c := range l {
			if c == Damaged {
				if p == Damaged {
					cpattern[len(cpattern)-1]++
				} else {
					cpattern = append(cpattern, 1)
				}
			} else if c == Unknown {
				if seed&i == i {
					if p == Damaged {
						cpattern[len(cpattern)-1]++
					} else {
						cpattern = append(cpattern, 1)
					}
					c = Damaged
				} else {
					c = Operational
				}
				i <<= 1
			}
			p = c
		}
		if len(cpattern) != len(pattern) {
			continue
		}
		match := true
		for i := 0; i < len(cpattern); i++ {
			if cpattern[i] != pattern[i] {
				match = false
				break
			}
		}
		if match {
			res++
		}
	}
	return res
}

func solve1(file string) int {
	res := 0
	if err := files.ReadLines(file, func(line string) {
		tmp1 := strings.SplitN(line, " ", 2)
		counts := make([]int, 0)
		for _, c := range strings.Split(tmp1[1], ",") {
			v, _ := strconv.Atoi(c)
			counts = append(counts, v)
		}
		res += testPattern1(tmp1[0], counts)
	}); err != nil {
		panic(err)
	}
	return res
}

func solve2(file string) int {
	res := 0
	if err := files.ReadLines(file, func(line string) {
		tmp1 := strings.SplitN(line, " ", 2)
		tcounts := make([]int, 0)
		for _, c := range strings.Split(tmp1[1], ",") {
			v, _ := strconv.Atoi(c)
			tcounts = append(tcounts, v)
		}

	}); err != nil {
		panic(err)
	}
	return res
}

func main() {
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 2: ", solve2(Data+"/example.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
