package main

import (
	"aoc2023/pkg/files"
	"strings"
)

const Data = "data/day08"

type MapEntry struct {
	left, right string
}

func solve1(file string) int {
	decisions := ""
	first := true
	maps := make(map[string]MapEntry)
	if err := files.ReadLines(file, func(line string) {
		if first {
			first = false
			decisions = line
			return
		}
		if line == "" {
			return
		}
		tmp1 := strings.SplitN(line, " = ", 2)
		tmp2 := strings.Split(tmp1[1][1:len(tmp1[1])-1], ", ")
		maps[tmp1[0]] = MapEntry{tmp2[0], tmp2[1]}
	}); err != nil {
		panic(err)
	}

	location := "AAA"
	steps := 0
	decision_count := len(decisions)
	for location != "ZZZ" {
		decision := decisions[steps%decision_count]
		steps++
		if decision == 'L' {
			location = maps[location].left
		} else {
			location = maps[location].right
		}
	}

	return steps
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(integers ...int) int {
	l := len(integers)
	if l == 0 {
		return 0
	} else if l == 1 {
		return integers[0]
	}
	res := integers[0] * integers[1] / GCD(integers[0], integers[1])
	for i := 2; i < len(integers); i++ {
		res = LCM(res, integers[i])
	}
	return res
}

func solve2(file string) int {
	decisions := ""
	first := true
	maps := make(map[string]MapEntry)
	if err := files.ReadLines(file, func(line string) {
		if first {
			first = false
			decisions = line
			return
		}
		if line == "" {
			return
		}
		tmp1 := strings.SplitN(line, " = ", 2)
		tmp2 := strings.Split(tmp1[1][1:len(tmp1[1])-1], ", ")
		maps[tmp1[0]] = MapEntry{tmp2[0], tmp2[1]}
	}); err != nil {
		panic(err)
	}

	locations := make([]string, 0)
	for k := range maps {
		if k[len(k)-1] == 'A' {
			locations = append(locations, k)
		}
	}

	// Since all paths are the same length, we can determine the number of steps
	// to reach a valid goal and then determine the least common multiple to get
	// all of them into one at the same time.
	periods := make([]int, len(locations))
	decision_count := len(decisions)
	for i, location := range locations {
		periods[i] = 0
		loc := location
		for {
			if decisions[periods[i]%decision_count] == 'L' {
				loc = maps[loc].left
			} else {
				loc = maps[loc].right
			}
			periods[i]++
			if loc[len(loc)-1] == 'Z' {
				break
			}
		}
	}

	return LCM(periods...)
}

func main() {
	print("Example 1a: ", solve1(Data+"/example1.txt"), "\n")
	print("Example 1b: ", solve1(Data+"/example2.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 2: ", solve2(Data+"/example3.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
