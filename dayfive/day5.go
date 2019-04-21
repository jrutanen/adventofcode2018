package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

func main() {
	//read input data from file
	input, err := ioutil.ReadFile("input_day5.txt")
	if err != nil {
		fmt.Print(err)
	}
	// convert byte slice to rune slice
	var polymers = []rune(string(input))
	var finished = false
	var counter = 0
	for !finished {
		for i, v := range polymers {
			if i >= len(polymers)-1 {
				finished = true
				break
			}
			if unicode.IsLower(v) {
				if unicode.IsUpper(polymers[i+1]) {
					vNext := unicode.ToLower(polymers[i+1])
					if v == vNext {
						polymers = append(polymers[:i], polymers[i+2:]...)
						counter = 0
						//slice is changed let's start from the start again
						break
					}
				}
			} else {
				if unicode.IsLower(polymers[i+1]) {
					vNext := unicode.ToUpper(polymers[i+1])
					if v == vNext {
						polymers = append(polymers[:i], polymers[i+2:]...)
						counter = 0
						//slice is changed let's start from the start again
						break
					}
				}
			}
		}
		//there were no more changes to the polymer
		if counter >= len(polymers)-1 {
			finished = true
		}
	}
	//-1 for len is for linefeed
	fmt.Printf("Result: %d", len(polymers)-1)
}
