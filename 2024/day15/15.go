package main

import (
	"fmt"
	"os"
	"strings"
)

func getData(path string) string {
    content, err := os.ReadFile(path)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return strings.TrimRight(string(content), "\n")
}

func parseData(data string) ([][]string, [][]string, [][2]int, [2]int, [2]int) {
	splittedData := strings.Split(data, "\r\n\r\n")

	player := [2]int{0, 0}
	
	mLines := strings.Split(splittedData[0], "\r\n")
	matrix := make([][]string, len(mLines))
	for i, line := range mLines {
		matrix[i] = strings.Split(line, "")
		for j, s := range matrix[i] {
			if s == "@" {
				player = [2]int{i, j}
				//matrix[i][j] = "."
			}
		}

	}

	matrix2 := make([][]string, len(mLines))
	player2 := [2]int{0, 0}
	for i, line := range mLines {
		row := strings.Split(line, "")
		newRow := []string{}
		for j, s := range row {
			if s == "@" {
				newRow = append(newRow, "@")
				newRow = append(newRow, ".")
				player2 = [2]int{i, j*2}
			}
			if s == "." {
				newRow = append(newRow, ".")
				newRow = append(newRow, ".")
			}
			if s == "#" {
				newRow = append(newRow, "#")
				newRow = append(newRow, "#")
			}
			if s == "O" {
				newRow = append(newRow, "[")
				newRow = append(newRow, "]")
			}
		}
		matrix2[i] = newRow
	}


	sInst := strings.Split(strings.ReplaceAll(splittedData[1], "\r\n", ""), "")
	inst := make([][2]int, len(sInst))
	for i, s := range sInst {
		if s == "^" {
			inst[i] = [2]int{-1, 0}
		} else if s == "v" {
			inst[i] = [2]int{1, 0}
		} else if s == ">" {
			inst[i] = [2]int{0, 1}
		} else if s == "<" {
			inst[i] = [2]int{0, -1}
		}
	}

	
	return matrix, matrix2, inst, player, player2
}

func RecMove(matrix *[][]string, inst [2]int, player *[2]int) {

	swap := func (a, b [2]int) {
		(*matrix)[a[0]][a[1]], (*matrix)[b[0]][b[1]] = (*matrix)[b[0]][b[1]], (*matrix)[a[0]][a[1]]
	}

	var move func([2]int, [2]int) bool
	move = func(pos, dir [2]int) bool {
		nextPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		next := (*matrix)[nextPos[0]][nextPos[1]]
		if next == "#" {
			return false
		}
		if next == "." {
			swap(pos, nextPos)
			return true
		}
		if next == "O" {
			movable := move(nextPos, dir)
			if movable {
				swap(pos, nextPos)
				return true
			}
		}
		return false
	}

	if move(*player, inst) {
		(*player)[0] += inst[0]
		(*player)[1] += inst[1]
	}
}

func print (matrix [][]string) {
	for _, line := range matrix {
		fmt.Println(strings.Join(line, ""))
	}
}

func part1 (matrix [][]string, inst [][2]int, player [2]int) {
	for _, i := range inst {
		RecMove(&matrix, i, &player)
		// for _, line := range matrix {
		// 	fmt.Println(strings.Join(line, ""))
		// }
	}

	sum := 0
	for i, line := range matrix {
		for j, s := range line {
			if s == "O" {
				sum += 100*i+j
			}
		}
	}
	fmt.Println("Part 1:", sum)
}

func RecMove2(matrix *[][]string, inst [2]int, player *[2]int) {
	visited := make(map[[2]int]bool)

	swap := func (a, b [2]int) {
		visited[a] = false
		(*matrix)[a[0]][a[1]], (*matrix)[b[0]][b[1]] = (*matrix)[b[0]][b[1]], (*matrix)[a[0]][a[1]]
	}

	var checkmove func([2]int, [2]int) bool
	checkmove = func(pos, dir [2]int) bool {
		nextPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		next := (*matrix)[nextPos[0]][nextPos[1]]

		if visited[pos] {
			return visited[pos]
		}
		if next == "#" {
			return false
		}
		if next == "." {
			visited[pos] = true
			return true
		}
		if dir[0] == 0 && (next == "[" || next == "]") {
			if checkmove(nextPos, dir) {
				visited[pos] = true
				return true
			} else {return false}
		}
		if dir[1] == 0 && next == "[" {
			nextPos2 := [2]int{nextPos[0], nextPos[1]+1}
			if checkmove(nextPos, dir) && checkmove(nextPos2, dir){
				visited[pos] = true
				return true
			} else {return false}
		}
		if dir[1] == 0 && next == "]" {
			nextPos2 := [2]int{nextPos[0], nextPos[1]-1}
			if checkmove(nextPos, dir) && checkmove(nextPos2, dir){
				visited[pos] = true
				return true
			} else {return false}
		}
		return false
	}

	var move func([2]int, [2]int)
	move = func(pos, dir [2]int) {
		nextPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		next := (*matrix)[nextPos[0]][nextPos[1]]
		if next == "." {
			swap(pos, nextPos)
		} else if dir[0] == 0 && (next == "[" || next == "]") {
			move(nextPos, dir)
			swap(pos, nextPos)
		} else if dir[1] == 0 && next == "[" {
			nextPos2 := [2]int{nextPos[0], nextPos[1]+1}
			move(nextPos, dir)
			move(nextPos2, dir)
			swap(pos, nextPos)
		} else if dir[1] == 0 && next == "]" {
			nextPos2 := [2]int{nextPos[0], nextPos[1]-1}
			move(nextPos, dir)
			move(nextPos2, dir)
			swap(pos, nextPos)
		}
	}


	if checkmove(*player, inst) {
		move(*player, inst)
		player[0] += inst[0]
		player[1] += inst[1]
	}
}

func part2 (matrix [][]string, inst [][2]int, player [2]int) {
	for _, i := range inst {
		RecMove2(&matrix, i, &player)
	}
	sum := 0
	for i, line := range matrix {
		for j, s := range line {
			if s == "[" {
				sum += 100*i+j
			}
		}
	}
	fmt.Println("Part 2:", sum)
}

func main() {
	data := getData("input.txt")
	matrix, matrix2, inst, player, player2 := parseData(data)
	part1(matrix, inst, player)
	part2(matrix2, inst, player2)
}