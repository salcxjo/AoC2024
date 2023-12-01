package main

import (
	"aoc2023/pkg/files"
	"aoc2023/pkg/strings"
)

func solve1(file string) int {
	res := 0
	if err := files.ReadLines(file, func(line string) {
		res += strings.FindFirstDigit(line)*10 + strings.FindLastDigit(line)
	}); err != nil {
		panic(err)
	}
	return res
}

func solve2(file string) int {
	res := 0
	if err := files.ReadLines(file, func(line string) {
		res += strings.FindFirstDigit(line, true)*10 + strings.FindLastDigit(line, true)
	}); err != nil {
		panic(err)
	}
	return res
}

func main() {
	print("Example 1: ", solve1("data/day01/example1.txt"), "\n")
	print("Solution 1: ", solve1("data/day01/input.txt"), "\n")

	print("Example 1: ", solve2("data/day01/example2.txt"), "\n")
	print("Solution 2: ", solve2("data/day01/input.txt"), "\n")
}
