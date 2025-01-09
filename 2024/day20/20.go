package main

import (
	"fmt"
	"os"
	"strings"
	"container/heap"
)

type Node struct {
	last *Node
	pos [2]int
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
	*pq = old[0:n-1]
	return x
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func heuristic(a, b [2]int) int {
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

func parseData(data string) ([][]bool, [2]int, [2]int) {
	lines := strings.Split(data, "\n")
	var grid [][]bool
	var start, end [2]int
	for i, line := range lines {
		var row []bool
		for j, char := range line {
			row = append(row, char == '#')
			if char == 'S' {
				start = [2]int{i, j}
			}
			if char == 'E' {
				end = [2]int{i, j}
			}
		}
		grid = append(grid, row)
	}
	return grid, start, end
}



func astar(grid [][]bool, start [2]int, end [2]int) *Node{
	pq := make(PriorityQueue, 1)
	pq[0] = &Node{nil, start, 0, heuristic(start, end)}
	heap.Init(&pq)
	visited := make(map[[2]int]bool)
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Node)
		if current.pos == end {
			return current
		}
		visited[current.pos] = true
		neighbours := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, neighbour := range neighbours {
			newPos := [2]int{current.pos[0] + neighbour[0], current.pos[1] + neighbour[1]}
			if newPos[0] < 0 || newPos[1] < 0 || newPos[0] >= len(grid) || newPos[1] >= len(grid[0]) {
				continue
			}
			if grid[newPos[0]][newPos[1]] {
				continue
			}
			if visited[newPos] {
				continue
			}
			heap.Push(&pq, &Node{current, newPos, current.value + 1, heuristic(newPos, end)})
		}
	}
	return nil
}


func part1(grid [][]bool, start [2]int, end [2]int) {
	fnode := astar(grid, start, end)
	fmt.Println(fnode.value)

	grid_path := make([][]int, len(grid))
	for i := range grid {
		grid_path[i] = make([]int, len(grid[i]))
		for j := range grid_path[i] {
			grid_path[i][j] = -1
		}
	}

	node := fnode
	for node != nil {
		grid_path[node.pos[0]][node.pos[1]] = node.value
		node = node.last
	}


	saved := make(map[int]int)
	dir := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	node = fnode
	for node != nil {
		for _, d := range dir {
			newPos := [2]int{node.pos[0] + d[0], node.pos[1] + d[1]}
			newPos2 := [2]int{node.pos[0] + 2*d[0], node.pos[1] + 2*d[1]}
			if newPos2[0] < 0 || newPos2[1] < 0 || newPos2[0] >= len(grid) || newPos2[1] >= len(grid[0]) {
				continue
			}
			if !grid[newPos[0]][newPos[1]] || grid[newPos2[0]][newPos2[1]] {
				continue
			}
			saves := grid_path[node.pos[0]][node.pos[1]] - grid_path[newPos2[0]][newPos2[1]] - 2
			if saves < 0 {
				continue
			}
			saved[saves]++
		}
		node = node.last
	}

	thresh := 100
	count := 0
	for k := range saved {
		if k < thresh {
			continue
		}
		count += saved[k]
	}
	fmt.Println(count)	
}

type floodNode struct {
	pos [2]int
	dist int
	usedCheat int
}

func floodfill(grid [][]bool, node *Node, grid_path [][]int, saved *map[int]int) {

	startPos := node.pos
	startVal := grid_path[startPos[0]][startPos[1]]
	
	dir := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	visited := make(map[[2]int]bool)
	queue := []floodNode{{startPos, 0, 0}}
	visited[[2]int{startPos[0], startPos[1]}] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.dist > 0 && !grid[current.pos[0]][current.pos[1]] {
			savedVal := startVal - grid_path[current.pos[0]][current.pos[1]] - current.dist

			if savedVal > 0 {
				(*saved)[savedVal]++
			}
		}

		for _, d := range dir {
			newPos := [2]int{current.pos[0] + d[0], current.pos[1] + d[1]}

			if newPos[0] < 0 || newPos[1] < 0 || newPos[0] >= len(grid) || newPos[1] >= len(grid[0]) {
				continue
			}

			newDist := current.dist + 1
			newUsed := current.usedCheat + 1

			if newDist > 20 {
				continue
			}

			stateKey := [2]int{newPos[0], newPos[1]}
			if visited[stateKey] {
				continue
			}

			visited[stateKey] = true

			queue = append(queue, floodNode{newPos, newDist, newUsed})
		}
	}
}


func part2(grid [][]bool, start [2]int, end [2]int) {
	fnode := astar(grid, start, end)
	
	grid_path := make([][]int, len(grid))
	for i := range grid {
		grid_path[i] = make([]int, len(grid[i]))
		for j := range grid_path[i] {
			grid_path[i][j] = -1
		}
	}

	node := fnode
	for node != nil {
		grid_path[node.pos[0]][node.pos[1]] = node.value
		node = node.last
	}

	saved := make(map[int]int)

	// flood fill
	node = fnode
	for node != nil {
		floodfill(grid, node, grid_path, &saved)
		node = node.last
	}
	threshold := 100
	count := 0
	// savedList := make([][2]int, 0)
	for k := range saved {
		if k < threshold {
			continue
		}
		count += saved[k]
		// savedList = append(savedList, [2]int{k, saved[k]})
	}
	
	//sort
	// for i := 0; i < len(savedList); i++ {
	// 	for j := i; j < len(savedList); j++ {
	// 		if savedList[i][0] > savedList[j][0] {
	// 			savedList[i], savedList[j] = savedList[j], savedList[i]
	// 		}
	// 	}
	// }
	// for _, s := range savedList {
	// 	fmt.Println(s)
	// }

	fmt.Println(count)
}

func main() {
	data := getData("input.txt")
	grid, start, end := parseData(data)
	fmt.Println(start, end)
	part1(grid, start, end)
	part2(grid, start, end)
}