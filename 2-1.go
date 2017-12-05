package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sort"
	"strconv"
)

func SortThatShit(A string, N string) int {
	a := strings.Split(A,"	")
	n, _ := strconv.Atoi(N)
	b := make([]int, n)
	for i, v := range a {
		b[i], _ = strconv.Atoi(v)
	}
	sort.Ints(b)
	c := make([]string, len(b))
	for i, y := range b {
		c[i] = strconv.Itoa(y)
	}
	first, err := strconv.Atoi(c[0])
	if err != nil {
		//lol error handling
	}
	last, err := strconv.Atoi(c[len(c)-1])
	if err != nil {
		//lol error handling
	}
	sumShit := last - first
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
		answer += SortThatShit(currentLine, "16")

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(answer)
}
