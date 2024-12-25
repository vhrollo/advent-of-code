package main

import (
    "fmt"
    "os"
    "strings"
	"container/list"
	"strconv"
	"math"
)

func getData(path string) string {
    content, err := os.ReadFile(path)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return strings.TrimRight(string(content), "\n")
}

func makeLinkedList(dataStr string) *list.List {
	data := strings.Split(dataStr, " ")
	l := list.New()
	for _, v := range data {
		num, _ := strconv.Atoi(v)
		l.PushBack(num)
	}
	return l
}

func isEven(n int) bool { // length of the number
	return len(strconv.Itoa(n)) % 2 == 0
}

// could be optimized
func splitInHalf(n int) (int, int) {
	s := strconv.Itoa(n)
	half := len(s) / 2
	firstHalf, _ := strconv.Atoi(s[:half])
	secondHalf, _ := strconv.Atoi(s[half:])
	return firstHalf, secondHalf
}

func part1(dataList *list.List) {
	const N int = 25
	const magic int = 2024

	for i := 0; i < N; i++ {
		for e := dataList.Front(); e != nil; {
			currentVal := e.Value.(int)

			nextElem := e.Next()

			if currentVal == 0 {
				e.Value = 1
			} else if isEven(currentVal) {
				firstHalf, secondHalf := splitInHalf(currentVal)
				e.Value = firstHalf
				// Insert the second half after 'e'
				dataList.InsertAfter(secondHalf, e)
				// We do *not* advance 'e' to the newly inserted node; we skip it
			} else {
				e.Value = currentVal * magic
			}

			// Move on to nextElem, *not* the newly inserted element
			e = nextElem
		}
	}
	
	// for e := dataList.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value)
	// }

	fmt.Println(dataList.Len())
}

////////////////////////////

var memo = make(map[string]int)

// to check how effective memoization is
var count int = 0
var total int = 0
var saved float64 = 0


func multiplyBy2024(s string) string {
	n, _ := strconv.Atoi(s)
	return strconv.Itoa(n * 2024)
}

func splitInHalfString(s string) (string, string) {
    half := len(s) / 2
    left := s[:half]
    right := s[half:]

    left = strings.TrimLeft(left, "0")
    if left == "" {
        left = "0"
    }

    right = strings.TrimLeft(right, "0")
    if right == "" {
        right = "0"
    }

    return left, right
}

func countStones(s string, Blinks int) int {
	if Blinks == 0 {
		return 1
	}
	
	key := s + "#" + strconv.Itoa(Blinks)
	if val, ok := memo[key]; ok {
		count++
		saved += math.Pow(1.5, float64(Blinks))
		return val
	}
	total++

	var result int

	if s == "0" {
		result = countStones("1", Blinks - 1)
	} else if len(s) % 2 == 0 {
		left, right := splitInHalfString(s)
		result = countStones(left, Blinks - 1) + countStones(right, Blinks - 1)
	} else {
		multiplied := multiplyBy2024(s)
		result = countStones(multiplied, Blinks - 1)
	}

	memo[key] = result
	return result
}


// this has to be rewritten as the other has run time of O(2^n), ouch!
func part2(data string, Blinks int) {
	dataList := strings.Split(data, " ")	

	sum := 0
	for _, v := range dataList {
		sum += countStones(v, Blinks)
	}

	fmt.Println(sum)
}

func main () {
	dataStr := getData("input.txt")
	dataList := makeLinkedList(dataStr)
	for e := dataList.Front(); e != nil; e = e.Next() {
		//fmt.Println(e.Value)
	}
	part1(dataList)
	part2(dataStr, 75)
	
	a := float32(count) / float32(total)
	fmt.Println("Memoization saved", a * 100, "% of the calculations")
	fmt.Println("Total calculations:", total)
	fmt.Println("How many saved due to memoization:", saved)
}