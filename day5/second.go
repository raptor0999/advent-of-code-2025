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

func main() {
	// read in file input
	contents, err := os.Open("first_input.txt")
	check(err)
	
	scanner := bufio.NewScanner(contents)

	for scanner.Scan() {
		// read lines and split by commas
		line := scanner.Text()

		fmt.Println(line)
	}
}