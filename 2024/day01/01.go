package main

import (
    "strings"
    "fmt"
    "os"
    "sort"
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

func part1(col1, col2 []int) {

    sort.Ints(col1)    
    sort.Ints(col2)    

    var abs int

    for i := 0; i < len(col1); i++ {
        abs += intAbs(col1[i] - col2[i])
    }

    fmt.Println(abs)

}

func part2(col1, col2 []int) {
    counts := make(map[int]int) 
    for _, val := range col2 {
        counts[val]++
    }

    var sum int

    for _, val := range col1 {
        count, exists := counts[val]
        if exists {
            sum += val * count
        }
    }

    fmt.Println(sum)
}

func main() {
    var data string = get_data("input.txt")

    var lines []string = strings.Split(data, "\n")

    var col1 []int
    var col2 []int

    for _, line := range lines {
        var parts []string = strings.Fields(line)
        var value1, value2 int 
        
        fmt.Sscanf(parts[0], "%d", &value1)
        fmt.Sscanf(parts[1], "%d", &value2)

        col1 = append(col1, value1)
        col2 = append(col2, value2)
    }

    //part1(col1, col2)
    part2(col1, col2)
}
