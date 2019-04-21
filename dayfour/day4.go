package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type sleeping struct {
	start time.Time
	stop  time.Time
}

type guard struct {
	id             int
	totalSleepTime int
	sleeping       []sleeping
}

type guards map[int]*guard

func highestFrequency(g guards, idSleepy int) (int, int) {
	layout := "2006-01-02 15:04"
	//Check what times the guard is most likely to be asleep
	var table map[string][60]int
	table = make(map[string][60]int)
	for i := range g[idSleepy].sleeping {
		var mins [60]int
		startTime := (g[idSleepy].sleeping[i].start.Format(layout))
		endTime := (g[idSleepy].sleeping[i].stop.Format(layout))
		end := strings.Index(startTime, " ")
		date := startTime[:end]
		start := strings.Index(startTime, ":") + 1
		startMinute, _ := strconv.Atoi(startTime[start:])
		start = strings.Index(endTime, ":") + 1
		endMinute, _ := strconv.Atoi(endTime[start:])

		//check if date is already added to the map
		_, ok := table[date]
		if !ok {
			for i = startMinute; i < endMinute; i++ {
				mins[i] = 1
			}
			//not in the map, add to map and set number of occurences to 1
			table[date] = mins
		} else {
			mins = table[date]
			for i = startMinute; i < endMinute; i++ {
				mins[i] = 1
			}
		}
	}
	//Find most frequent minute
	var frequency [60]int
	for i := range table {
		for j := range table[i] {
			if table[i][j] == 1 {
				frequency[j]++
			}
		}
	}
	//get index of the highest frequency
	iMax := 0
	for i := range frequency {
		if frequency[i] > frequency[iMax] {
			iMax = i
		}
	}
	return iMax, frequency[iMax]
}

func main() {
	layout := "2006-01-02 15:04"
	//read input data from file
	shiftFile, err := os.Open("input_day4.txt")
	fmt.Print(err)
	defer shiftFile.Close()

	//copy rows to slice
	shifts := []string{}
	scanner := bufio.NewScanner(shiftFile)
	for scanner.Scan() {
		shifts = append(shifts, scanner.Text())
	}

	//sort items in slice
	sort.Strings(shifts)

	//go through the list and create guards and add them to a map
	//inialize a map
	g := make(guards)
	guardId := 0
	newShift := true
	var sleepStart time.Time
	var sleepStop time.Time
	var sleepTime time.Time

	//loop trough rows and add data to each guard and add it to the map
	for _, row := range shifts {
		//New shift
		if strings.Contains(row, "Guard") {
			start := strings.Index(row, "#") + 1
			end := strings.Index(row, " begins shift")
			guardId, _ = strconv.Atoi(string(row[start:end]))

			//store information from previous shift
			if newShift {
				//check if guard is already added to the map
				_, ok := g[guardId]
				if !ok {
					//not in the map, add to map and set number of occurences to 1
					g[guardId] = &guard{id: guardId}
				}
			}
			newShift = true
		}

		if strings.Contains(row, "falls asleep") {
			start := strings.Index(row, "[") + 1
			end := strings.Index(row, "]")
			timeString := string(row[start:end])
			sleepStart, _ = time.Parse(layout, timeString)
		}

		if strings.Contains(row, "wakes up") {
			start := strings.Index(row, "[") + 1
			end := strings.Index(row, "]")
			timeString := string(row[start:end])
			sleepStop, _ = time.Parse(layout, timeString)
			sleepTime = sleepTime.Add(sleepStop.Sub(sleepStart))
			g[guardId].sleeping = append(g[guardId].sleeping,
				sleeping{start: sleepStart, stop: sleepStop})
			g[guardId].totalSleepTime = g[guardId].totalSleepTime + int(sleepStop.Sub(sleepStart).Minutes())
		}
	}
	//Strategy 1: Find the guard that has the most minutes asleep.
	//What minute does that guard spend asleep the most?
	//Find guard that has slept the most
	idSleepy := 0
	maxSleepTime := 0
	for _, i := range g {
		if i.totalSleepTime > maxSleepTime {
			maxSleepTime = i.totalSleepTime
			idSleepy = i.id
		}
	}
	//The minute guard sleeps most
	iMax, _ := highestFrequency(g, idSleepy)
	//result
	result := idSleepy * iMax
	fmt.Println("Part one: " + strconv.Itoa(idSleepy) + " * " + strconv.Itoa(iMax) + " = " + strconv.Itoa(result))

	//Strategy 2: Of all guards, which guard is most frequently asleep on the same minute?
	record := 0
	minute := 0
	var guard int

	for _, i := range g {
		min, freq := highestFrequency(g, i.id)
		if freq > record {
			record = freq
			guard = i.id
			minute = min
		}
	}
	fmt.Printf("Guard %d slept %d times during minute %d\n", guard, record, minute)
	result = guard * minute
	fmt.Printf("Result: %d\n", result)
}
