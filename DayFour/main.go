package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fullyContained := 0

	for scanner.Scan() {

		line := scanner.Text()

		r := strings.NewReplacer(",", " ", "-", " ")
		parsed := r.Replace(line)

		var stringSlices []string
		var intSlices []int

		numberString := ""
		for i := 0; i < len(parsed); i++ {
			if parsed[i] != ' ' {
				numberString += string(parsed[i])
			} else if numberString != "" {
				numberString = strings.TrimSpace(numberString)
				stringSlices = append(stringSlices, numberString)
				numberString = ""
			}
		}
		if numberString != "" {
			numberString = strings.TrimSpace(numberString)
			stringSlices = append(stringSlices, numberString)
		}

		for _, item := range stringSlices {
			atoi, err := strconv.Atoi(item)
			if err != nil {
				return
			}
			intSlices = append(intSlices, atoi)

		}
		if (intSlices[0] <= intSlices[2] && intSlices[1] >= intSlices[3]) || (intSlices[2] <= intSlices[0] && intSlices[3] >= intSlices[1]) {
			fullyContained += 1
		}
	}
	fmt.Println(fullyContained)
}
