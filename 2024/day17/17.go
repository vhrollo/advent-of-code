package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func parseData(data string) ([]int, []int) {
	parts := strings.Split(data, "\n\n")

	var registerValues []int
	for _, line := range strings.Split(parts[0], "\n") {
		parts := strings.Split(line, ": ")
		value, _ := strconv.Atoi(parts[1])
		registerValues = append(registerValues, value)
	}

	programParts := strings.Split(parts[1], ": ")
	var programOpcodes []int
	for _, line := range strings.Split(programParts[1], ",") {
		value, _ := strconv.Atoi(line)
		programOpcodes = append(programOpcodes, value)
	}

	return registerValues, programOpcodes
}

// part 1

func getOperand(combo int, registerValues *[]int) int {
	if 0 <= combo && combo <= 3 {
		return combo
	} else if 4 <= combo && combo <= 6 {
		return (*registerValues)[combo - 4]
	} else if combo == 7 {
		fmt.Println("This program is not valid")
		os.Exit(1)
	} else {
		os.Exit(1)
	}
	return 0
}

//op 0 devision, saves to A
func adv(a int, registerValues *[]int) {
	combo := getOperand(a, registerValues)
	(*registerValues)[0] = int(float64((*registerValues)[0]) / math.Pow(2, float64(combo)))
}

//op 1 xor, saves to B
func bxl(a int, registerValues *[]int) {
	(*registerValues)[1] = (*registerValues)[1] ^ a
}

//op 2 modulo, saves to B
func bst (a int, registerValues *[]int) {
	combo := getOperand(a, registerValues)
	(*registerValues)[1] = combo % 8
}

//op 3 - jump
func jnz (a int, registerValues *[]int, ptr *int) {
	if (*registerValues)[0] != 0 {
		*ptr = a
	} else {
		*ptr = *ptr + 2
	}
}

//op 4 - xor
func bxc (registerValues *[]int) {
	(*registerValues)[1] = (*registerValues)[1] ^ (*registerValues)[2]
}

//op 5 - out to stdstr
func out (a int, registerValues *[]int, stdstr *[]int) {
	combo := getOperand(a, registerValues)
	print := combo % 8
	*stdstr = append(*stdstr, print)
}

//op 6 adv but to B
func bdv (a int, registerValues *[]int) {
	combo := getOperand(a, registerValues)
	(*registerValues)[1] = int(float64((*registerValues)[0]) / math.Pow(2, float64(combo)))
}

//op 7 adv but to C
func cdv (a int, registerValues *[]int) {
	combo := getOperand(a, registerValues)
	(*registerValues)[2] = int(float64((*registerValues)[0]) / math.Pow(2, float64(combo)))
}


func part1(registerValues []int, programOpcodes []int) {
	ptr := 0
	var stdstr []int
	for ptr < len(programOpcodes) - 1 {
		opcode := programOpcodes[ptr]
		combo := programOpcodes[ptr + 1]
		switch opcode {
		case 0:
			adv(combo, &registerValues)
		case 1:
			bxl(combo, &registerValues)
		case 2:
			bst(combo, &registerValues)
		case 3:
			jnz(combo, &registerValues, &ptr)
			ptr -= 2
		case 4:
			bxc(&registerValues)
		case 5:
			out(combo, &registerValues, &stdstr)
		case 6:
			bdv(combo, &registerValues)
		case 7:
			cdv(combo, &registerValues)
		default:
			fmt.Println("This program is not valid")
		}
		ptr += 2
	}
	fmt.Println("Part 1")
	fmt.Println("Register values:", registerValues)
	str := ""
	for i, s := range stdstr {
		if i == 0 {
			str += strconv.Itoa(s)
		} else {
			str += "," + strconv.Itoa(s)
		}
	}
	fmt.Println("Output:", str)	
}


// part 2
func part2(originalRegisters []int, programOpcodes []int) {
    fmt.Println("Part 2")

    // We'll start from A = 1 if "lowest positive" is required.
    for A := 1; A < 999999999; A++ {
        // Make a fresh copy of the registers.
        copyOfRegisters := make([]int, len(originalRegisters))
        copy(copyOfRegisters, originalRegisters)

        // Set register A in this fresh copy
        copyOfRegisters[0] = A

        // Run helper
        if helper(copyOfRegisters, programOpcodes) {
            // If it matched, print and break
            fmt.Println(A)
            return
        }
    }
    fmt.Println("No solution found up to a very large A...")
}


func helper(registerValues []int, programOpcodes []int) bool {
	ptr := 0
	compareptr := 0
	var stdstr []int
	for ptr < len(programOpcodes) - 1 {
		opcode := programOpcodes[ptr]
		combo := programOpcodes[ptr + 1]
		switch opcode {
		case 0:
			adv(combo, &registerValues)
		case 1:
			bxl(combo, &registerValues)
		case 2:
			bst(combo, &registerValues)
		case 3:
			jnz(combo, &registerValues, &ptr)
			ptr -= 2
		case 4:
			bxc(&registerValues)
		
		case 5: // out
			out(combo, &registerValues, &stdstr)

			// If we've already output as many values as the program has instructions,
			// it's probably time to decide success/failure right away.
			if compareptr >= len(programOpcodes) {
				return false // We produced more output than we have instructions to match!
			}

			if stdstr[compareptr] != programOpcodes[compareptr] {
				// Mismatch => immediately return false
				return false
			}

			compareptr++
			// If we have matched all opcodes exactly, success!
			if compareptr == len(programOpcodes) {
				return true
			}
		case 6:
			bdv(combo, &registerValues)
		case 7:
			cdv(combo, &registerValues)
		default:
			fmt.Println("This program is not valid")
		}
		ptr += 2
	}
	return false
}


func main() {
	data := getData("input.txt")
	registerValues, programOpcodes := parseData(data)
	part1(registerValues, programOpcodes)
	part2(registerValues, programOpcodes)
}