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

func parse_data(dataStr string) [][]string {
    rows := strings.Split(dataStr, "\n")
    matrix := make([][]string, len(rows))

    for i, row := range rows {
        matrix[i] = strings.Split(row, "")
    }
    return matrix
}


func countAtSymbols(matrix [][]string) int {
    count := 0
    for _, row := range matrix {
        for _, cell := range row {
            if cell == "@" {
                count++
            }
        }
    }
    return count
}


func part1(matrix [][]string) {
    // pos down i 
    // pos right j
    dir := [2]int{0, -1}

    x, y := -1, -1
    maxY, maxX := len(matrix), len(matrix[0])


    loop:
    for i, row := range matrix {
        for j, cell := range row {
            if cell == "^" {
                y, x = i, j
                break loop
            }
        }
    }
    
    for {
        newY := y + dir[1]
        newX := x + dir[0]

        if newY < 0 || newY >= maxY || newX < 0 || newX >= maxX {
            matrix[y][x] = "@"
            break
        }

        if matrix[newY][newX] == "#" {
            dir = [2]int{-dir[1], dir[0]}
            continue
        } 

        matrix[y][x] = "@"

        y = newY
        x = newX
    }

    count := countAtSymbols(matrix)
    fmt.Printf("Total '@' symbols: %d\n", count)
}



func main() {
    data := get_data("test.txt")
    matrix := parse_data(data)
    part1(matrix)
}
