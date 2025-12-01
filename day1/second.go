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

func computeNextDialValue(currentVal int64, leftOrRight string, modifier int64) (int64, int64) {
	var zeroes int64 = 0


	if leftOrRight == "L" {
		// subtract
		for i := 0; i < int(modifier); i++ {
			currentVal -= 1

			if currentVal%100 == 0 {
				zeroes += 1
			}
		}
	} else {
		// add
		for i := 0; i < int(modifier); i++ {
			currentVal += 1

			if currentVal%100 == 0 {
				zeroes += 1
			}
		}
	}

	fmt.Print("Dial Value modified to: ")
	fmt.Println(currentVal)

	if currentVal >= 0 && currentVal < 100 {
		fmt.Print("Dial Value is a now: ")
		fmt.Println(currentVal)

		return currentVal, zeroes
	} else {
		if currentVal < 0 {
			for currentVal < 0 {
				currentVal += 100
			}

			
		} 

		if currentVal > 99 {
			for currentVal > 99 {
				currentVal -= 100
			}

			
		}
	}

	fmt.Print("Dial Value is a now: ")
	fmt.Println(currentVal)

	return currentVal, zeroes
}

func main() {
	// set initial vars for work
	var dialValue int64 = 50
	var zeroes int64 = 0
	var zeroesSum int64 = 0

	// read in file input
	contents, err := os.Open("first_input.txt")
	check(err)
	
	scanner := bufio.NewScanner(contents)

	fmt.Println("Dial value is: ", dialValue)

	for scanner.Scan() {
		line := scanner.Text()

		currentVal, err := strconv.ParseInt(line[1:len(line)], 10, 64)
		check(err)
		dialValue, zeroes = computeNextDialValue(dialValue, line[0:1], int64(currentVal))

		zeroesSum += zeroes

		fmt.Println("Zeros Sum: ")
		fmt.Println(zeroesSum)
	}

	fmt.Println("Zeros Sum: ")
	fmt.Println(zeroesSum)
}