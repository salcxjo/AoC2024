package main

import (
	"aoc2023/pkg/files"
	"strconv"
	"strings"
)

const Data = "data/day19"

const (
	UNKNOWN      = 0
	COMPARE_LESS = 1
	COMPARE_MORE = 2
	ASSIGN       = 3
)

type WorkflowEntry struct {
	op     int
	target string
	index  byte
	value  int
}

type Workflow []WorkflowEntry
type Workflows map[string]Workflow

func solve1(file string) int {
	res := 0

	workflows := make(Workflows, 0)
	read_workflows := true
	if err := files.ReadLines(file, func(line string) {
		if line == "" {
			read_workflows = false
			return
		}
		if read_workflows {
			tmp := strings.Split(line, "{")
			wf := Workflow{}
			for _, entry := range strings.Split(tmp[1][:len(tmp[1])-1], ",") {
				tmp1 := strings.SplitN(entry, ":", 2)
				if len(tmp1) == 1 {
					wf = append(wf, WorkflowEntry{op: ASSIGN, target: tmp1[0]})
				} else if tmp2 := strings.SplitN(tmp1[0], "<", 2); len(tmp2) > 1 {
					v, _ := strconv.Atoi(tmp2[1])
					wf = append(wf, WorkflowEntry{op: COMPARE_LESS, index: tmp2[0][0], target: tmp1[1], value: v})
				} else if tmp2 := strings.SplitN(tmp1[0], ">", 2); len(tmp2) > 1 {
					v, _ := strconv.Atoi(tmp2[1])
					wf = append(wf, WorkflowEntry{op: COMPARE_MORE, index: tmp2[0][0], target: tmp1[1], value: v})
				} else {
					panic("Unknown operation")
				}
			}
			workflows[tmp[0]] = wf
		} else {
			part := make(map[rune]int)
			for _, e := range strings.SplitN(line[1:len(line)-1], ",", 4) {
				tmp := strings.SplitN(e, "=", 2)
				v, _ := strconv.Atoi(tmp[1])
				part[rune(tmp[0][0])] = v
			}
			current := "in"
			for {
				wf := workflows[current]
				new := ""
				for _, e := range wf {
					switch e.op {
					case ASSIGN:
						new = e.target
					case COMPARE_LESS:
						if part[rune(e.index)] < e.value {
							new = e.target
						}
					case COMPARE_MORE:
						if part[rune(e.index)] > e.value {
							new = e.target
						}
					}
					if new != "" {
						current = new
						break
					}
				}
				if current == "R" {
					break
				} else if current == "A" {
					res += part['x'] + part['m'] + part['a'] + part['s']
					break
				}
			}
		}
	}); err != nil {
		panic(err)
	}
	return res
}

func solve2(file string) int {
	res := 0

	workflows := make(Workflows, 0)
	read_workflows := true
	if err := files.ReadLines(file, func(line string) {
		if line == "" {
			read_workflows = false
			return
		}
		if read_workflows {
			tmp := strings.Split(line, "{")
			wf := Workflow{}
			for _, entry := range strings.Split(tmp[1][:len(tmp[1])-1], ",") {
				tmp1 := strings.SplitN(entry, ":", 2)
				if len(tmp1) == 1 {
					wf = append(wf, WorkflowEntry{op: ASSIGN, target: tmp1[0]})
				} else if tmp2 := strings.SplitN(tmp1[0], "<", 2); len(tmp2) > 1 {
					v, _ := strconv.Atoi(tmp2[1])
					wf = append(wf, WorkflowEntry{op: COMPARE_LESS, index: tmp2[0][0], target: tmp1[1], value: v})
				} else if tmp2 := strings.SplitN(tmp1[0], ">", 2); len(tmp2) > 1 {
					v, _ := strconv.Atoi(tmp2[1])
					wf = append(wf, WorkflowEntry{op: COMPARE_MORE, index: tmp2[0][0], target: tmp1[1], value: v})
				} else {
					panic("Unknown operation")
				}
			}
			workflows[tmp[0]] = wf
		} else {
			return
		}
	}); err != nil {
		panic(err)
	}

	for _, wf := range workflows {
		for _, e := range wf {
			_ = e
		}
	}
	return res
}

func main() {
	print("Example 1: ", solve1(Data+"/example.txt"), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt"), "\n")

	print("Example 2: ", solve2(Data+"/example.txt"), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
