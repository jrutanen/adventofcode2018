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

func main() {
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
	//	sleeping := false
	layout := "2006-01-02 15:04"

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
	//find guard that has slept the most
	idSleepy := 0
	maxSleepTime := 0
	first := true
	for _, i := range g {
		if first {
			first = false
			maxSleepTime = i.totalSleepTime
		} else {
			if i.totalSleepTime > maxSleepTime {
				maxSleepTime = i.totalSleepTime
				idSleepy = i.id
			}
		}
	}
	fmt.Println(strconv.Itoa(idSleepy) + ": " + strconv.Itoa(g[idSleepy].totalSleepTime))

	//Check what times the sleepy one is most likely to be asleep
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
		for i := range table {
			fmt.Println(table[i])
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
	fmt.Println(frequency)
	//get index of the highest frequency
	iMax := 0
	for i := range frequency {
		if frequency[i] > frequency[iMax] {
			iMax = i
		}
	}
	result := idSleepy * iMax
	fmt.Println(strconv.Itoa(idSleepy) + " * " + strconv.Itoa(iMax) + " = " + strconv.Itoa(result))
}
