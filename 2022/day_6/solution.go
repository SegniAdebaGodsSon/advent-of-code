package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var datastream string
	if scanner.Scan() {
		datastream = scanner.Text()
	}

	solution := part_two(datastream)
	println(solution)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func part_one(datastream string) int {
	length := len(datastream)
	tracker := make(map[rune]int, 4)

	for i := 0; i < length; {
		char := rune(datastream[i])

		if index, ok := tracker[char]; ok {
			i = index + 1
			for k := range tracker {
				delete(tracker, k)
			}
		} else {
			tracker[char] = i
			if len(tracker) == 4 {
				return i + 1
			}
			i++
		}

	}
	return -1
}

func part_two(datastream string) int {
	length := len(datastream)
	tracker := make(map[rune]int, 14)

	for i := 0; i < length; {
		char := rune(datastream[i])

		if index, ok := tracker[char]; ok {
			i = index + 1
			for k := range tracker {
				delete(tracker, k)
			}
		} else {
			tracker[char] = i
			if len(tracker) == 14 {
				return i + 1
			}
			i++
		}

	}
	return -1
}
