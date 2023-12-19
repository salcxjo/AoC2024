package main

import (
	"aoc2023/pkg/city"
	"aoc2023/pkg/files"
)

const Data = "data/day17"

func solve1(file string) int {
	city := city.ReadFile(file)
	return city.GetHeatLoss()
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
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	// print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	// print("Example 2: ", solve2(Data+"/example.txt"), "\n")
	// print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
