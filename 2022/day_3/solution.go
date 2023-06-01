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
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	cycle := 1
	var group [3]string
	for scanner.Scan() {
		token := scanner.Text()
		group[cycle-1] = token
		if cycle%3 == 0 {
			total += helper(group)
			cycle = 0
		}
		cycle++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Final answer: ", total)
}

func helper(group [3]string) int {
	rucksack_1 := group[0]
	rucksack_2 := group[1]
	rucksack_3 := group[2]
	map_1 := make(map[rune]int)
	map_2 := make(map[rune]int)

	for _, char := range rucksack_1 {
		map_1[char] = int(char)
	}

	for _, char := range rucksack_2 {
		map_2[char] = int(char)
	}

	for _, char := range rucksack_3 {
		if _, ok_1 := map_1[char]; ok_1 {
			if _, ok_2 := map_2[char]; ok_2 {
				return priorityOf(char)
			}
		}
	}
	return 0
}

func priorityOf(char rune) int {
	if 'a' <= char && char <= 'z' {
		return int(char) - int('a') + 1
	} else {
		return int(char) - int('A') + 1 + 26
	}
}
