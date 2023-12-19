package city

import (
	"aoc2023/pkg/files"
)

const (
	None  = 0
	Up    = -1
	Left  = -2
	Down  = +1
	Right = +2
)

type City struct {
	Tiles  map[int]map[int]int
	Width  int
	Height int
}

type cityNode struct {
	X         int
	Y         int
	Direction int
	Straight  int
	Costs     int
	Previous  *cityNode
}

func ReadFile(file string) *City {
	res := &City{Tiles: make(map[int]map[int]int)}
	if err := files.ReadLines(file, func(line string) {
		res.AddLine(line)
	}); err != nil {
		panic(err)
	}
	return res
}

func (c *City) AddLine(line string) {
	c.Tiles[c.Height] = make(map[int]int)
	for i, t := range line {
		c.Tiles[c.Height][i] = int(t - '0')
	}
	c.Width = len(line)
	c.Height++
}

func (c *City) GetHeatLoss() int {
	res := -1
	// Initialize the cost matrix
	costs := make([][]int, c.Height)
	for i := 0; i < c.Height; i++ {
		costs[i] = make([]int, c.Width)
		for j := 0; j < c.Width; j++ {
			costs[i][j] = -1
		}
	}

	// Initialize the queue with the starting node
	queue := []cityNode{{X: 0, Y: 0, Direction: None, Straight: 0, Costs: 0, Previous: nil}}
	costs[0][0] = 0

	// Perform breadth-first search
	for len(queue) > 0 {
		// Dequeue the current node
		node := queue[0]
		queue = queue[1:]

		for _, dir := range []int{Right, Down, Left, Up} {
			if dir == -node.Direction {
				continue
			}

			// Calculate the new position
			newX := node.X
			newY := node.Y
			if dir == Up {
				newY--
			} else if dir == Left {
				newX--
			} else if dir == Down {
				newY++
			} else if dir == Right {
				newX++
			}

			// Check if the new position is valid
			if newX >= 0 && newX < c.Width && newY >= 0 && newY < c.Height {
				// Calculate the cost of the new position
				newCost := node.Costs + c.Tiles[newY][newX]

				if costs[newY][newX] == -1 || newCost <= costs[newY][newX] {
					costs[newY][newX] = newCost
					queue = append(queue, cityNode{X: newX, Y: newY, Direction: dir, Costs: newCost, Previous: &node})
				}
			}
		}

		continue

		// Check if reached the goal
		if node.X == c.Width-1 && node.Y == c.Height-1 {
			// return node.Costs
			/*
				println()
				for y := 0; y < c.Height; y++ {
					for x := 0; x < c.Width; x++ {
						found := false
						for n := &node; n != nil; n = n.Previous {
							if n.X == x && n.Y == y {
								found = true
								switch n.Direction {
								case Up:
									print("^")
								case Left:
									print("<")
								case Down:
									print("v")
								case Right:
									print(">")
								case None:
									print("S")
								default:
									print("?")
								}
								break
							}
						}
						if !found {
							print(".")
						}
					}
					println()
				}
			*/
			if res == -1 || node.Costs < res {
				res = node.Costs
			}
			continue
		}

		// Explore the neighbors
		// for _, dir := range []int{Up, Left, Down, Right} {
		for _, dir := range []int{Right, Down, Left, Up} {
			// Skip the opposite direction
			if dir == -node.Direction {
				continue
			}

			// Calculate the new position
			newX := node.X
			newY := node.Y
			if dir == Up {
				newY--
			} else if dir == Left {
				newX--
			} else if dir == Down {
				newY++
			} else if dir == Right {
				newX++
			}

			// Check if the new position is valid
			if newX >= 0 && newX < c.Width && newY >= 0 && newY < c.Height {
				// Calculate the cost of the new position
				newCost := node.Costs + c.Tiles[newY][newX]

				found := false
				for prev := node.Previous; prev != nil; prev = prev.Previous {
					if prev.X == newX && prev.Y == newY {
						found = true
						break
					}
				}
				if found {
					continue
				}

				// Check if the cost is within the limit
				//if costs[newY][newX] == -1 || newCost < costs[newY][newX] {
				nn := cityNode{X: newX, Y: newY, Direction: dir, Straight: 1, Costs: newCost, Previous: &node}
				if dir == node.Direction {
					if node.Straight == 3 {
						continue
					}
					nn.Straight += node.Straight
				}

				// Remove queued nodes with higher cost but same target
				for i := 0; i < len(queue); i++ {
					if queue[i].X == newX && queue[i].Y == newY && queue[i].Costs > newCost {
						queue = append(queue[:i], queue[i+1:]...)
						i--
					}
				}

				// Update the cost matrix and enqueue the new node
				costs[newY][newX] = newCost
				queue = append(queue, nn)
				//}
			}
		}
	}

	return res //costs[c.Height-1][c.Width-1]
}
