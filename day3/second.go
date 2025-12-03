package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv")

func check(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
		panic(e)
	}
}

func main() {
	// read in file input
	contents, err := os.Open("first_input.txt")
	check(err)
	var joltageSum int64 = 0
	
	scanner := bufio.NewScanner(contents)

	for scanner.Scan() {
		// read lines and split by commas
		line := scanner.Text()

		var largestJoltageString string = ""
		var joltageCountMax int = 12
		var startPoint int = 0
		
		// damn this one is hard
		// lets step through until we find the largest number we can with the remaining count still left
		for joltageCount := 0; joltageCount < joltageCountMax; joltageCount++ {
			var largestJoltageNow int64 = 0

			for i := startPoint; i <= len(line) - (joltageCountMax - joltageCount); i++ {
				joltageString := string(line[i:i+1])

				potentialJoltage, err := strconv.ParseInt(joltageString, 10, 64)
				check(err)

				if potentialJoltage > largestJoltageNow {
					largestJoltageNow = potentialJoltage
					startPoint = i + 1
				}
			}

			largestJoltageString += strconv.FormatInt(largestJoltageNow, 10)
		}

		largestJoltage, err := strconv.ParseInt(largestJoltageString, 10, 64)
		check(err)

		joltageSum += largestJoltage
	}

	fmt.Print("Joltage sum: ")
	fmt.Println(joltageSum)

}