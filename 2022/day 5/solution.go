package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	length := len(s.items)
	if length == 0 {
		panic("Stack is empty")
	}
	item := s.items[length-1]
	s.items = s.items[:length-1]
	return item
}

func (s *Stack[T]) Peek() T {
	length := len(s.items)
	if length == 0 {
		panic("Stack is empty")
	}
	return s.items[length-1]
}

func (s *Stack[T]) Print() {
	length := len(s.items)
	for i := length - 1; i > -1; i-- {
		fmt.Printf("[ %v ]\n", s.items[i])
	}
}

func NewStack[T any](items []T) *Stack[T] {
	return &Stack[T]{
		items: items,
	}
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	stack_1_items := []string{"D", "L", "V", "T", "M", "H", "F"}
	stack_1 := NewStack(stack_1_items)

	stack_2_items := []string{"H", "Q", "G", "J", "C", "T", "N", "P"}
	stack_2 := NewStack(stack_2_items)

	stack_3_items := []string{"R", "S", "D", "M", "P", "H"}
	stack_3 := NewStack(stack_3_items)

	stack_4_items := []string{"L", "B", "V", "F"}
	stack_4 := NewStack(stack_4_items)

	stack_5_items := []string{"N", "H", "G", "L", "Q"}
	stack_5 := NewStack(stack_5_items)

	stack_6_items := []string{"W", "B", "D", "G", "R", "M", "P"}
	stack_6 := NewStack(stack_6_items)

	stack_7_items := []string{"G", "M", "N", "R", "C", "H", "L", "Q"}
	stack_7 := NewStack(stack_7_items)

	stack_8_items := []string{"C", "L", "W"}
	stack_8 := NewStack(stack_8_items)

	stack_9_items := []string{"R", "D", "L", "Q", "J", "Z", "M", "T"}
	stack_9 := NewStack(stack_9_items)

	stacks := []*Stack[string]{stack_1, stack_2, stack_3, stack_4, stack_5, stack_6, stack_7, stack_8, stack_9}

	scanner := bufio.NewScanner(file)
	inputReached := false
	for scanner.Scan() {
		token := scanner.Text()
		if inputReached {
			splitToken := strings.Split(token, " ")
			move, err := strconv.Atoi(splitToken[1])
			if err != nil {
				panic(err)
			}

			from, err := strconv.Atoi(splitToken[3])
			if err != nil {
				panic(err)
			}

			to, err := strconv.Atoi(splitToken[5])
			if err != nil {
				panic(err)
			}
			part_2(move, from-1, to-1, stacks)
		}
		if token == "" {
			inputReached = true
		}
	}

	result := ""

	for i := 0; i < len(stacks); i++ {
		result += stacks[i].Peek()
	}
	println(result)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func part_1(move int, from int, to int, stacks []*Stack[string]) {
	for move > 0 {
		peek := stacks[from].Pop()
		stacks[to].Push(peek)
		move--
	}
}

func part_2(move int, from int, to int, stacks []*Stack[string]) {
	tempStack := NewStack([]string{})
	for i := move; i > 0; i-- {
		peek := stacks[from].Pop()
		tempStack.Push(peek)
	}

	for i := move; i > 0; i-- {
		peek := tempStack.Pop()
		stacks[to].Push(peek)
	}
}

func printStacks(stacks []*Stack[string]) {
	for i := 0; i < len(stacks); i++ {
		fmt.Printf("______STACK: %v ______ \n", i+1)
		stacks[i].Print()
		fmt.Println()
	}
}
