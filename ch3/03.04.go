package main

import (
	"fmt"
)

func main() {
	str := "съешь еще этих мягких французских булок, да выпей чаю"
	res := CountLetters(str)
	for key, val := range res {
		fmt.Printf("%c - %d\n", key, val)
	}
}

func CountLetters(str string) map[rune]int {
	var res = make(map[rune]int)
	for _, ch := range str {
		res[ch]++
	}
	return res
}
