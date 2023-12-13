package main

import (
	"aoc2023/pkg/files"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const Printing = false

const Data = "data/day12"

const (
	Undefined   = ' '
	Operational = '.'
	Damaged     = '#'
	Unknown     = '?'
)

func countVariations(input string, pattern []int, need_operational bool, debug string) int {
	_debug := debug
	need_operational = false
	input_length := len(input)
	pattern_length := len(pattern)
	if input_length == 0 {
		if pattern_length == 0 || (pattern_length == 1 && pattern[0] == 0) {
			if Printing {
				print(_debug+input, " ok 1!\n")
			}
			return 1
		}
		if Printing {
			print(_debug+input, " fail 1!\n")
		}
		return 0
	}
	if pattern_length == 0 {
		if input_length > 0 && strings.IndexRune(input, Damaged) == -1 {
			if Printing {
				print(_debug+input, " ok 2!\n")
			}
			return 1
		}
		return 0
	}
	if pattern[0] == 0 {
		if input[0] == Damaged {
			if Printing {
				print(_debug+input, " fail 2!\n")
			}
			return 0
		}
		return countVariations(input[1:], pattern[1:], false, debug+".")
	}

	length_required := pattern_length - 1
	for _, v := range pattern {
		length_required += v
	}
	if length_required > input_length {
		if Printing {
			print(_debug+input, " fail length!\n")
		}
		return 0
	}

	for i, c := range input {
		if c == Operational {
			need_operational = false
			debug += string(Operational)
			continue
		}
		if c != Operational {
			if c == Damaged && need_operational {
				if Printing {
					print(_debug+input, " fail 3!\n")
				}
				return 0
			}
			res := 0

			// Assume this is operational
			if c == Unknown {
				res += countVariations(input[i+1:], pattern, false, debug+string(Operational))
			}

			count := 1
			for i = i + 1; i < input_length && (input[i] != Operational); i++ {
				if /*input[j] == Unknown &&*/ count == pattern[0] {
					_pattern := append([]int{0}, pattern[1:]...)
					res += countVariations(input[i:], _pattern, false, debug+(strings.Repeat(string(Damaged), count)))
					return res
				}
				count++
			}
			if count == pattern[0] {
				_pattern := append([]int{0}, pattern[1:]...)
				res += countVariations(input[i:], _pattern, false, debug+(strings.Repeat(string(Damaged), count)))
			}
			return res
		}
	}
	if Printing {
		print(_debug+input, " fail(!!)!\n")
	}
	return 0
}

func solve1(file string) int {
	res := 0
	if err := files.ReadLines(file, func(line string) {
		tmp1 := strings.SplitN(line, " ", 2)
		counts := make([]int, 0)
		for _, c := range strings.Split(tmp1[1], ",") {
			v, _ := strconv.Atoi(c)
			counts = append(counts, v)
		}
		if Printing {
			println()
			println(tmp1[0], " (", tmp1[1], ")")
		}
		cur := countVariations(tmp1[0], counts, false, "")
		if Printing {
			println(tmp1[0], " = ", cur)
		}
		res += cur
	}); err != nil {
		panic(err)
	}
	return res
}

func solve2(file string) int {
	res := 0
	lines := make([]string, 0)
	if err := files.ReadLines(file, func(line string) {
		lines = append(lines, line)
	}); err != nil {
		panic(err)
	}
	count := len(lines)
	done := 0
	print("\n")

	cache := make(map[string]int)
	files.ReadLines(file+".cache", func(line string) {
		if line[0] == ';' {
			return
		}
		tmp := strings.SplitN(line, " = ", 2)
		cache[tmp[0]], _ = strconv.Atoi(tmp[1])
	})

	wg := sync.WaitGroup{}

	strm, err := os.OpenFile(file+".cache", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		strm, _ = os.OpenFile(file+".cache", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}
	cache_writer := io.Writer(strm)

	// ch := make(chan struct{}, 8)
	for j, line := range lines {
		if v, ok := cache[line]; ok {
			res += v
			done++
			print("\rC: ", done, "/", count, " (line ", j, " = ", res, ")")
			continue
		}

		// ch <- struct{}{}
		wg.Add(1)
		go func(j int, line string) {
			// defer func() { <-ch }()
			start := time.Now()
			key := line
			tmp1 := strings.SplitN(line, " ", 2)
			tcounts := make([]int, 0)
			for _, c := range strings.Split(tmp1[1], ",") {
				v, _ := strconv.Atoi(c)
				tcounts = append(tcounts, v)
			}
			line = tmp1[0]
			for i := 0; i < 4; i++ {
				line += "?" + tmp1[0]
			}
			counts := make([]int, 0)
			for i := 0; i < 5; i++ {
				counts = append(counts, tcounts...)
			}
			v := countVariations(line, counts, false, "")
			done++
			print("\rN: ", done, "/", count, " (line ", j, " = ", res, ")")
			cache_writer.Write([]byte("; " + time.Since(start).String() + "\n" + key + " = " + strconv.Itoa(v) + "\n"))
			res += v
			wg.Done()
		}(j, line)
	}
	wg.Wait()
	print("\n")
	strm.Close()
	return res
}

func main() {
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 2: ", solve2(Data+"/example.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
