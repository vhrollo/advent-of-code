package main

import (
    "fmt"
    "os"
    "strings"
)

func get_data(path string) (string) {
    var content []byte
    var err error
    
    content, err = os.ReadFile(path)
    if err != nil {
        fmt.Println(err)
    }

    return strings.TrimRight(string(content), "\n") // for some reason it added a newline
}

func findWord(matrix [][]string, i int, j int) int {
    occurrences := 0
    word := "MAS"
    maxX := len(matrix)
    maxY := len(matrix[0])
    wordLen := len(word)
    directions := [][2]int{
        {-1, 0}, {1, 0},  // Up, Down
		{0, -1}, {0, 1},  // Left, Right
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1}, // Diagonals
    }

    for _, dir := range directions {
        k := 0
        x, y := i, j        
        
        for k < wordLen {
            newX := x + dir[0]
            newY := y + dir[1]

            if newX < 0 || newX >= maxX || newY < 0 || newY >= maxY {
                break
            }

            if string(matrix[newX][newY]) != string(word[k]) {
                break
            }

            x, y = newX, newY
            k++
        }

        if k == wordLen {
            occurrences++
        }
    }
    return occurrences
}

func findXWord(matrix [][]string, i int, j int) int {
    if i == 0 || i == len(matrix) - 1 || j == 0 || j == len(matrix[0]) - 1 {
        return 0
    }
    diagonals := [2][2][2]int{
        {{i - 1, j - 1}, {i + 1, j + 1}},
        {{i - 1, j + 1}, {i + 1, j - 1}},
    }
   
    for _, diag := range diagonals {
        foundM, foundS := false, false 
        for _, pos := range diag {
            x, y := pos[0], pos[1]
            if matrix[x][y] == "M" { foundM = true }
            if matrix[x][y] == "S" { foundS = true }
        }
        if foundM && foundS {
            continue
        }
        return 0
    }
    return 1
}

func part1(matrix [][]string) {
    sum := 0
    for i, row := range matrix {
		for j, char := range row {
			if char == "X" {
                sum += findWord(matrix, i, j)
            }
		}
	}
    fmt.Println(sum)
}

func part2(matrix [][]string) {
    sum := 0
    for i, row := range matrix {
		for j, char := range row {
			if char == "A" {
                sum += findXWord(matrix, i, j)
            }
		}
	}
    fmt.Println(sum)
}


func main() {
    data := get_data("test.txt")
    
    rows := strings.Split(data, "\n")

    matrix := make([][]string, len(rows))

    for i, row := range rows {
        matrix[i] = strings.Split(row, "")
    }
    //part1(matrix)
    part2(matrix)
}
