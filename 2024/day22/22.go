package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"sync"
)

const magic = 0x1000000 // 16777216

type sequence [4]int

func getData() []int {
	dat, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	intLines := make([]int, len(lines))
	for i, line := range lines {
		intLines[i], _ = strconv.Atoi(line)
	}
	return intLines
}

func step(S int) int {
	S = (S ^ (S <<  6)) % magic
	S = (S ^ (S >>  5)) % magic
	S = (S ^ (S << 11)) % magic
	return S
}

func part1(data []int) {
	total := 0
	for _, val := range data {
		for i := 0; i < 2000; i++ {
			val = step(val)
		}
		total += val
	}
	fmt.Println("Part1:", total)
}

func generatePricesAndChanges(initial int, n int) ([]int, []int) {
	prices := make([]int, n)
	changes := make([]int, n)
	for i := 0; i < n; i++ {
		prices[i] = initial % 10
		initial = step(initial)
		if i == 0 {
			changes[i] = 0
		} else {
			changes[i] = prices[i] - prices[i-1]
		}
	}
	return prices, changes
}

func bestSequence(prices, changes []int) map[sequence]int {
	localMap := make(map[sequence]int)
	for i := 4; i < len(changes); i++ {
		seq := sequence{
			changes[i-3], 
			changes[i-2], 
			changes[i-1], 
			changes[i],
		}
		// the earliest occurrence of a sequence is the will be taken
		if _, exists := localMap[seq]; !exists {
			localMap[seq] = prices[i]
		}
	}

	return localMap
}

func part2(data []int) {
	globalMap := make(map[sequence]int)

	var wg sync.WaitGroup
	localMaps := make(chan map[sequence]int, len(data))
	
	for _, val := range data {
		wg.Add(1)
		go func(secret int) {
			defer wg.Done()
			prices, changes := generatePricesAndChanges(secret, 2000)
			localMaps <- bestSequence(prices, changes)
		}(val)
	}

	go func() {
		wg.Wait()
		close(localMaps)
	}()

	for m := range localMaps {
		for k, v := range m {
			globalMap[k] += v
		}
	}

	highest := 0
	var bestSeq sequence
	for k, sumVal := range globalMap {
		if sumVal > highest {
			highest = sumVal
			bestSeq = k
		}
	}
	fmt.Println("Part2:", highest, bestSeq)	
}

func main() {
	data := getData()
	part1(data)
	part2(data)
}