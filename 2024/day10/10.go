package main

import (
    "fmt"
    "os"
    "strings"
)

type Cell struct {
    checked  bool
    hilltops map[[2]int]struct{}
    paths int
}

func getData(path string) string {
    content, err := os.ReadFile(path)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return strings.TrimRight(string(content), "\n")
}

func buildMatrix(dataStr string) [][]int {
    rows := strings.Split(dataStr, "\r\n")
    if len(rows) == 1 {
        rows = strings.Split(dataStr, "\n")
    }

    matrix := make([][]int, len(rows))
    for i, row := range rows {
        cols := strings.Split(row, "")
        matrix[i] = make([]int, len(cols))
        for j, c := range cols {
            matrix[i][j] = int(c[0] - '0')
        }
    }
    return matrix
}

func fillNeighbors(matrix [][]int, dp *[][]Cell, i, j int) {
    (*dp)[i][j].checked = true

    directions := []struct{ di, dj int }{
        {-1, 0}, // up
        {1, 0},  // down
        {0, -1}, // left
        {0, 1},  // right
    }

    for _, d := range directions {
        ni, nj := i+d.di, j+d.dj

        // boundry check
        if ni < 0 || ni >= len(matrix) || nj < 0 || nj >= len(matrix[0]) {
            continue
        }

        if matrix[ni][nj] == matrix[i][j]+1 {
            neighborCell := &(*dp)[ni][nj]
            if !neighborCell.checked {
                fillNeighbors(matrix, dp, ni, nj)
                for hilltop := range neighborCell.hilltops {
                    (*dp)[i][j].hilltops[hilltop] = struct{}{}
                }
                (*dp)[i][j].paths += neighborCell.paths
            } else {
                for hilltop := range neighborCell.hilltops {
                    (*dp)[i][j].hilltops[hilltop] = struct{}{}
                }
                (*dp)[i][j].paths += neighborCell.paths
            }
        }
    }
}

func solve(matrix [][]int) {
    dp := make([][]Cell, len(matrix))
    for i := range dp {
        dp[i] = make([]Cell, len(matrix[i]))
        for j := range dp[i] {
            dp[i][j] = Cell{
                checked:  false,
                hilltops: make(map[[2]int]struct{}),
                paths:    0,
            }
        }
    }

    for i := range matrix {
        for j := range matrix[i] {
            if matrix[i][j] == 9 {
                dp[i][j].checked = true
                dp[i][j].hilltops[[2]int{i, j}] = struct{}{}
                dp[i][j].paths = 1
            }
        }
    }

    var trailheads [][2]int
    for i := range matrix {
        for j := range matrix[i] {
            if matrix[i][j] == 0 {
                trailheads = append(trailheads, [2]int{i, j})
            }
        }
    }

    for _, th := range trailheads {
        fillNeighbors(matrix, &dp, th[0], th[1])
    }

    totalScore := 0
    totalPaths := 0
    for _, th := range trailheads {
        totalScore += len(dp[th[0]][th[1]].hilltops)
        totalPaths += dp[th[0]][th[1]].paths
    }

    fmt.Println(totalScore)
    fmt.Println(totalPaths)
}

func main() {
    data := getData("input.txt")
    matrix := buildMatrix(data)
    solve(matrix)
}
