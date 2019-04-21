package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

func collapsePolymer(poly []rune) int {
	var finished = false
	for !finished {
		for i, v := range poly {
			if i >= len(poly)-1 {
				finished = true
				break
			}
			if unicode.IsLower(v) {
				if unicode.IsUpper(poly[i+1]) {
					vNext := unicode.ToLower(poly[i+1])
					if v == vNext {
						poly = append(poly[:i], poly[i+2:]...)
						//slice is changed let's start from the start again
						break
					}
				}
			} else {
				if unicode.IsLower(poly[i+1]) {
					vNext := unicode.ToUpper(poly[i+1])
					if v == vNext {
						poly = append(poly[:i], poly[i+2:]...)
						//slice is changed let's start from the start again
						break
					}
				}
			}
		}
	}
	//-1 for len is for linefeed
	result := len(poly) - 1
	return result
}

func removeUnit(letter rune, runes []rune) []rune {
	var result []rune
	for _, c := range runes {
		if c != letter && c != unicode.ToUpper(letter) {
			result = append(result, c)
		}
	}
	return result
}

func main() {
	//read input data from file
	input, err := ioutil.ReadFile("input_day5.txt")
	if err != nil {
		fmt.Print(err)
	}
	// convert byte slice to rune slice
	var polymers = []rune(string(input))

	//Part One
	fmt.Printf("Part One Result: %d\n", collapsePolymer(polymers))

	//Part Two
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var bestLetter rune
	shortest := 70000
	for _, letter := range alphabet {
		newPolymers := removeUnit(letter, []rune(string(input)))
		count := collapsePolymer(newPolymers)
		if count < shortest {
			shortest = count
			bestLetter = letter
		}
	}
	fmt.Printf("Part Two Result: %c: %d\n", bestLetter, shortest)
}
