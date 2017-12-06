package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
)

func main() {
	//var answer int = 0
	i := 0
	jumps := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	intArray := make([]int, 0)
	for scanner.Scan() {
		a := scanner.Text()
		currentLine, _ := strconv.Atoi(a)
		intArray = append(intArray, currentLine)
		i++
	}
	i = 0
	for i >= 0 && i < len(intArray){
		fmt.Println("Since we're standing on ", intArray[i], " we should jump ", intArray[i], "steps, landing on ", intArray[i + intArray[i]], " then increase ", intArray[i]," with 1, making it",intArray[i] +1)
		jumps++
		fmt.Println(". We have jumped ",jumps," times.")
		intArray[i]++
		i += intArray[i]-1


	}

}
