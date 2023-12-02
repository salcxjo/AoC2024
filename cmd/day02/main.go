package main

import (
	"aoc2023/pkg/files"
	"strconv"
	"strings"
)

const Data = "data/day02"

func solve1(file string) int {
	res := 0
	if err := files.ReadLines(file, func(line string) {
		tmp1 := strings.SplitN(line, ": ", 2)
		tmp2 := strings.SplitN(tmp1[0], " ", 2)
		game_id, _ := strconv.Atoi(tmp2[1])

		limits := make(map[string]int)
		limits["red"] = 12
		limits["green"] = 13
		limits["blue"] = 14

		sets := strings.Split(tmp1[1], "; ")
		possible := true
		for _, set := range sets {
			cubes := map[string]int{"red": 0, "green": 0, "blue": 0}
			content := strings.Split(set, ", ")
			for _, cube := range content {
				tmp1 = strings.SplitN(cube, " ", 2)
				count, _ := strconv.Atoi(tmp1[0])
				cubes[tmp1[1]] += count
			}
			for k, v := range cubes {
				if v > limits[k] {
					possible = false
					break
				}
			}
			if !possible {
				break
			}
		}
		if possible {
			res += game_id
		}
	}); err != nil {
		panic(err)
	}
	return res
}

func solve2(file string) int {
	res := 0
	if err := files.ReadLines(file, func(line string) {
		tmp1 := strings.SplitN(line, ": ", 2)

		cubes := map[string]int{"red": 0, "green": 0, "blue": 0}

		sets := strings.Split(tmp1[1], "; ")
		for _, set := range sets {
			content := strings.Split(set, ", ")
			for _, cube := range content {
				tmp1 = strings.SplitN(cube, " ", 2)
				count, _ := strconv.Atoi(tmp1[0])
				id := tmp1[1]
				if cubes[id] < count {
					cubes[id] = count
				}
			}
		}
		power := 1
		for _, v := range cubes {
			power *= v
		}
		res += power
	}); err != nil {
		panic(err)
	}
	return res
}

func main() {
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 1: ", solve2(Data+"/example.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
