package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	points := 0

	for scanner.Scan() {
		line := scanner.Text()

		switch line[0] {
		case 'A':
			switch line[2] {
			case 'X':
				points += (1 + 3)
			case 'Y':
				points += (2 + 6)
			case 'Z':
				points += (3 + 0)
			}
		case 'B':
			switch line[2] {
			case 'X':
				points += (1 + 0)
			case 'Y':
				points += (2 + 3)
			case 'Z':
				points += (3 + 6)
			}
		case 'C':
			switch line[2] {
			case 'X':
				points += (1 + 6)
			case 'Y':
				points += (2 + 0)
			case 'Z':
				points += (3 + 3)
			}
		}
	}
	fmt.Print(points)
}
