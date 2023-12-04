package main

import (
	"aoc2023/pkg/files"
	"slices"
	"strings"
)

const Data = "data/day04"

func solve1(file string) int {
	total_points := 0
	if err := files.ReadLines(file, func(line string) {
		tmp1 := strings.SplitN(line, ": ", 2)
		tmp2 := strings.SplitN(tmp1[1], " | ", 2)
		winning_cards := slices.DeleteFunc(strings.Split(tmp2[0], " "), func(s string) bool { return s == "" })
		owned_cards := slices.DeleteFunc(strings.Split(tmp2[1], " "), func(s string) bool { return s == "" })
		points := 0
		for _, winning_card := range winning_cards {
			for _, owned_card := range owned_cards {
				if winning_card == owned_card {
					if points == 0 {
						points = 1
					} else {
						points *= 2
					}
					break
				}
			}
		}
		total_points += points
	}); err != nil {
		panic(err)
	}
	return total_points
}

func solve2(file string) int {
	card_count := 0
	// Store the extra copies won of the next 25 cards (maximum owned plus current one)
	extras := make([]int, 26)
	extras_index := 0
	if err := files.ReadLines(file, func(line string) {
		tmp1 := strings.SplitN(line, ": ", 2)
		tmp2 := strings.SplitN(tmp1[1], " | ", 2)
		winning_cards := slices.DeleteFunc(strings.Split(tmp2[0], " "), func(s string) bool { return s == "" })
		owned_cards := slices.DeleteFunc(strings.Split(tmp2[1], " "), func(s string) bool { return s == "" })
		points := 0
		for _, winning_card := range winning_cards {
			for _, owned_card := range owned_cards {
				if winning_card == owned_card {
					points++
					break
				}
			}
		}
		// Get the number of copies of the current card
		number_cards := 1 + extras[extras_index]
		card_count += number_cards
		for i := 1; i <= points; i++ {
			extras[(extras_index+i)%26] += number_cards
		}
		extras[extras_index] = 0
		extras_index = (extras_index + 1) % 26
	}); err != nil {
		panic(err)
	}
	return card_count
}

func main() {
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 1: ", solve2(Data+"/example.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
