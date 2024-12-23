package main

import (
    "fmt"
    "os"
    "strings"
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


//will give the pos of every char of type x in a string
func parse_data(data string) (map[string][][2]int, int, int) {
	var pos_map = make(map[string][][2]int)

	lines := strings.Split(data, "\r\n")
	
	maxX := len(lines[0])
	maxY := len(lines)

	for j, line := range lines {
		for i, char := range line {
			if char == '.' {continue}

			if _, ok := pos_map[string(char)]; !ok {
				pos_map[string(char)] = [][2]int{}
			}
			pos_map[string(char)] = append(pos_map[string(char)], [2]int{i, j})
		}
	}
	return pos_map, maxX, maxY
}

//get pairs of list in k size in go
// takes pos v - pos u = new vector and return v + new vector of size k
// checks if the new pos is within the boundaries
//and does the same for u - new vector 
func get_new_pos(pos1 [2]int, pos2 [2]int, maxX int, maxY int) [][2]int {
	diff := [2]int{pos2[0] - pos1[0], pos2[1] - pos1[1]}
	new_pos := [][2]int{
		{pos2[0] + diff[0], pos2[1] + diff[1]},
		{pos1[0] - diff[0], pos1[1] - diff[1]},
	}

	checked_pos := [][2]int{}
	for _, pos := range new_pos {
		if pos[0] >= 0 && pos[0] < maxX && pos[1] >= 0 && pos[1] < maxY {
			checked_pos = append(checked_pos, pos)
		}
	}
	return checked_pos
}

// will check every interval in the line and return the new pos
func get_new_pos_freq(pos1 [2]int, pos2 [2]int, maxX int, maxY int) [][2]int {
	diff := [2]int{pos2[0] - pos1[0], pos2[1] - pos1[1]}
	new_pos := [][2]int{}
	// check both directions, and check out of bounds for each
	for i := 1; i < maxX; i++ {
		new_pos = append(new_pos, [2]int{pos2[0] + diff[0] * i, pos2[1] + diff[1] * i})
		new_pos = append(new_pos, [2]int{pos1[0] - diff[0] * i, pos1[1] - diff[1] * i})
	}

	checked_pos := [][2]int{}
	for _, pos := range new_pos {
		if pos[0] >= 0 && pos[0] < maxX && pos[1] >= 0 && pos[1] < maxY {
			checked_pos = append(checked_pos, pos)
		}
	}

	return checked_pos
}

func print_hashmap(hashmap map[string][][2]int) {
	for key, value := range hashmap {
		fmt.Println(key, len(value))
	}
}

func part1 (data string) {
	pos_map, maxX, maxY := parse_data(data)
	fmt.Println(maxX, maxY)
	//print_hashmap(pos_map)

	set := make(map[[2]int]bool)
	for _, value := range pos_map {
		for i, pos1 := range value {
			for _, pos2 := range value[i+1:] {
				new_pos := get_new_pos(pos1, pos2, maxX, maxY)
				for _, pos := range new_pos {
					set[pos] = true
				}
			}
		}
	}
	fmt.Println(len(set))
}

func part2 (data string) {
	pos_map, maxX, maxY := parse_data(data)
	fmt.Println(maxX, maxY)
	// print_hashmap(pos_map)

	set := make(map[[2]int]bool)
	for _, value := range pos_map {
		for i, pos1 := range value {
			for _, pos2 := range value[i+1:] {
				new_pos := get_new_pos_freq(pos1, pos2, maxX, maxY)
				for _, pos := range new_pos {
					set[pos] = true
				}
			}
		}
	}

	// also add the pos of antennas
	for _, value := range pos_map {
		for _, pos := range value {
			set[pos] = true
		}
	}

	fmt.Println(len(set))
}
	

func main() {
	data := get_data("input.txt")
	part2(data)
}