package main

import (
	"aoc2023/pkg/files"
	"strconv"
	"strings"
)

const Data = "data/day15"

func hash(input string) int {
	res := 0
	for _, c := range input {
		res += int(c)
		res *= 17
		res %= 256
	}
	return res
}

func solve1(file string) int {
	res := 0
	if err := files.ReadLines(file, func(line string) {
		tmp := strings.Split(line, ",")
		for _, t := range tmp {
			res += hash(t)
		}
	}); err != nil {
		panic(err)
	}
	return res
}

func solve2(file string) int {
	res := 0
	boxes := make(map[int][]string)
	if err := files.ReadLines(file, func(line string) {
		tmp1 := strings.Split(line, ",")
		for _, t := range tmp1 {
			p := strings.IndexAny(t, "-=")
			op := t[p]
			label := t[:p]
			id := hash(label)
			v := t[p+1:]
			switch op {
			case '-':
				for i, b := range boxes[id] {
					if strings.HasPrefix(b, label+" ") {
						boxes[id] = append(boxes[id][:i], boxes[id][i+1:]...)
						break
					}
				}
			case '=':
				_, ok := boxes[id]
				if !ok {
					boxes[id] = make([]string, 1)
					boxes[id][0] = label + " " + v
				} else {
					found := false
					for i, b := range boxes[id] {
						if strings.HasPrefix(b, label+" ") {
							boxes[id][i] = label + " " + v
							found = true
							break
						}
					}
					if !found {
						boxes[id] = append(boxes[id], label+" "+v)
					}
				}
			}
		}
	}); err != nil {
		panic(err)
	}
	for i, b := range boxes {
		for j, c := range b {
			tmp := strings.Split(c, " ")
			v, _ := strconv.Atoi(tmp[1])
			res += (1 + i) * (j + 1) * v
		}
	}
	return res
}

func main() {
	println(hash("HASH"))
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 2: ", solve2(Data+"/example.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
