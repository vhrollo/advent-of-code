package main

import (
    "fmt"
    "os"
    "strings"
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

func parse(data string) ( [][2]int, [][]int){
    splitted := strings.Split(data, "\n\n") 

    //rules parse
    rules_raw := strings.Split(splitted[0], "\n")
    rules := make([][2]int, 0, len(rules_raw))

    for _, rule := range rules_raw {
        parts := strings.Split(rule, "|")
        nr1, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
        nr2, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

        if err1 != nil || err2 != nil {
            fmt.Println("uhoh")
        }

        rules = append(rules, [2]int{nr1, nr2})
    }


    //updates parse
    updates_raw := strings.Split(splitted[1], "\n")
    updates := make([][]int, 0, len(updates_raw))

    for _, update := range updates_raw {

        parts := strings.Split(update, ",")
        currentUpdate := make([]int, 0, len(parts))

        for _, part := range parts {
            num, err := strconv.Atoi(strings.TrimSpace(part))
            if err != nil {
                fmt.Println("uhoh.v2")
            }
            currentUpdate = append(currentUpdate, num)
        }

        updates = append(updates, currentUpdate)
    }

    return rules, updates
}


//dependency map
func preproccessRules(rules [][2]int) map[int]map[int]struct{} {
    mustComeAfter := make(map[int]map[int]struct{})

    for _, rule := range rules {
        A, B := rule[0], rule[1]
        
        if _, exists := mustComeAfter[B]; !exists {
            mustComeAfter[B] = make(map[int]struct{})
        }
        mustComeAfter[B][A] = struct{}{}
    }
    return mustComeAfter
}

func identifyCorrectUpdate(dep map[int]map[int]struct{}, updates [][]int) ([]int, []int) {        
    correct := []int{}
    inc := []int{}

    for i, update := range updates {
       
        // so i can map each page to its idx
        pagePosition := make(map[int]int)
        for idx, page := range update {
            pagePosition[page] = idx
        }

        isCorrect := true

        for idx, page := range update {
            dependencies, exists := dep[page]    

            if !exists  {
                continue
            }

            for depPage := range dependencies {
                depPos, existsDep := pagePosition[depPage]
                if !existsDep {
                    continue
                }

                if depPos > idx {
                    isCorrect = false
                    break
                }
            }

            if !isCorrect {break}
        }

        if isCorrect {
            correct = append(correct, i)
        } else {
            inc = append(inc, i)
        }

    }
    return correct, inc
}

func contains(slice []int, item int) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

func topoSort(pages []int, dep map[int]map[int]struct{}) []int {
    graph := make(map[int][]int)
    inDegree := make(map[int]int)

    for _, page := range pages {
        inDegree[page] = 0
    }

    for _, page := range pages {
        for depPage := range dep[page] {
            if contains(pages, depPage) {
                graph[depPage] = append(graph[depPage], page)
                inDegree[page]++
            }
        }
    }

    queue := []int{}
    for _, page := range pages {
       if inDegree[page] == 0 {
            queue = append(queue, page)
        }
    }

    sorted := []int{}

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        sorted = append(sorted, current)

        for _, neighbor := range graph[current] {
            inDegree[neighbor]--
            if inDegree[neighbor] == 0 {
                queue = append(queue, neighbor)
            }
        }
    }

    if len(sorted) != len(pages) {
        fmt.Println("cycle")
    }

    return sorted
}



func sumMiddleVals(idxs []int, updates [][]int) {
    sum := 0
    for _, i := range idxs {
        update := updates[i] 
        uplen  := len(update)
        middle := uplen / 2 
        sum += update[middle]
    }
    fmt.Println(sum)
}

func main() {
    data := get_data("input.txt")
    rules, updates := parse(data)

    depmap := preproccessRules(rules)
    correct, inc := identifyCorrectUpdate(depmap, updates)
    fmt.Println(correct)
    sumMiddleVals(correct, updates)
    
    part2 := 0
    for _, i := range inc {
        sorted := topoSort(updates[i], depmap) 
        uplen  := len(sorted)
        middle := uplen / 2 
        part2 += sorted[middle]
    }
    fmt.Println(part2)
}
