package main

import (
    "fmt"
    "os"
    "strings"
	"strconv"
    "math"
)

type Equation struct {
    TestValue int
    Numbers   []int
}


func get_data(path string) (string) {
    var content []byte
    var err error
    
    content, err = os.ReadFile(path)
    if err != nil {
        fmt.Println(err)
    }

    return strings.TrimRight(string(content), "\n") // for some reason it added a newline
}

func padding(n int) int {
    return int(math.Pow(10, math.Floor(math.Log10(float64(n))+1)))
}

func canAchieveValue(numbers []int, target int, allowConcat bool) bool {

    var recurse func(index int, current int) bool 
    recurse = func(index int, current int) bool {

        if current > target {
            return false
        }

        if index == len(numbers) {
            return current == target
        }

        nextNum := numbers[index]

        if recurse(index+1, current+nextNum) {
            return true
        }
        
        if recurse(index+1, current*nextNum) {
            return true
        }

        if allowConcat && recurse(index+1, current*padding(nextNum)+nextNum) {
            return true
        }

        return false
    }

    return recurse(1, numbers[0])
}



func parseData(dataStr string) ([]Equation, error) {
	var equations []Equation
	rows := strings.Split(dataStr, "\n")

	for _, row := range rows {
		parts := strings.Split(row, ":")

		testValueStr := strings.TrimSpace(parts[0])
		testValue, err := strconv.Atoi(testValueStr)
		if err != nil {
			return nil, fmt.Errorf("invalid test value '%s': %v", testValueStr, err)
		}

		numbersStr := strings.TrimSpace(parts[1])
        numberStrs := strings.Fields(numbersStr)
        var numbers []int
        for _, numStr := range numberStrs {
            num, err := strconv.Atoi(numStr)
            if err != nil {
                return nil, fmt.Errorf("invalid number '%s': %v", numStr, err)
            }
            numbers = append(numbers, num)
        }
        
        equations = append(equations, Equation{
            TestValue: testValue,
            Numbers:   numbers,
        })
	}
	return equations, nil
}

func solve(data []Equation, concat bool) {
    sum := 0
    for _, equation := range data {
        if canAchieveValue(equation.Numbers, equation.TestValue, concat) {
            sum += equation.TestValue
        }
    }
    fmt.Printf("count: %d\n", sum)
}


func main() {
	dataStr := get_data("input.txt")
    data, err := parseData(dataStr)
	if err != nil {
		fmt.Println(err)
		return
	}
    solve(data, true)
}