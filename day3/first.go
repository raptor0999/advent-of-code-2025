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
		var largestJoltage int64 = 0
		
		// we are gonna brute force that b and take every joltage possible and compare them to each other to find the largest
		for i := 0; i < len(line); i++ {
			for j := i+1; j < len(line); j++ {
				joltageString := string(line[i:i+1]) + string(line[j:j+1])

				potentialJoltage, err := strconv.ParseInt(joltageString, 10, 64)
				check(err)

				if potentialJoltage > largestJoltage {
					largestJoltage = potentialJoltage
				}
			}
		}

		joltageSum += largestJoltage
	}

	fmt.Print("Joltage sum: ")
	fmt.Println(joltageSum)

}