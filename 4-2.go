package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"sort"
)

func solveShit(A string) (int) {
	splittedShit := strings.Split(A, " ")
	shit := 0
	nope := false
	for i := range splittedShit {
		if !nope {
			tempSplit := strings.Split(splittedShit[i], "")
			sort.Strings(tempSplit)
			joinedTempSplit := strings.Join(tempSplit, "")
			for j := range splittedShit {
				tempSplitJ := strings.Split(splittedShit[j], "")
				sort.Strings(tempSplitJ)
				joinedTempSplitJ := strings.Join(tempSplitJ, "")
				if j != i && joinedTempSplit == joinedTempSplitJ {
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