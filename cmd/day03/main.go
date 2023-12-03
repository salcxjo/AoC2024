package main

import (
	"aoc2023/pkg/files"
	"regexp"
	"strconv"
)

const Data = "data/day03"

var rx_numbers = regexp.MustCompile(`\d+`)
var rx_icons = regexp.MustCompile(`[^\.\d]+`)
var rx_stars = regexp.MustCompile(`\*`)

func solve1(file string) int {
	res := 0
	lines := []string{}
	if err := files.ReadLines(file, func(line string) {
		lines = append(lines, line)
	}); err != nil {
		panic(err)
	}

	line_count := len(lines)

	for i, line := range lines {
		matches := rx_numbers.FindAllStringIndex(line, -1)
		for _, match := range matches {
			valid := false
			for j := max(0, i-1); j <= min(line_count-1, i+1); j++ {
				icon_matches := rx_icons.FindAllStringIndex(lines[j], -1)
				for _, icon_match := range icon_matches {
					if match[0] <= icon_match[1] && match[1] >= icon_match[0] {
						valid = true
						break
					}
				}
				if valid {
					v, _ := strconv.Atoi(line[match[0]:match[1]])
					res += v
					break
				}
			}
		}
	}

	return res
}

func solve2(file string) int {
	res := 0
	lines := []string{}
	if err := files.ReadLines(file, func(line string) {
		lines = append(lines, line)
	}); err != nil {
		panic(err)
	}

	line_count := len(lines)

	for i, line := range lines {
		matches := rx_stars.FindAllStringIndex(line, -1)
		for _, match := range matches {
			count := 0
			ratio := 0
			for j := max(0, i-1); j <= min(line_count-1, i+1); j++ {
				number_matches := rx_numbers.FindAllStringIndex(lines[j], -1)
				for _, number_match := range number_matches {
					if match[0] <= number_match[1] && match[1] >= number_match[0] {
						if count == 0 {
							ratio, _ = strconv.Atoi(lines[j][number_match[0]:number_match[1]])
						} else if count == 1 {
							v, _ := strconv.Atoi(lines[j][number_match[0]:number_match[1]])
							ratio *= v
						}
						count++
					}
				}
			}
			if count == 2 {
				res += ratio
			}
		}
	}

	return res
}

func main() {
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 1: ", solve2(Data+"/example.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
