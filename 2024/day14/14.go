package main

import (
	"fmt"
	"os"
	"strings"
)
type Data struct {
	pos, vel [2]int 
}

func getData(path string) string {
    content, err := os.ReadFile(path)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return strings.TrimRight(string(content), "\n")
}

func parseData(data string) []Data {
	result := []Data{}
	lines := strings.Split(data, "\r\n")
	for _, line := range lines {
		var pos, vel [2]int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &pos[0], &pos[1], &vel[0], &vel[1])
		result = append(result, Data{pos, vel})
	}
	return result
}

func mod(a, b int) int {
    return (a % b + b) % b
}

func part1(data []Data) {
	N := 100
	//w, h := 11, 7
	w, h := 101, 103
	quadrant := [4]int{}
	for _, d := range data {
		x := mod(d.pos[0] + N*d.vel[0],w)
		y := mod(d.pos[1] + N*d.vel[1],h)
		if x < w/2 && y < h/2 {
			quadrant[0]++
		} 
		if x > w/2 && y < h/2 {
			quadrant[1]++
		}
		if x < w/2 && y > h/2 {
			quadrant[2]++
		}
		if x > w/2 && y > h/2 {
			quadrant[3]++
		}
	}
	sum := 1
	for _, q := range quadrant {
		sum *= q
	}
	fmt.Println(sum)
}

// funny that the final image was unique in points
func part2(data []Data) {
	//w, h := 11, 7
	w, h := 101, 103
	maxIter := 100000
	for i := 0; i < maxIter; i++ {
		points := make(map[[2]int]bool)
		duplicate := false
		for _, d := range data {
			x := mod(d.pos[0] + i*d.vel[0], w)
			y := mod(d.pos[1] + i*d.vel[1], h)
			point := [2]int{x, y}
			if _, ok := points[point]; ok {
				duplicate = true
				break
			}
			points[point] = true
		}
		
		if !duplicate {
			fmt.Println(i)
			showEgg(data, i)
			break
		}
	}
}

func showEgg(data []Data, i int) {
	w, h := 101, 103
	egg := make([][]bool, h)
	for i := range egg {
		egg[i] = make([]bool, w)
	}
	for _, d := range data {
		x := mod(d.pos[0] + i*d.vel[0], w)
		y := mod(d.pos[1] + i*d.vel[1], h)
		egg[y][x] = true
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if egg[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	data := getData("input.txt")
	cleanData := parseData(data)
	part1(cleanData)
	part2(cleanData)
}