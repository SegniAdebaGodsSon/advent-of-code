package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	var matrix [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token := scanner.Text()
		row := strings.Split(token, "")
		var intRow []int
		for _, str := range row {
			num := strToInt(str)
			intRow = append(intRow, num)
		}

		matrix = append(matrix, intRow)
	}

	printIntMatrix(matrix)

	visibleCount := partOneNaive(matrix)
	println("Visible count: ", visibleCount)
	println()

	println("Max scenic score: ", partTwoNaive(matrix))

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
func partOneNaive(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	visible := make([][]bool, rows)

	for i := range visible {
		visible[i] = make([]bool, cols)
	}

	for rowInd, row := range matrix {
		visible[rowInd] = visibleFromLeft(row)
		visibleFromRightArr := visibleFromRight(row)
		for colInd, val := range visible[rowInd] {
			visible[rowInd][colInd] = val || visibleFromRightArr[colInd]
		}
	}

	// check visibility from top
	for col := 0; col < cols; col++ {
		max := -1
		for row := 0; row < rows; row++ {
			value := matrix[row][col]
			visValue := visible[row][col]

			if value > max {
				visible[row][col] = visValue || true
				max = value
			} else {
				visible[row][col] = visValue || false
			}
		}
	}

	// check visibility from bottom
	for col := 0; col < cols; col++ {
		max := -1
		for row := rows - 1; row > -1; row-- {
			value := matrix[row][col]
			visValue := visible[row][col]

			if value > max {
				visible[row][col] = visValue || true
				max = value
			} else {
				visible[row][col] = visValue || false
			}
		}
	}

	visibleCount := 0

	for _, row := range visible {
		for _, col := range row {
			if col {
				visibleCount++
			}
		}
	}

	printBoolMatrix(visible)
	return visibleCount
}

func partTwoNaive(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	maxScenicScore := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			scenicScore := countDistantDown(row, col, matrix) *
				countDistantUp(row, col, matrix) *
				countDistantLeft(row, col, matrix) *
				countDistantRight(row, col, matrix)
			maxScenicScore = int(math.Max(float64(maxScenicScore), float64(scenicScore)))
		}
	}
	return maxScenicScore
}

func countDistantUp(row int, col int, matrix [][]int) int {
	count := 0
	value := matrix[row][col]
	row--
	for ; row > -1; row-- {
		currVal := matrix[row][col]
		if value > currVal {
			count++
		} else {
			return count + 1
		}
	}
	return count
}

func countDistantDown(row int, col int, matrix [][]int) int {
	rows := len(matrix)
	count := 0
	value := matrix[row][col]
	row++
	for ; row < rows; row++ {
		currVal := matrix[row][col]
		if value > currVal {
			count++
		} else {
			return count + 1
		}
	}
	return count
}

func countDistantLeft(row int, col int, matrix [][]int) int {
	count := 0
	value := matrix[row][col]
	col--
	for ; col > -1; col-- {
		currVal := matrix[row][col]
		if value > currVal {
			count++
		} else {
			return count + 1
		}
	}
	return count
}

func countDistantRight(row int, col int, matrix [][]int) int {
	cols := len(matrix[0])
	count := 0
	value := matrix[row][col]
	col++
	for ; col < cols; col++ {
		currVal := matrix[row][col]
		if value > currVal {
			count++
		} else {
			return count + 1
		}
	}
	return count
}

func strToInt(strNum string) int {
	num, err := strconv.Atoi(strNum)
	if err != nil {
		panic(err)
	}
	return num
}

func printIntMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Print("[")
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Print("]\n")
	}
}

func printBoolMatrix(matrix [][]bool) {
	for _, row := range matrix {
		fmt.Print("[")
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Print("]\n")
	}
}

func visibleFromLeft(row []int) (isVisible []bool) {
	max := -1

	for _, value := range row {
		if value > max {
			isVisible = append(isVisible, true)
			max = value
		} else {
			isVisible = append(isVisible, false)
		}
	}
	return
}

func visibleFromRight(row []int) (isVisible []bool) {
	max := -1

	for index := len(row) - 1; index > -1; index-- {
		value := row[index]
		if value > max {
			isVisible = append(isVisible, true)
			max = value
		} else {
			isVisible = append(isVisible, false)
		}
	}

	return reverseArray(isVisible)
}

func reverseArray(array []bool) []bool {
	length := len(array)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}
