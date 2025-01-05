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

func parseData(data string) {
	strings.Split(data, "\r\n\r\n")

}

func main() {
	data := getData("test.txt")
	fmt.Println(data)
}