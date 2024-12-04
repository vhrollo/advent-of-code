package main

import (
    "strings"
    "os"
    "fmt"
    "regexp"
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

func part1(text string) {
    pattern := `mul\((\d+),(\d+)\)` 

    re := regexp.MustCompile(pattern)
    matches := re.FindAllStringSubmatch(text, -1)
    
    sum := 0

    for _, match := range matches {
        num1, err1 := strconv.Atoi(match[1])
        num2, err2 := strconv.Atoi(match[2])
        if err1 != nil || err2 != nil {
            fmt.Println("uhoh")
        }
        sum += num1 * num2
    }
    fmt.Println(sum)
}

func part2(text string) {
    //basically gonna remove all 
    // (?s:.*?) bc regex in go doesnt searches newline by default
    // could've also doen this sequentially, but nah
    pattern := `don't\(\)(?s:.*?)(do\(\)|$)`
    
    re := regexp.MustCompile(pattern)
    new_text := re.ReplaceAllString(text, "#")

    //fmt.Println(new_text)
    part1(new_text)
}


func main() {
    text := get_data("input.txt")

    //part1(text)
    part2(text)
}
