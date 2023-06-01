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

	count := 0
	for scanner.Scan() {
		token := scanner.Text()
		sections := strings.Split(token, ",")
		sectionOne := sections[0]
		sectionTwo := sections[1]
		if partTwo(sectionOne, sectionTwo) {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(count)
}

func partOne(sectionOne string, sectionTwo string) bool {
	sectionOneSplit := strings.Split(sectionOne, "-")
	sectionTwoSplit := strings.Split(sectionTwo, "-")

	sectionOneFrom, err := strconv.Atoi(sectionOneSplit[0])
	if err != nil {
		panic(err)
	}

	sectionOneTo, err := strconv.Atoi(sectionOneSplit[1])
	if err != nil {
		panic(err)
	}

	sectionTwoFrom, err := strconv.Atoi(sectionTwoSplit[0])
	if err != nil {
		panic(err)
	}
	sectionTwoTo, err := strconv.Atoi(sectionTwoSplit[1])
	if err != nil {
		panic(err)
	}

	if sectionOneFrom == sectionTwoFrom || sectionOneTo == sectionTwoTo {
		return true
	} else {
		if sectionOneFrom < sectionTwoFrom && sectionOneTo > sectionTwoTo {
			return true
		}

		if sectionOneFrom > sectionTwoFrom && sectionOneTo < sectionTwoTo {
			return true
		}
	}

	return false
}

func partTwo(sectionOne string, sectionTwo string) bool {
	sectionOneSplit := strings.Split(sectionOne, "-")
	sectionTwoSplit := strings.Split(sectionTwo, "-")

	sectionOneFrom, err := strconv.Atoi(sectionOneSplit[0])
	if err != nil {
		panic(err)
	}

	sectionOneTo, err := strconv.Atoi(sectionOneSplit[1])
	if err != nil {
		panic(err)
	}

	sectionTwoFrom, err := strconv.Atoi(sectionTwoSplit[0])
	if err != nil {
		panic(err)
	}
	sectionTwoTo, err := strconv.Atoi(sectionTwoSplit[1])
	if err != nil {
		panic(err)
	}

	if sectionOneFrom == sectionTwoFrom {
		return true
	}

	if sectionOneFrom < sectionTwoFrom && sectionOneTo >= sectionTwoFrom {
		return true
	}

	if sectionTwoFrom < sectionOneFrom && sectionTwoTo >= sectionOneFrom {
		return true
	}
	return false
}

// sectionOneFrom .... sectionOneTo
// sectionTwoFrom .... sectionTwoTo
