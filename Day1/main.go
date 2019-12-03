package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main_part_1() {
	// open the file and defer its closure
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Failed to open the file...", err)
	}
	defer func() {
		if err = inputFile.Close(); err != nil {
			log.Fatal("Failed to close file...", err)
		}
	}()

	// iterate over each line (we roughly know there are 100 lines, so we can
	// limit the size of the array we want it to grow to)
	intSlice := make([]int, 0, 100)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		// grab the line, trim it, cast it to an int, fatal error if problems
		line := scanner.Text()
		v, err := strconv.Atoi(strings.Trim(line, "\r\n"))
		if err != nil {
			log.Fatal("Ill formed integer string: ", err)
		}
		// add the new int to the slice of integers
		intSlice = append(intSlice, v)
	}

	// calculate a running sum over the masses
	sum := 0
	for _, v := range intSlice {
		sum += calculateFuelRequirement(v)
	}

	log.Println("Sum of fuel of all components: ", sum)
}

func main() {
	// open the file and defer its closure
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Failed to open the file...", err)
	}
	defer func() {
		if err = inputFile.Close(); err != nil {
			log.Fatal("Failed to close file...", err)
		}
	}()

	// iterate over each line (we roughly know there are 100 lines, so we can
	// limit the size of the array we want it to grow to)
	intSlice := make([]int, 0, 100)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		// grab the line, trim it, cast it to an int, fatal error if problems
		line := scanner.Text()
		v, err := strconv.Atoi(strings.Trim(line, "\r\n"))
		if err != nil {
			log.Fatal("Ill formed integer string: ", err)
		}
		// add the new int to the slice of integers
		intSlice = append(intSlice, v)
	}

	// calculate a running sum over the masses
	sum := 0
	for _, v := range intSlice {
		sum += fullyCalculateFuel(v)
	}

	log.Println("Sum of fuel of all components: ", sum)
}

func calculateFuelRequirement(mass int) int {
	needed := int(math.Floor(float64(mass/3))) - 2
	if needed < 0 {
		return 0
	}
	return needed
}

func recursivelyCalculateFuel(mass int) int {
	// calculate the needed fuel
	fuelNeeded := calculateFuelRequirement(mass)
	// if we've reached a sufficiently small amount...
	if fuelNeeded == 0 {
		return 0
	}
	// otherwise, return the needed fuel plus whatever
	// else is needed for the fuel we added
	return fuelNeeded + recursivelyCalculateFuel(fuelNeeded)
}

func fullyCalculateFuel(mass int) int {
	currentFuelNeeds := calculateFuelRequirement(mass)
	runningSum := currentFuelNeeds
	for currentFuelNeeds > 0 {
		currentFuelNeeds = calculateFuelRequirement(currentFuelNeeds)
		runningSum += currentFuelNeeds
	}
	return runningSum
}
