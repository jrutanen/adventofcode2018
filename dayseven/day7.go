package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type step struct {
	id  rune
	pre []rune
}

var steps []*step

func setOrder() []rune {
	var order []rune
	var firstRune rune
	//get first rune (no previous rune available)
	for _, s := range steps {
		if len(s.pre) == 0 {
			firstRune = s.id
			order = append(order, firstRune)
			break
		}
	}

	//get all steps that have first rune as previous step
	var nextSteps []rune
	for _, s := range steps {
		if stepInPrevious(s, getStep(firstRune)) {
			nextSteps = append(nextSteps, s.id)
		}
	}
	for _, r := range nextSteps {
		order = append(order, r)
	}

	return order
}

func stepInPrevious(stepID *step, toFind *step) bool {
	for _, s := range steps {
		if s.id == stepID.id {
			for _, p := range s.pre {
				if p == toFind.id {
					return true
				}
			}
		}
	}
	return false
}

func stepInSteps(stepID rune) bool {
	for _, s := range steps {
		if s.id == stepID {
			return true
		}
	}
	return false
}

func getStep(stepID rune) *step {
	for _, s := range steps {
		if s.id == stepID {
			return s
		}
	}
	return nil
}

func createStep(stepID rune, previous rune) {
	oldStep := stepInSteps(stepID)
	oldPrevious := stepInSteps(previous)

	if !oldStep {
		newStep := step{id: stepID}
		newStep.pre = append(newStep.pre, previous)
		steps = append(steps, &newStep)
	} else {
		s := getStep(stepID)
		s.pre = append(s.pre, previous)
	}
	if !oldPrevious {
		newStep := step{id: previous}
		steps = append(steps, &newStep)
	}
}

func createSliceOfSteps(file *os.File) []*step {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		start := strings.Index(scanner.Text(), "step ") + 5
		end := strings.Index(scanner.Text(), " can")
		stepRune := []rune(scanner.Text()[start:end])
		end = strings.Index(scanner.Text(), " must")
		previousRune := []rune(scanner.Text()[5:end])
		stepID := stepRune[0]
		previous := previousRune[0]
		createStep(stepID, previous)
	}
	return steps
}

func main() {
	//read input data from file
	input, err := os.Open("input_day7.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer input.Close()

	//copy coordinates to struct
	createSliceOfSteps(input)

	//Part one
	//What is the size of the largest area that isn't infinite?

	fmt.Printf("Part One result: %d\n", 1)

	//Part two
}
