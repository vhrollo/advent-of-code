// 30 min max to solve this problem, suprising that only 700 people solved it!


package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func toInt (a string) int {
	b, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return b
}

func astar (grid [][]bool, start, end [2]int) int {
	pq := make(PriorityQueue, 1)
	pq[0] = &Node{nil, start, 0, heuristic(start, end)}
	heap.Init(&pq)
	visited := make(map[[2]int]bool)
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Node)
		if current.pos == end {
			return current.value
		}
		visited[current.pos] = true
		neighbours := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, neighbour := range neighbours {
			newPos := [2]int{current.pos[0] + neighbour[0], current.pos[1] + neighbour[1]}
			if newPos[0] < 0 || newPos[1] < 0 || newPos[0] >= len(grid) || newPos[1] >= len(grid[0]) {
				continue
			}
			if grid[newPos[1]][newPos[0]] {
				continue
			}
			if visited[newPos] {
				continue
			}
			heap.Push(&pq, &Node{current, newPos, current.value + 1, heuristic(newPos, end)})
		}
	}
	return -1
}

func part1 (data string) { 
	grid := parseData(data, 1024)
	fmt.Println(astar(grid, [2]int{0, 0}, [2]int{70, 70}))
}

func part2 (data string) {
	splitted := strings.Split(data, "\n")
	datalen := len(splitted)
	for i := 1; i < datalen; i++ {
		grid := parseData(data, i)
		if astar(grid, [2]int{0, 0}, [2]int{70, 70}) == -1 {
			fmt.Println(splitted[i-1])
			break
		}
	}
}

func getData(path string) string {
    content, err := os.ReadFile(path)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return strings.TrimRight(string(content), "\n")
}

func parseData(data string, n int) [][]bool {
	// size := 6 //test
	size := 70 //input

	grid := make([][]bool, size+1)
	for i := range grid {
		grid[i] = make([]bool, size+1)
	}

	splitted := strings.Split(data, "\n")
	for i := 0; i < n; i++ {
		pos := strings.Split(splitted[i], ",")
		posInt := [2]int{toInt(pos[0]), toInt(pos[1])}
		grid[posInt[1]][posInt[0]] = true
	}
	return grid
}


func main() {
	data := getData("input.txt")
	part1(data)
	part2(data)
}