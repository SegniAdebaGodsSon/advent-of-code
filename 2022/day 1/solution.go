package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var topThreeElfs [3]int
	currCount := 0
	elf := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		token := scanner.Text()
		if token == "" {
			sort.Ints(topThreeElfs[:])
			topThreeElfs[0] = int(math.Max(float64(topThreeElfs[0]), float64(currCount)))
			currCount = 0
		} else {
			tokenNumber, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			currCount += tokenNumber
		}
		elf++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	totalCalories := 0
	for _, value := range topThreeElfs {
		totalCalories += value
	}
	fmt.Println("Max calories: ", topThreeElfs)
	fmt.Println("Sum of top 3 calories: ", totalCalories)

}
