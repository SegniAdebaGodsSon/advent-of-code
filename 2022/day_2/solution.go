package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	p1_rock    = 'A'
	p1_paper   = 'B'
	p1_scissor = 'C'

	p2_rock    = 'X'
	p2_paper   = 'Y'
	p2_scissor = 'Z'

	outcome_loss = 'X'
	outcome_draw = 'Y'
	outcome_win  = 'Z'
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token := scanner.Text()
		characters := strings.Split(token, " ")
		playerOne := rune(characters[0][0])
		outcome := rune(characters[1][0])
		score += play(playerOne, outcome)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(score)
}

func play(playerOne rune, outcome rune) int {
	const loss, draw, win = 0, 3, 6
	result := 0
	var playerTwo rune

	if outcome == outcome_draw {
		playerTwo = playerOne + 23
		result += draw
	} else if outcome == outcome_loss {
		previous := playerOne - rune(1)
		if previous < 65 {
			previous = 67
		}
		playerTwo = previous + 23
		result += loss
	} else {
		next := playerOne + rune(1)
		if next > 67 {
			next = 65
		}
		playerTwo = next
		result += win
	}

	return result + int(playerTwo) - 87
}
