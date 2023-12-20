package main

import (
	"aoc2023/pkg/files"
	"strings"
)

const Data = "data/day20"

const (
	LOW  = false
	HIGH = true

	BROADCASTER = 0
	FLIP_FLOP   = 1
	CONJUNCTION = 2
)

type Module struct {
	state   map[string]bool
	targets []string
	kind    int
}

type Pulse struct {
	recipient string
	sender    string
	value     bool
}

type Pulses []Pulse

func solve1(file string, presses int) int {
	pulses := make(Pulses, 0)
	modules := make(map[string]Module)

	if err := files.ReadLines(file, func(line string) {
		tmp1 := strings.SplitN(line, " -> ", 2)
		targets := strings.Split(tmp1[1], ", ")
		switch tmp1[0][0] {
		case '%':
			name := tmp1[0][1:]
			modules[name] = Module{
				state:   map[string]bool{},
				targets: targets,
				kind:    FLIP_FLOP,
			}
		case '&':
			name := tmp1[0][1:]
			modules[name] = Module{
				state:   map[string]bool{},
				targets: targets,
				kind:    CONJUNCTION,
			}
		default:
			name := tmp1[0]
			modules[name] = Module{
				state:   map[string]bool{},
				targets: targets,
				kind:    BROADCASTER,
			}
		}
	}); err != nil {
		panic(err)
	}

	// "wire up" current states for conjunctions
	for name, module := range modules {
		for _, target := range module.targets {
			if _, ok := modules[target]; ok {
				modules[target].state[name] = LOW
			}
		}
	}

	lows, highs := 0, 0
	for i := 0; i < presses; i++ {
		pulses = append(pulses, Pulse{
			recipient: "broadcaster",
			sender:    "button",
			value:     LOW,
		})
		for len(pulses) > 0 {
			pulse := pulses[0]
			pulses = pulses[1:]
			module := modules[pulse.recipient]

			if pulse.value == LOW {
				lows++
			} else {
				highs++
			}

			if presses == 1 {
				if pulse.value == HIGH {
					print(pulse.sender, " -high-> ", pulse.recipient, "\n")
				} else {
					print(pulse.sender, " -low-> ", pulse.recipient, "\n")
				}
			}

			switch module.kind {
			case BROADCASTER:
				for _, target := range module.targets {
					pulses = append(pulses, Pulse{
						recipient: target,
						sender:    pulse.recipient,
						value:     pulse.value,
					})
				}
			case FLIP_FLOP:
				if pulse.value == LOW {
					nv := !module.state["FF"]
					module.state["FF"] = nv
					for _, target := range module.targets {
						pulses = append(pulses, Pulse{
							recipient: target,
							sender:    pulse.recipient,
							value:     nv,
						})
					}
				}
			case CONJUNCTION:
				module.state[pulse.sender] = pulse.value
				out := LOW
				for _, v := range module.state {
					if v == LOW {
						out = HIGH
						break
					}
				}
				for _, target := range module.targets {
					pulses = append(pulses, Pulse{
						recipient: target,
						sender:    pulse.recipient,
						value:     out,
					})
				}
			}
		}
	}

	return lows * highs
}

func solve2(file string) int {
	pulses := make(Pulses, 0)
	modules := make(map[string]Module)

	if err := files.ReadLines(file, func(line string) {
		tmp1 := strings.SplitN(line, " -> ", 2)
		targets := strings.Split(tmp1[1], ", ")
		switch tmp1[0][0] {
		case '%':
			name := tmp1[0][1:]
			modules[name] = Module{
				state:   map[string]bool{},
				targets: targets,
				kind:    FLIP_FLOP,
			}
		case '&':
			name := tmp1[0][1:]
			modules[name] = Module{
				state:   map[string]bool{},
				targets: targets,
				kind:    CONJUNCTION,
			}
		default:
			name := tmp1[0]
			modules[name] = Module{
				state:   map[string]bool{},
				targets: targets,
				kind:    BROADCASTER,
			}
		}
	}); err != nil {
		panic(err)
	}

	// "wire up" current states for conjunctions
	for name, module := range modules {
		for _, target := range module.targets {
			if _, ok := modules[target]; ok {
				modules[target].state[name] = LOW
			}
		}
	}

	presses := 0
	for {
		pulses = append(pulses, Pulse{
			recipient: "broadcaster",
			sender:    "button",
			value:     LOW,
		})
		presses++
		if presses%1000000 == 0 {
			print("\rSo far:", presses)
		}
		for len(pulses) > 0 {
			pulse := pulses[0]
			pulses = pulses[1:]
			module := modules[pulse.recipient]

			if pulse.recipient == "rx" && pulse.value == LOW {
				return presses
			}

			switch module.kind {
			case BROADCASTER:
				for _, target := range module.targets {
					pulses = append(pulses, Pulse{
						recipient: target,
						sender:    pulse.recipient,
						value:     pulse.value,
					})
				}
			case FLIP_FLOP:
				if pulse.value == LOW {
					nv := !module.state["FF"]
					module.state["FF"] = nv
					for _, target := range module.targets {
						pulses = append(pulses, Pulse{
							recipient: target,
							sender:    pulse.recipient,
							value:     nv,
						})
					}
				}
			case CONJUNCTION:
				module.state[pulse.sender] = pulse.value
				out := LOW
				for _, v := range module.state {
					if v == LOW {
						out = HIGH
						break
					}
				}
				for _, target := range module.targets {
					pulses = append(pulses, Pulse{
						recipient: target,
						sender:    pulse.recipient,
						value:     out,
					})
				}
			}
		}
	}
}

func main() {
	print("Example 1:\n", solve1(Data+"/example1.txt", 1), "\n")
	print("Example 2:\n", solve1(Data+"/example2.txt", 1), "\n")
	print("Example 1: ", solve1(Data+"/example1.txt", 1000), "\n")
	print("Example 2: ", solve1(Data+"/example2.txt", 1000), "\n")
	print("Solution 1: ", solve1(Data+"/input.txt", 1000), "\n")
	print("Solution 2: ", solve2(Data+"/input.txt"), "\n")

	// print("Example 2: ", solve2(Data+"/example.txt"), "\n")
	// print("Solution 2: ", solve2(Data+"/input.txt"), "\n")
}
