package main

import (
	"fmt"
	"os"
	"testing"
)

func TestOrder(t *testing.T) {
	expected := [6]rune{'C', 'A', 'B', 'D', 'F', 'E'}
	input, err := os.Open("input_day7_test.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer input.Close()

	createSliceOfSteps(input)

	runes := setOrder()

	for i, r := range runes {
		if r != expected[i] {
			t.Errorf("Wrong rune. Expected %U got %U\n", expected[i], r)
		}
	}

	if len(runes) != 6 {
		t.Errorf("Not all runes returned\n")
	}
}
