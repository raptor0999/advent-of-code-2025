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
				halfOrSo := stringIntLength / 2 + stringIntLength % 2
				var counted bool = false

				// half or so because we do not need to go further to find repeats than half of the length of the string
				for j := 1; j <= halfOrSo; j++ {
					
					// only if the string length is evenly divisible
					if stringIntLength % j == 0 {
						var lastSection string = ""
						var invalid bool = true

						// split the string into equal sections
						for k := 0; k <= stringIntLength - j; k += j {
							// make sure to automatically count numbers of length 1 as valid
							if (lastSection != "" && lastSection != stringInt[k:k+j]) || stringIntLength == 1 {
								invalid = false
							}

							lastSection = stringInt[k:k+j]
						}

						if invalid && !counted {
							fmt.Println("Invalid value: " + stringInt)
							invalidSum += i
							counted = true
						}
					}
				}
			}

			fmt.Println("End")
		}
	}

	fmt.Println("Invalid sum is : " + strconv.FormatInt(int64(invalidSum), 10))
}