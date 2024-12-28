package main

import (
	"fmt"
	"os"
	"sort"
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

func buildMatrix(data string) [][]string {
    rows := strings.Split(data, "\r\n")
    matrix := make([][]string, len(rows))
    for i, row := range rows {
        matrix[i] = strings.Split(row, "")
    }
    return matrix
}

func bfsRegion(matrix [][]string, visited [][]bool, startR, startC int) [][2]int {
    R := len(matrix)
    C := len(matrix[0])
    letter := matrix[startR][startC]
    var region [][2]int

    neighbors := func(r, c int) [][2]int {
        dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
        var result [][2]int
        for _, d := range dirs {
            nr, nc := r+d[0], c+d[1]
            if nr >= 0 && nr < R && nc >= 0 && nc < C {
                result = append(result, [2]int{nr, nc})
            }
        }
        return result
    }

    queue := [] [2]int { {startR, startC} }
    visited[startR][startC] = true

    for len(queue) > 0 {
        r, c := queue[0][0], queue[0][1]
        queue = queue[1:]
        region = append(region, [2]int{r, c})
        
        for _, nbr := range neighbors(r, c) {
            nr, nc := nbr[0], nbr[1]
            if !visited[nr][nc] && matrix[nr][nc] == letter {
                visited[nr][nc] = true
                queue = append(queue, [2]int{nr, nc})
            }
        }
    }

    return region
}

func part1(matrix [][]string) {
    R := len(matrix)
    C := len(matrix[0])
    visited := make([][]bool, R)
    for i := 0; i < R; i++ {
        visited[i] = make([]bool, C)
    }

    totalPrice := 0

    for r := 0; r < R; r++ {
        for c := 0; c < C; c++ {
            if !visited[r][c] {
                regionCells := bfsRegion(matrix, visited, r, c)
                area := len(regionCells)

                // Perimeter calculation:
                perimeter := 0
                for _, cell := range regionCells {
                    rr, cc := cell[0], cell[1]
                    // Check up, down, left, right
                    dirs := [][2]int{{rr-1, cc}, {rr+1, cc}, {rr, cc-1}, {rr, cc+1}}
                    for _, d := range dirs {
                        nr, nc := d[0], d[1]
                        // If out of bounds or different letter, contributes to perimeter
                        if nr < 0 || nr >= R || nc < 0 || nc >= C || matrix[nr][nc] != matrix[rr][cc] {
                            perimeter++
                        }
                    }
                }

                totalPrice += area * perimeter
            }
        }
    }

    fmt.Println("Part 1:", totalPrice)
}



// Part 2
type horizontalEdge struct {
    row         int
    startColumn int
    endColumn   int
    regionBelow bool
}

type verticalEdge struct {
    col         int
    startRow    int
    endRow      int
    regionRight bool
}

func part2(matrix [][]string) {
    R := len(matrix)
    C := len(matrix[0])
    visited := make([][]bool, R)
    for i := 0; i < R; i++ {
        visited[i] = make([]bool, C)
    }

    totalPrice := 0

    for r := 0; r < R; r++ {
        for c := 0; c < C; c++ {
            if !visited[r][c] {
                regionCells := bfsRegion(matrix, visited, r, c)
                area := len(regionCells)
                
                hEdges, vEdges := collectBoundaryEdges(matrix, regionCells)
                mergedSides := len(mergeHorizontal(hEdges)) + len(mergeVertical(vEdges))

                totalPrice += area * mergedSides
            }
        }
    }

    fmt.Println("Part 2:", totalPrice)
}


func collectBoundaryEdges(grid [][]string, region [][2]int) ([]horizontalEdge, []verticalEdge) {
    R := len(grid)
    C := len(grid[0])

    regionSet := make(map[[2]int]bool, len(region))
    for _, rc := range region {
        regionSet[rc] = true
    }

    var hEdges []horizontalEdge
    var vEdges []verticalEdge

    for _, rc := range region {
        r, c := rc[0], rc[1]

        if c == 0 || !regionSet[[2]int{r, c - 1}] {
            vEdges = append(vEdges, verticalEdge{
                col:      c, 
                startRow: r,
                endRow:   r + 1,
                regionRight: true,
            })
        }

        if c+1 == C || !regionSet[[2]int{r, c + 1}] {
            vEdges = append(vEdges, verticalEdge{
                col:      c + 1,
                startRow: r,
                endRow:   r + 1,
                regionRight: false,
            })
        }

        if r == 0 || !regionSet[[2]int{r - 1, c}] {
            hEdges = append(hEdges, horizontalEdge{
                row:         r,
                startColumn: c,
                endColumn:   c + 1,
                regionBelow: true,
            })
        }

        if r+1 == R || !regionSet[[2]int{r + 1, c}] {
            hEdges = append(hEdges, horizontalEdge{
                row:         r + 1,
                startColumn: c,
                endColumn:   c + 1,
                regionBelow: false,
            })
        }
    }

    return hEdges, vEdges
}

// first sortt the edges to then merge them, this way we can merge them in one pass
// could be done better, but it got bloated

func mergeHorizontal(edges []horizontalEdge) []horizontalEdge {
    sort.Slice(edges, func(i, j int) bool {
        if edges[i].row == edges[j].row {
            return edges[i].startColumn < edges[j].startColumn
        }
        return edges[i].row < edges[j].row
    })

    var merged []horizontalEdge
    // this works since all rows are contiguous
    for _, e := range edges {
        n := len(merged)
        if n == 0 {
            merged = append(merged, e)
            continue
        }
        last := &merged[n-1]
        if last.row == e.row &&
           last.endColumn == e.startColumn &&
           last.regionBelow == e.regionBelow {

            last.endColumn = e.endColumn
        } else {
            merged = append(merged, e)
        }
    }
    return merged
}

func mergeVertical(edges []verticalEdge) []verticalEdge {
    sort.Slice(edges, func(i, j int) bool {
        if edges[i].col == edges[j].col {
            return edges[i].startRow < edges[j].startRow
        }
        return edges[i].col < edges[j].col
    })

    var merged []verticalEdge
    for _, e := range edges {
        n := len(merged)
        if n == 0 {
            merged = append(merged, e)
            continue
        }
        last := &merged[n-1]
        if last.col == e.col && 
           last.endRow == e.startRow &&
           last.regionRight == e.regionRight {

            last.endRow = e.endRow
        } else {
            merged = append(merged, e)
        }
    }
    return merged
}


func main() {
	dataStr := getData("input.txt")
    dataMatrix := buildMatrix(dataStr)
    part1(dataMatrix)
    part2(dataMatrix)
}