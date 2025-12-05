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
	// southwest
	if x > 0 && y < len(theArray)-1 {
		if (theArray[y+1][x-1:x] == "@") {
			adjacentRollCount++
		}
	}

	// west
	if x > 0 {
		if (theArray[y][x-1:x] == "@") {
			adjacentRollCount++
		}
	}

	// northwest
	if x > 0 && y > 0 {
		if (theArray[y-1][x-1:x] == "@") {
			adjacentRollCount++
		}
	}

	// north
	if y > 0 {
		if (theArray[y-1][x:x+1] == "@") {
			adjacentRollCount++
		}
	}

	// northeast
	if x < len(theArray[y])+1 && y > 0 {
		if x == len(theArray[y])-1 {
			if (theArray[y-1][x-1:] == "@") {
				adjacentRollCount++
			}
		} else {
			if (theArray[y-1][x+1:x+2] == "@") {
				adjacentRollCount++
			}
		}
		
	}

	// east 
	if x < len(theArray[y])+1 {
		if x == len(theArray[y])-1 {
			if (theArray[y][x-1:] == "@") {
				adjacentRollCount++
			}
		} else {
			if (theArray[y][x+1:x+2] == "@") {
				adjacentRollCount++
			}
		}
	}

	// southeast
	if x < len(theArray[y])+1 && y < len(theArray)-1 {
		if x == len(theArray[y])-1 {
			if (theArray[y+1][x-1:] == "@") {
				adjacentRollCount++
			}
		} else {
			if (theArray[y+1][x+1:x+2] == "@") {
				adjacentRollCount++
			}
		}
	}

	// south
	if y < len(theArray)-1 {
		if (theArray[y+1][x:x+1] == "@") {
			adjacentRollCount++
		}
	}

	return adjacentRollCount
}

func main() {
	// read in file input
	contents, err := os.Open("first_input.txt")
	check(err)
	var twoD []string
	var rollsRemoved int = 0
	
	scanner := bufio.NewScanner(contents)

	for scanner.Scan() {
		// read lines and split by commas
		line := scanner.Text()

		twoD = append(twoD, line)
	}

	for true {
		var currentRollsRemoved int = 0
		var rollsToRemove []int

		for i := 0; i < len(twoD); i++ {
			for j := 0; j < len(twoD[i]); j++ {
				if twoD[i][j:j+1] == "@" {
					if checkPositions(j, i, twoD) < 4 {
						rollsRemoved++
						currentRollsRemoved++

						rollsToRemove = append(rollsToRemove, i)
						rollsToRemove = append(rollsToRemove, j)
					}
				}
			}
		}

		// actually remove the marked rolls since we are done with an iteration
		for k := 0; k < len(rollsToRemove); k += 2 {
			// remove all these rolls
			// aka mark them . in original array
			
			var pos1 int = rollsToRemove[k]
			var pos2 int = rollsToRemove[k+1]

			// weird rune shit Go needs to do since strings are immutable
			// Thanks for wasting hours of my time Go
			var r = []rune(twoD[pos1])
			r[pos2] = '.'
			twoD[pos1] = string(r)
		}

		if currentRollsRemoved < 1 {
			break
		} else {
			currentRollsRemoved = 0
		}
	}

	fmt.Print("Rolls removed: ")
	fmt.Println(rollsRemoved)
}