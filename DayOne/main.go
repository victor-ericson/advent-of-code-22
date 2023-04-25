package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//skapar current och max calories
	currentCalories := 0
	maxCalories := 0

	//
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
		} else {
			sum, _ := strconv.Atoi(line)
			currentCalories += sum
		}
	}
	fmt.Println(maxCalories)
}
