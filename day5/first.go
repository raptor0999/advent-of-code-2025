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
	var ids []int64
	var processingFreshIds bool = true
	var freshCount int = 0

	// read in file input
	contents, err := os.Open("first_input.txt")
	check(err)
	
	scanner := bufio.NewScanner(contents)

	for scanner.Scan() {
		// read lines and split by commas
		
		line := scanner.Text()

		if line == "" {
			processingFreshIds = false
		} else {
			if processingFreshIds {
				// we are catalouging all the fresh ids in the ranges
				items := strings.Split(line, "-")

				start, err := strconv.ParseInt(items[0], 10, 64)
				check(err)
				end, err := strconv.ParseInt(items[1], 10, 64)
				check(err)

				ids = append(ids, start)
				ids = append(ids, end)
			} else {
				// we are checking ids against the fresh ids now
				var isIn bool = false
				id, err := strconv.ParseInt(line, 10, 64)
				check(err)

				for i := 0; i < len(ids); i += 2 {
					if id >= ids[i] && id <= ids[i+1] {
						isIn = true
					}
				}

				if isIn {
					freshCount++
				}
			}
		}
	}

	fmt.Println(freshCount)
}