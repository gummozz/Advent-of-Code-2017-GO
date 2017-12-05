package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func SortThatShit(A string, N int) int {
	var sumShit int = 0
	a := strings.Split(A,"	")
	b := make([]int, N)
	for i, v := range a {
		b[i], _ = strconv.Atoi(v)
	}
	for i := range b {
		for j := range b {
			if b[i]%b[j] == 0 && b[i]/b[j] != 1 {
				sumShit = b[i] / b[j]
			}
		}
	}
	return sumShit
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	answer := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		answer += SortThatShit(currentLine, 16)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(answer)
}
