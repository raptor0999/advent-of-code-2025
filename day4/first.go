package main

import (
	"bufio"
	"fmt"
	"os")

func check(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
		panic(e)
	}
}

func checkPositions(x int, y int, theArray []string) int {
	var adjacentRollCount int = 0

	// start from southwest and go clockwise
	fmt.Println(x)
	fmt.Println(y)

	// southwest
	if x > 0 && y < len(theArray)-1 {
		if (theArray[y+1][x-1:x] == "@") {
			adjacentRollCount++
			fmt.Print("sw ")
		}
	}

	// west
	if x > 0 {
		if (theArray[y][x-1:x] == "@") {
			adjacentRollCount++
			fmt.Print("w ")
		}
	}

	// northwest
	if x > 0 && y > 0 {
		if (theArray[y-1][x-1:x] == "@") {
			adjacentRollCount++
			fmt.Print("nw ")
		}
	}

	// north
	if y > 0 {
		if (theArray[y-1][x:x+1] == "@") {
			adjacentRollCount++
			fmt.Print("n ")
		}
	}

	// northeast
	if x < len(theArray[y])+1 && y > 0 {
		if x == len(theArray[y])-1 {
			if (theArray[y-1][x-1:] == "@") {
				adjacentRollCount++
				fmt.Print("nE ")
			}
		} else {
			if (theArray[y-1][x+1:x+2] == "@") {
				adjacentRollCount++
				fmt.Print("ne ")
			}
		}
		
	}

	// east 
	if x < len(theArray[y])+1 {
		if x == len(theArray[y])-1 {
			if (theArray[y][x-1:] == "@") {
				adjacentRollCount++
				fmt.Print("E ")
			}
		} else {
			if (theArray[y][x+1:x+2] == "@") {
				adjacentRollCount++
				fmt.Print("e ")
			}
		}
	}

	// southeast
	if x < len(theArray[y])+1 && y < len(theArray)-1 {
		if x == len(theArray[y])-1 {
			if (theArray[y+1][x-1:] == "@") {
				adjacentRollCount++
				fmt.Print("sE ")
			}
		} else {
			if (theArray[y+1][x+1:x+2] == "@") {
				adjacentRollCount++
				fmt.Print("se ")
			}
		}
	}

	// south
	if y < len(theArray)-1 {
		if (theArray[y+1][x:x+1] == "@") {
			adjacentRollCount++
			fmt.Print("s ")
		}
	}

	fmt.Println("")

	return adjacentRollCount
}

func main() {
	// read in file input
	contents, err := os.Open("first_input.txt")
	check(err)

	var twoD []string
	var rollsAvailable int = 0
	
	scanner := bufio.NewScanner(contents)

	for scanner.Scan() {
		// read lines and split by commas
		line := scanner.Text()

		twoD = append(twoD, line)
	}

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[i]); j++ {
			if twoD[i][j:j+1] == "@" {
				if checkPositions(j, i, twoD) < 4 {
					rollsAvailable++
				}
			}
		}
	}

	fmt.Print("Rolls available: ")
	fmt.Println(rollsAvailable)
}