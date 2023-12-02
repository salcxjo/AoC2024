package main

import (
	"aoc2023/pkg/files"
)

const Data = "data/day##"

func solve1(file string) int {
	res := 0
	if err := files.ReadLines(file, func(line string) {
	}); err != nil {
		panic(err)
	}
	return res
}

func solve2(file string) int {
	res := 0
	if err := files.ReadLines(file, func(line string) {
	}); err != nil {
		panic(err)
	}
	return res
}

func main() {
	print("Example 1: ", solve1(Data+"/example1.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 1: ", solve2(Data+"/example2.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
