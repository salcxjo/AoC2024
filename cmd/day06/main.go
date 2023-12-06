package main

import (
	"aoc2023/pkg/files"
	"slices"
	"strconv"
	"strings"
)

const Data = "data/day06"

func countSolutions(time int, record int) int {
	res := 0
	for a := 1; a < time-1; a++ {
		distance := (time - a) * a
		if distance > record {
			res++
		}
	}
	return res
}

func solve1(file string) int {
	res := 1
	times := make([]int, 0)
	distances := make([]int, 0)
	if err := files.ReadLines(file, func(line string) {
		tmp1 := strings.SplitN(line, ":", 2)
		target := &times
		if tmp1[0] == "Distance" {
			target = &distances
		}
		for _, s := range slices.DeleteFunc(strings.Split(tmp1[1], " "), func(s string) bool { return s == "" }) {
			v, _ := strconv.Atoi(s)
			*target = append(*target, v)
		}
	}); err != nil {
		panic(err)
	}
	for i, t := range times {
		res *= countSolutions(t, distances[i])
	}
	return res
}

func solve2(file string) int {
	res := 0
	time := 0
	distance := 0
	if err := files.ReadLines(file, func(line string) {
		tmp1 := strings.SplitN(line, ":", 2)
		if tmp1[0] == "Time" {
			time, _ = strconv.Atoi(strings.Join(slices.DeleteFunc(strings.Split(tmp1[1], " "), func(s string) bool { return s == "" }), ""))
		} else {
			distance, _ = strconv.Atoi(strings.Join(slices.DeleteFunc(strings.Split(tmp1[1], " "), func(s string) bool { return s == "" }), ""))
		}
	}); err != nil {
		panic(err)
	}
	res = countSolutions(time, distance)
	return res
}

func main() {
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 1: ", solve2(Data+"/example.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
