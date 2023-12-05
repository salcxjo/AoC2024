package main

import (
	"aoc2023/pkg/files"
	"slices"
	"strconv"
	"strings"
)

const Data = "data/day05"

type Mapping struct {
	target int
	length int
}

var mapping map[string]map[int]Mapping

func solve1(file string) int {
	mapping = make(map[string]map[int]Mapping)
	current := ""
	tables := []string{}
	if err := files.ReadLines(file, func(line string) {
		if line == "" {
			return
		}

		tmp1 := strings.SplitN(line, ":", 2)
		first := tmp1[0][0]
		if first < '0' || first > '9' {
			tmp2 := strings.SplitN(tmp1[0], "-", 2)
			current = tmp2[0]
			m, ok := mapping[current]
			if !ok {
				m = make(map[int]Mapping)
				mapping[current] = m
			}

			if current == "seeds" {
				for _, s := range slices.DeleteFunc(strings.Split(tmp1[1], " "), func(s string) bool { return s == "" }) {
					n, _ := strconv.Atoi(s)
					m[n] = Mapping{target: -1, length: 1}
				}
			} else {
				tables = append(tables, current)
			}
		} else {
			m, _ := mapping[current]
			tmp3 := strings.SplitN(line, " ", 3)
			trg, _ := strconv.Atoi(tmp3[0])
			src, _ := strconv.Atoi(tmp3[1])
			lng, _ := strconv.Atoi(tmp3[2])
			m[src] = Mapping{target: trg, length: lng}
		}
	}); err != nil {
		panic(err)
	}

	lowest := -1
	for seed := range mapping["seeds"] {
		current := seed
		for _, table := range tables {
			m, _ := mapping[table]
			for idx, info := range m {
				if idx <= current && current < idx+info.length {
					current += info.target - idx
					break
				}
			}
		}
		if lowest == -1 || current < lowest {
			lowest = current
		}
	}
	return lowest
}

func solve2(file string) int {
	mapping = make(map[string]map[int]Mapping)
	current := ""
	tables := []string{}
	if err := files.ReadLines(file, func(line string) {
		if line == "" {
			return
		}

		tmp1 := strings.SplitN(line, ":", 2)
		first := tmp1[0][0]
		if first < '0' || first > '9' {
			tmp2 := strings.SplitN(tmp1[0], "-", 2)
			current = tmp2[0]
			m, ok := mapping[current]
			if !ok {
				m = make(map[int]Mapping)
				mapping[current] = m
			}

			if current == "seeds" {
				entries := slices.DeleteFunc(strings.Split(tmp1[1], " "), func(s string) bool { return s == "" })
				for i := 0; i < len(entries); i += 2 {
					n, _ := strconv.Atoi(entries[i])
					l, _ := strconv.Atoi(entries[i+1])
					m[n] = Mapping{target: -1, length: l}
				}
			} else {
				tables = append(tables, current)
			}
		} else {
			m, _ := mapping[current]
			tmp3 := strings.SplitN(line, " ", 3)
			trg, _ := strconv.Atoi(tmp3[0])
			src, _ := strconv.Atoi(tmp3[1])
			lng, _ := strconv.Atoi(tmp3[2])
			m[src] = Mapping{target: trg, length: lng}
		}
	}); err != nil {
		panic(err)
	}

	lowest := -1

	// TODO

	return lowest
}

func main() {
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 1: ", solve2(Data+"/example.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
