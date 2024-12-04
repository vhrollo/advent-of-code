package main

import (
    "strings"
    "os"
    "fmt"
    "strconv"
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

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isSafe(reports []int) bool {
    delta := reports[1] - reports[0]

    if intAbs(delta) < 1 || intAbs(delta) > 3 {
        return false
    }

    direction := 1
    if delta < 0 {
        direction = -1
    }


    for i:=1; i < len(reports)-1; i++ {
        currentDelta := reports[i+1] - reports[i]

        if (currentDelta > 0 && direction < 0) || (currentDelta < 0 && direction > 0) {
            return false
        }

        if intAbs(currentDelta) < 1 || intAbs(currentDelta) > 3 {
            return false
        }
    }
    return true
}

func part1(table [][]int) {
    var sum int

    for _, reports := range table {
        if isSafe(reports) { 
            sum++
        }
    }

    fmt.Println(sum)
}

func part2(table [][]int) {
    var sum int

    for _, reports := range table {
        if isSafe(reports) {
            sum++
            continue
        }

        //fmt.Println("line:", j)
        //fmt.Println("original", reports)
        for i := 0; i < len(reports); i++ {
            copied := []int{}
            copied = append(copied, reports[:i]...)
            copied= append(copied, reports[i+1:]...)
            //fmt.Println("org",reports)
            //fmt.Println(copied)
            if isSafe(copied) {
                sum++
                break
            }
        }
    }

    fmt.Println(sum)
}


func main() {
    var data string = get_data("input.txt")

    var lines []string = strings.Split(data, "\n")

    var table [][]int

    for _, report := range lines {
        var levels []string = strings.Fields(report)
        
        row := make([]int, len(levels))

        for j, level := range levels {
            num, err := strconv.Atoi(level)
            if err != nil {
                fmt.Println("uhoh")
                os.Exit(0)
            }
            row[j] = num
        }
        table = append(table,row)
    }

    //part1(table)
    part2(table)
}
