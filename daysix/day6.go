package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x    int
	y    int
	area int
}

// Abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findBorders(coords []*coordinate) (int, int, int, int) {
	xMin := -1
	xMax := -1
	yMin := -1
	yMax := -1
	for _, coord := range coords {
		if coord.x > xMax {
			xMax = coord.x
		}
		if coord.x < xMin || xMin == -1 {
			xMin = coord.x
		}
		if coord.y > yMax {
			yMax = coord.y
		}
		if coord.y < yMin || yMin == -1 {
			yMin = coord.y
		}
	}
	return xMin, xMax, yMin, yMax
}

func getDistance(from *coordinate, to *coordinate) int {
	return abs(from.x-to.x) + abs(from.y-to.y)
}

func calculateArea(coords []*coordinate) {
	xMin, xMax, yMin, yMax := findBorders(coords)
	//area 0 means infinity
	for ci, coord := range coords {
		if coord.x <= xMin || coord.x >= xMax {
			coord.area = 0
		} else if coord.y <= yMin || coord.y >= yMax {
			coord.area = 0
		} else {
			area := 0
			for i := xMin; i < xMax; i++ {
				for j := yMin; j < yMax; j++ {
					currentPos := createCoordinate(i, j)
					distanceToMe := getDistance(coord, currentPos)
					distanceToOther := -1
					for cj, coordJ := range coords {
						//ignore self
						if ci != cj {
							distance := getDistance(coordJ, currentPos)
							if distance < distanceToOther || distanceToOther == -1 {
								distanceToOther = distance
							}
						}
					}
					if distanceToMe < distanceToOther {
						area++
					}
				}
			}
			coord.area = area
		}
	}
}

func calculateAreaTwo(coords []*coordinate) {
	xMin, xMax, yMin, yMax := findBorders(coords)
	//area 0 means infinity
	for ci, coord := range coords {
		if coord.x <= xMin || coord.x >= xMax {
			coord.area = 0
		} else if coord.y <= yMin || coord.y >= yMax {
			coord.area = 0
		} else {
			area := 0
			for i := xMin - 100; i < xMax+100; i++ {
				for j := yMin - 100; j < yMax+100; j++ {
					currentPos := createCoordinate(i, j)
					distanceToMe := getDistance(coord, currentPos)
					distanceToOther := -1
					for cj, coordJ := range coords {
						//ignore self
						if ci != cj {
							distance := getDistance(coordJ, currentPos)
							if distance < distanceToOther || distanceToOther == -1 {
								distanceToOther = distance
							}
						}
					}
					if distanceToMe < distanceToOther {
						area++
					}
				}
			}
			if coord.area < area {
				coord.area = 0
			}
		}
	}
}

func calculateDistances(coords []*coordinate, maxDistance int) int {
	xMin, xMax, yMin, yMax := findBorders(coords)
	//area 0 means infinity
	var okLocations []coordinate

	for i := xMin; i < xMax; i++ {
		for j := yMin; j < yMax; j++ {
			currentPos := createCoordinate(i, j)
			totalDistance := 0
			for _, coord := range coords {
				distance := getDistance(coord, currentPos)
				totalDistance = totalDistance + distance
			}
			if totalDistance < maxDistance {
				okLocations = append(okLocations, *currentPos)
			}
		}
	}
	return len(okLocations)
}

func createCoordinate(x int, y int) *coordinate {
	coord := new(coordinate)
	coord.x = x
	coord.y = y
	return coord
}

func createSliceOfCoordinates(file *os.File) []*coordinate {
	coords := []*coordinate{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		comma := strings.Index(scanner.Text(), ",")
		x, _ := strconv.Atoi(scanner.Text()[:comma])
		y, _ := strconv.Atoi(scanner.Text()[comma+2:])
		coords = append(coords, createCoordinate(x, y))
	}
	return coords
}

func main() {
	//read input data from file
	input, err := os.Open("input_day6.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer input.Close()

	//copy coordinates to struct
	coordinates := createSliceOfCoordinates(input)

	//Part one
	//What is the size of the largest area that isn't infinite?
	calculateArea(coordinates)
	//run through second time with bigger total area to see what areas are still growing (infinity)
	calculateAreaTwo(coordinates)

	//print the biggest area
	maxArea := 0
	for _, coord := range coordinates {
		if coord.area > maxArea {
			maxArea = coord.area
		}
	}
	fmt.Printf("Part One result: %d\n", maxArea)

	//Part two
	//What is the size of the region containing all locations which have a total
	//distance to all given coordinates of less than 10000?
	fmt.Printf("Part Two result: %d", calculateDistances(coordinates, 10000))
}
