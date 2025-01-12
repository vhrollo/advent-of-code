package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var NUMERIC = []string{
	"789",
	"456",
	"123",
	" 0A",
}

var DIRECTIONAL = []string{
	" ^A",
	"<v>",
}

func walk(keypad []string, x, y int, path string) bool {
	for _, r := range path {
		switch r {
		case '^': y--
		case 'v': y++
		case '<': x--
		case '>': x++
		}
		if x<0 || y<0 || y>=len(keypad) || x>=len(keypad[0]) || keypad[y][x]==' ' {
			return true
		}
	}
	return false
}

func findCoord(keypad []string, key rune) (int, int) {
	for y, row := range keypad {
		for x, r := range row {
			if r == key {
				return x, y
			}
		}
	}
	return -1, -1
}

func pathsBetween(keypad []string, start, end rune) []string {
	x1, y1 := findCoord(keypad, start)
	x2, y2 := findCoord(keypad, end)

	dx := x2 - x1
	var hor string
	if dx > 0 {
		hor = strings.Repeat(">", dx)
	} else {
		hor = strings.Repeat("<", -dx)
	}

	dy := y2 - y1
	var ver string
	if dy > 0 {
		ver = strings.Repeat("v", dy)
	} else {
		ver = strings.Repeat("^", -dy)
	}

	candidates := []string{hor + ver, ver + hor}
	var result []string
	for _, path := range candidates {
		if walk(keypad, x1, y1, path) {
			continue
		}
		result = append(result, path+"A")
	}
	return result
}


var memoCostBetween sync.Map

type costBetweenKey struct {
	KeypadID string
	start, end rune
	links int
}

func costBetween(keypad []string, keypadID string, start, end rune, links int) int {
	k := costBetweenKey{keypadID, start, end, links}
	if val, ok := memoCostBetween.Load(k); ok {
		return val.(int)
	}
	
	if links == 0 {
		memoCostBetween.Store(k, 1)
		return 1
	}

	minVal := int(^uint(0) >> 1) // inf
	for _, path := range pathsBetween(keypad, start, end) {
		c := cost(DIRECTIONAL, "DIRECTIONAL", path, links-1)
		if c < minVal {
			minVal = c
		}
	}
	memoCostBetween.Store(k, minVal)
	return minVal
}

var memoCost sync.Map

type costKey struct {
	KeypadID string
	keys     string
	links    int
}

func cost(keypad []string, keypadID, keys string, links int) int {
	k := costKey{keypadID, keys, links}
	if val, ok := memoCost.Load(k); ok {
		return val.(int)
	}
	r := []rune("A"+keys)
	total := 0
	for i:=0; i<len(r)-1; i++ {
		total += costBetween(keypad,keypadID,r[i],r[i+1],links)
	}
	memoCost.Store(k, total)
	return total
}

func complexity(code string, robots int) int {
	val, _ := strconv.Atoi(code[:len(code) - 1])
	return val * cost(NUMERIC, "NUMERIC", code, robots+1)
}

func main() {
	dat, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	sum2, sum25 := 0, 0
	for _, line := range lines {
		sum2  += complexity(line, 2)
		sum25 += complexity(line, 25)
	}
	fmt.Println(sum2)
	fmt.Println(sum25)
}