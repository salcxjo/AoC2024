package main

import (
	"aoc2023/pkg/files"
	"strconv"
	"strings"
)

const Data = "data/day09"

func predict1(values []int) int {
	diffs := make([]int, len(values)-1)
	non_zero := false
	for i := 0; i < len(values)-1; i++ {
		d := values[i+1] - values[i]
		diffs[i] = d
		if d != 0 {
			non_zero = true
		}
	}
	if !non_zero {
		return values[len(values)-1]
	}
	return values[len(values)-1] + predict1(diffs)
}

func predict2(values []int) int {
	diffs := make([]int, len(values)-1)
	non_zero := false
	for i := 0; i < len(values)-1; i++ {
		d := values[i+1] - values[i]
		diffs[i] = d
		if d != 0 {
			non_zero = true
		}
	}
	if !non_zero {
		return values[0]
	}
	return values[0] - predict2(diffs)
}

func solve1(file string) int {
	res := 0
	if err := files.ReadLines(file, func(line string) {
		tmp := strings.Split(line, " ")
		values := make([]int, len(tmp))
		for i, v := range tmp {
			values[i], _ = strconv.Atoi(v)
		}
		res += predict1(values)
	}); err != nil {
		panic(err)
	}
	return res
}

func solve2(file string) int {
	res := 0
	if err := files.ReadLines(file, func(line string) {
		tmp := strings.Split(line, " ")
		values := make([]int, len(tmp))
		for i, v := range tmp {
			values[i], _ = strconv.Atoi(v)
		}
		res += predict2(values)
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
