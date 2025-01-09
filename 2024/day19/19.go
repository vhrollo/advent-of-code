// we be doing a little tabulation

package main

import (
	"fmt"
	"os"
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

func parseData(data string) ([]string, []string){
	parts := strings.Split(data, "\n\n")

	patterns := strings.Split(parts[0], ", ")
	designs := strings.Split(parts[1], "\n")

	return patterns, designs
}

func canFormDesign(design string, patterns []string) bool {
	n := len(design)
	dp := make([]bool, n+1)
	dp[n] = true
	for i := n-1; i >= 0; i-- {
		for _, pattern := range patterns {
			if i+len(pattern) <= n && design[i:i+len(pattern)] == pattern && dp[i+len(pattern)] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}

func part1(patterns []string, designs []string) {
	count := 0
	for _, design := range designs {
		if canFormDesign(design, patterns) {
			count++
		}
	}
	fmt.Println(count)
}

func canFormDesigns(design string, patterns []string) int {
	n := len(design)
	dp := make([]int, n+1)
	dp[n] = 1
	for i := n-1; i >= 0; i-- {
		for _, pattern := range patterns {
			if i+len(pattern) <= n && design[i:i+len(pattern)] == pattern {
				dp[i] += dp[i+len(pattern)]
			}
		}
	}
	return dp[0]
}

func part2(patterns []string, designs []string) {
	count := 0
	for _, design := range designs {
		count += canFormDesigns(design, patterns)
	}
	fmt.Println(count)
}
	

func main() {
	data := getData("input.txt")
	patterns, designs := parseData(data)
	part1(patterns, designs)
	part2(patterns, designs)
}