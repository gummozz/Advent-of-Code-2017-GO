package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)

func solveShit(A string) (int) {
	nope := false
	shit := 0
	splittedShit := strings.Split(A, " ")
	for i := range splittedShit {
		if !nope {
			for j := range splittedShit {
				if j != i && splittedShit[i] == splittedShit[j] {
					nope = true
				}
			}
		}
	}
	if !nope {
		shit = 1
	}
	return shit
}

func main() {
	var answer int = 0
	fmt.Println("hej")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		answer += solveShit(currentLine)
	}
	fmt.Println(answer)
}
