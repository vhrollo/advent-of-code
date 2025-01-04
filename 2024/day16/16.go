package main

import (
	"fmt"
	"os"
	"strings"
	"container/heap"
	"math"
)

type Node struct {
	last *Node
	pos [2]int
	dir [2]int
	value int
	heuristic int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value + pq[i].heuristic < pq[j].value + pq[j].heuristic
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return x
}

func abs (a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func heuristic (a, b [2]int) int {
	return abs(a[0] - b[0]) + abs(a[1] - b[1])
}

func getData(path string) string {
    content, err := os.ReadFile(path)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return strings.TrimRight(string(content), "\n")
}

func astar (matrix [][]string, start [2]int, end [2]int) {
	//copy matrix so to print
	matrix_copy := make([][]string, len(matrix))
	for i := range matrix {
		matrix_copy[i] = make([]string, len(matrix[i]))
		copy(matrix_copy[i], matrix[i])
	}

	bestCost := map[[4]int]int{}
	
	directions := [][2]int{
        {-1, 0}, // up
        {0, 1},  // right
        {1, 0},  // down
        {0, -1}, // left
    }

    for r := 0; r < len(matrix); r++ {
        for c := 0; c < len(matrix[r]); c++ {
            for _, d := range directions {
                bestCost[[4]int{r, c, d[0], d[1]}] = math.MaxInt32
            }
        }
    }

	open := PriorityQueue{}
	heap.Init(&open)
	heuristic_n := heuristic(start, end)
	heap.Push(&open, &Node{last: nil, pos: start, dir: [2]int{0,1}, value: 0, heuristic: heuristic_n})
	final_node := &Node{}
	
	for len(open) > 0 {
		current := heap.Pop(&open).(*Node)
		currentState := [4]int{current.pos[0], current.pos[1], current.dir[0], current.dir[1]}
		cost := bestCost[currentState]

		if current.value > cost {
			continue
		}
		matrix_copy[current.pos[0]][current.pos[1]] = "O"

		bestCost[currentState] = current.value

		if current.pos == end {
			matrix_copy[current.pos[0]][current.pos[1]] = "X"
			if final_node.value == 0 || current.value < final_node.value {
				final_node = current
			}
			// add break if a star
		}
		// first check if we can move in the same direction
		// Forward
		nextPos := [2]int{current.pos[0] + current.dir[0], current.pos[1] + current.dir[1]}
		if matrix[nextPos[0]][nextPos[1]] != "#" {
			newCost := current.value + 1
			neighborState := [4]int{nextPos[0], nextPos[1], current.dir[0], current.dir[1]}

			if newCost < bestCost[neighborState] {
				bestCost[neighborState] = newCost
				heap.Push(&open, &Node{
					last:      current,
					pos:       nextPos,
					dir:       current.dir,
					value:     newCost,
					heuristic: heuristic(nextPos, end),
				})
			}
		}
	
		// then rotate left and right
		rotatedDirs := [2][2]int{
			{current.dir[1], -current.dir[0]},  // left
			{-current.dir[1], current.dir[0]},  // right
		}

		for _, ndir := range rotatedDirs {
			newCost := current.value + 1000
			neighborState := [4]int{current.pos[0], current.pos[1], ndir[0], ndir[1]}

			if newCost < bestCost[neighborState] {
				bestCost[neighborState] = newCost
				heap.Push(&open, &Node{
					last:      current,
					pos:       current.pos,
					dir:       ndir,
					value:     newCost,
					heuristic: heuristic(current.pos, end),
				})
			}
		}
	}

	final := final_node.value
	for final_node.last != nil {
		matrix_copy[final_node.pos[0]][final_node.pos[1]] = "*"
		final_node = final_node.last
	}
	for _, line := range matrix_copy {
		fmt.Println(strings.Join(line, ""))
	}

	backtrace(matrix, bestCost, start, end, final)
}

//part 2 
func backtrace(matrix [][]string, bestCost map[[4]int]int, start, end [2]int, final int) {
	counted := map[[4]int]bool{}

	var rec func(r, c int, dir [2]int)
	rec = func(r, c int, dir [2]int) {
		if r < 0 || r >= len(matrix) || c < 0 || c >= len(matrix[0]) {
			return
		}
		if matrix[r][c] == "#" {
			return
		}
		if counted[[4]int{r, c, dir[0], dir[1]}] {
			return
		}

		best := bestCost[[4]int{r, c, -dir[0], -dir[1]}]	
		if best == math.MaxInt32 {
			return
		}

		counted[[4]int{r, c, dir[0], dir[1]}] = true		

		if r == start[0] && c == start[1] {
			return
		}

		rotate := [2][2]int{
			{dir[1], -dir[0]},  // left
			{-dir[1], dir[0]},  // right
		}

		for _, ndir := range rotate {
			neighborState := [4]int{r, c, ndir[0], ndir[1]}
			if bestCost[neighborState] == best - 1000 {
				rec(r, c, [2]int{-ndir[0], -ndir[1]})
			}
		}

		if bestCost[[4]int{r + dir[0], c + dir[1], -dir[0], -dir[1]}] == best - 1 {
			rec(r + dir[0], c + dir[1], dir)
		}
	}

	for _, d := range [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		fmt.Println(bestCost[[4]int{end[0], end[1], -d[0], -d[1]}], final)
		if bestCost[[4]int{end[0], end[1], -d[0], -d[1]}] == final {
			rec(end[0], end[1], d)
		}
	}

	for k := range counted {
		matrix[k[0]][k[1]] = "@"
	}
	for _, line := range matrix {
		fmt.Println(strings.Join(line, ""))
	}
	fmt.Println(countUniquePositions(counted))
}

func countUniquePositions(hashmap map[[4]int]bool) int {
    uniquePositions := make(map[[2]int]bool)

    for key := range hashmap {
        pos := [2]int{key[0], key[1]} // Extract only the position part
        uniquePositions[pos] = true  // Add to unique set
    }

    return len(uniquePositions) // Number of unique positions
}


func parseData(data string) ([][]string, [2]int, [2]int ) {
	matrix := [][]string{}
	lines := strings.Split(data, "\r\n")
	var start, end [2]int
	
	for i, line := range lines {
		matrix = append(matrix, []string{})
		for j, char := range line {
			matrix[i] = append(matrix[i], string(char))
			if char == 'S' {
				start = [2]int{i, j}
			} else if char == 'E' {
				end = [2]int{i, j}
			}
		}
	}

	return matrix, start, end
}

func main () {
	data := getData("input.txt")
	fmt.Println(data)
	matrix, start, end := parseData(data)
	fmt.Println(start, end)
	astar(matrix, start, end)
}