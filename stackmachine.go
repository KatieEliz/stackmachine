package main

import (
	"errors"
	"strconv"
	"strings"
)

type StackMachine struct {
	stack []int
}

func NewStackMachine() *StackMachine {
	return &StackMachine{stack: []int{}}
}

func stackMachine(commands string) (int, error) {
	sm := NewStackMachine()
	cmdList := strings.Fields(commands)
	for _, cmd := range cmdList {
		switch cmd {
		case "POP":
			_, err := sm.Pop()
			if err != nil {
				return 0, errors.New("")
			}
		case "DUP":
			err := sm.Dup()
			if err != nil {
				return 0, errors.New("")
			}
		case "+":
			err := sm.Add()
			if err != nil {
				return 0, errors.New("")
			}
		case "-":
			err := sm.Subtract()
			if err != nil {
				return 0, errors.New("")
			}
		case "*":
			err := sm.Multiply()
			if err != nil {
				return 0, errors.New("")
			}
		case "SUM":
			err := sm.Sum()
			if err != nil {
				return 0, errors.New("")
			}
		case "CLEAR":
			sm.Clear()
		default:
			num, err := strconv.Atoi(cmd)
			if err != nil {
				return 0, errors.New("")
			}
			err = sm.Push(num)
			if err != nil {
				return 0, errors.New("")
			}
		}
	}

	// After processing all commands, return the top of the stack as the result
	top, err := sm.Top()
	if err != nil {
		return 0, errors.New("")
	}
	return top, nil
}

func (sm *StackMachine) Push(value int) error {
	if value < 0 || value > 50000 {
		return errors.New("integer must be between 0 and 50000 inclusive")
	}
	sm.stack = append(sm.stack, value)
	return nil
}

func (sm *StackMachine) Pop() (int, error) {
	if len(sm.stack) == 0 {
		return 0, errors.New("empty stack")
	}
	value := sm.stack[len(sm.stack)-1]
	sm.stack = sm.stack[:len(sm.stack)-1]
	return value, nil
}

func (sm *StackMachine) Dup() error {
	if len(sm.stack) == 0 {
		return errors.New("empty stack")
	}
	value := sm.stack[len(sm.stack)-1]
	sm.stack = append(sm.stack, value)
	return nil
}

func (sm *StackMachine) Add() error {
	if len(sm.stack) < 2 {
		return errors.New("too few elements to add")
	}
	a, _ := sm.Pop()
	b, _ := sm.Pop()
	sum := a + b
	if sum < 0 || sum > 50000 {
		return errors.New("integer must be between 0 and 50000 inclusive")
	}
	return sm.Push(sum)
}

func (sm *StackMachine) Subtract() error {
	if len(sm.stack) < 2 {
		return errors.New("too few elements to subtract")
	}
	a, _ := sm.Pop()
	b, _ := sm.Pop()
	diff := a - b
	if diff < 0 || diff > 50000 {
		return errors.New("integer must be between 0 and 50000 inclusive")
	}
	return sm.Push(diff)
}

func (sm *StackMachine) Multiply() error {
	if len(sm.stack) < 2 {
		return errors.New("too few elements to multiply")
	}
	a, _ := sm.Pop()
	b, _ := sm.Pop()
	product := a * b
	if product < 0 || product > 50000 {
		return errors.New("integer must be between 0 and 50000 inclusive")
	}
	return sm.Push(product)
}

func (sm *StackMachine) Clear() {
	sm.stack = []int{}
}
func (sm *StackMachine) Sum() error {
	if len(sm.stack) == 0 {
		return errors.New("empty stack")
	}
	sum := 0
	for _, num := range sm.stack {
		sum += num
	}
	sm.Clear()
	sm.Push(sum)
	return nil
}

func (sm *StackMachine) IsEmpty() bool {
	return len(sm.stack) == 0
}
func (sm *StackMachine) Top() (int, error) {
	if len(sm.stack) == 0 {
		return 0, errors.New("stack is empty")
	}
	return sm.stack[len(sm.stack)-1], nil
}
