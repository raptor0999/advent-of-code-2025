package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv")

func check(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
		panic(e)
	}
}

func main() {
	// set initial vars for work
	var invalidSum int64 = 0

	// read in file input
	contents, err := os.Open("first_input.txt")
	check(err)
	
	scanner := bufio.NewScanner(contents)

	for scanner.Scan() {
		// read lines and split by commas
		line := scanner.Text()
		items := strings.Split(line, ",")
		
		for _, item := range items {
			ranges := strings.Split(item, "-")

			// convert to integers so we can iterate over range
			beginRange, err := strconv.ParseInt(ranges[0], 10, 64)
			check(err)

			endRange, err := strconv.ParseInt(ranges[1], 10, 64)
			check(err)

			fmt.Println("Begin")

			for i := beginRange; i <= endRange; i++ {
				// convert each number back to string so we can split in half and see if they are duplicating (invalid)
				stringInt := strconv.FormatInt(int64(i), 10)
				stringIntLength := len(stringInt)
				half := stringIntLength / 2

				firstHalf := stringInt[0:half]
				secondHalf := stringInt[half:stringIntLength]

				if firstHalf == secondHalf {
					fmt.Println("Invalid")
					fmt.Println(stringInt)

					invalidSum += i
				}
				
			}

			fmt.Println("End")
		}
	}

	fmt.Println("Invalid sum is : " + strconv.FormatInt(int64(invalidSum), 10))
}