package main

import (
    "fmt"
    "os"
    "strings"
	"strconv"
)

type fileInfo struct {
    id     int
    length int
    start  int // leftmost block index of this file on the disk
    blocks []int
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

func parseDiskMap(s string) []int {
    var blocks []int
    fileID := 0

    for i := 0; i < len(s); i += 2 {
        fLen, _ := strconv.Atoi(string(s[i]))
        // Append fLen blocks for the current file ID
        for j := 0; j < fLen; j++ {
            blocks = append(blocks, fileID)
        }
        fileID++

        if i+1 < len(s) {
            sLen, _ := strconv.Atoi(string(s[i+1]))
            for j := 0; j < sLen; j++ {
                blocks = append(blocks, -1)
            }
        }
    }
    return blocks
}

func defrag(blocks []int) {
    left := 0
    right := len(blocks) - 1

    for left < right {
        // Advance 'left' until we find a free space
        for left < right && blocks[left] != -1 {
            left++
        }
        // Advance 'right' backward until we find a file block
        for right > left && blocks[right] == -1 {
            right--
        }
        if left < right {
            // Move block from 'right' to 'left'
            blocks[left], blocks[right] = blocks[right], -1
            left++
            right--
        }
    }
}

func part1(data string) {
	clean_data := parseDiskMap(data)
	defrag(clean_data)
	fmt.Println(computeChecksum(clean_data))	
}

// part 2
func collectFiles(blocks []int) []*fileInfo {
    fileMap := make(map[int]*fileInfo)

    for i, id := range blocks {
        if id >= 0 { // not free
            fi, ok := fileMap[id]
            if !ok {
                // first time we see this file ID
                fi = &fileInfo{id: id, start: i}
                fileMap[id] = fi
            }
            fi.blocks = append(fi.blocks, i)
            if i < fi.start {
                fi.start = i
            }
        }

    }

    files := make([]*fileInfo, 0, len(fileMap))
    for _, fi := range fileMap {
        fi.length = len(fi.blocks)
        files = append(files, fi)
    }
	
	// dont care to make a better sort
	for i := 0; i < len(files); i++ {
		for j := i + 1; j < len(files); j++ {
			if files[i].start > files[j].start {
				files[i], files[j] = files[j], files[i]
			}
		}
	}

    return files
}

func findFirstfit(blocks []int, fi *fileInfo) int {

	for i := 0; i <= fi.start-fi.length; i++ {
		if blocks[i] == -1 {
			j := 1
			for ; j < fi.length; j++ {
				if blocks[i+j] != -1 {
					break
				}
			}
			if j == fi.length {
				return i
			}
		}
	}
	return -1
}

func moveFile(blocks []int, fi *fileInfo, dest int) {
	for i, j := fi.start, dest; i < fi.start+fi.length; i, j = i+1, j+1 {
		blocks[j] = blocks[i]
		blocks[i] = -1
	}

	newBlocks := make([]int, fi.length)
    for i := 0; i < fi.length; i++ {
        newBlocks[i] = dest + i
    }
    fi.blocks = newBlocks
    fi.start = dest
}

func part2(data string) {
	clean_data := parseDiskMap(data)
	files := collectFiles(clean_data)
	// iterate from last file to first
	for i := len(files) - 1; i >= 0; i-- {
		fi := files[i]
		dest := findFirstfit(clean_data, fi)
		if dest != -1 {
			moveFile(clean_data, fi, dest)
		}
		// fmt.Println(clean_data)
	}
	fmt.Println(computeChecksum(clean_data))
	//fmt.Println(clean_data)
}

func computeChecksum(blocks []int) int64 {
    var sum int64
    for i, id := range blocks {
        if id != -1 {
            sum += int64(i) * int64(id)
        }
    }
    return sum
}

func main() {
	data := get_data("input.txt")
	part1(data)
	part2(data)
}