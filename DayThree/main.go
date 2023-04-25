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

	charMap := make(map[string]int64)

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 1; i < 53; i++ {
		charMap[string(letters[i-1])] = int64(i)
	}

	fmt.Println(charMap)

	var prioritySum int64 = 0

	for scanner.Scan() {

		line := scanner.Text()

		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]

		foundMatch := false

		for _, char1 := range firstHalf {
			for _, char2 := range secondHalf {
				if char1 == char2 && !foundMatch {
					prioritySum += charMap[string(char1)]
					foundMatch = true
				}
			}
		}
	}

	fmt.Println(prioritySum)
}
