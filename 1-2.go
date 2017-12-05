package main

import (
	"fmt"
	"container/ring"
	"strconv"
)

var input string = ""
var n int = len(input) /2
var answer int = 0

func main() {
	runes := []rune(input)
	r := ring.New(len(input))
	for i := 0; i < r.Len(); i++ {
		r.Value = string(runes[i])
		r = r.Next()
	}
	r = r.Move(n)
	for i := 0; i < len(input); i++ {
		if string(runes[i]) == r.value(string) {
			convertedRune, _ := strconv.Atoi(string(runes[i]))
			answer += convertedRune
		}
		r = r.Next()
	}
	fmt.Println(answer)
}
