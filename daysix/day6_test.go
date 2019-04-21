package main

import (
	"fmt"
	"os"
	"testing"
)

func TestArea(t *testing.T) {
	expected := [6]int{0, 0, 0, 9, 17, 0}
	input, err := os.Open("input_day6_test.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer input.Close()

	coordinates := createSliceOfCoordinates(input)
	calculateArea(coordinates)

	for i, coord := range coordinates {
		if coord.area != expected[i] {
			t.Errorf("Wrong area. Expected %d got %d", expected[i], coord.area)
		}
	}
}

func TestDistances(t *testing.T) {
	expected := 16
	input, err := os.Open("input_day6_test.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer input.Close()

	coordinates := createSliceOfCoordinates(input)

	area := calculateDistances(coordinates, 32)
	if area != expected {
		t.Errorf("Wrong area. Expected %d got %d", expected, area)
	}
}
