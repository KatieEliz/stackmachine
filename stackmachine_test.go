package main

import (
	"errors"
	"testing"
)

func TestStartsWithEmptyStack(t *testing.T) {
	_, err := stackMachine("")

	if err == nil {
		t.Error("expected error due to no results")
	}
}

func TestStackMachineInRange(t *testing.T) {
	sm := &StackMachine{}

	err := sm.Push(100)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

}
func TestStackMachineOutOfRange(t *testing.T) {
	sm := &StackMachine{}

	err := sm.Push(-1)
	if err == nil {
		t.Error("Expected error, integer out of range")
	}
}
func TestStackMachineOnBoundary(t *testing.T) {
	sm := &StackMachine{}

	err := sm.Push(50000)
	if err != nil {
		t.Error("Expected no error, num on boundary")
	}
}
func TestPushToMaximumCapacity(t *testing.T) {
	sm := &StackMachine{}

	err := sm.Push(50000)
	if err != nil {
		t.Error("unexpected error")
	}
}
func TestPushMulitipleElements(t *testing.T) {
	sm := &StackMachine{}

	for _, num := range []int{10, 20, 30} {
		err := sm.Push(num)
		if err != nil {
			t.Error("unexpected error")
		}
	}
}

/*func TestSequenceOfOperations(t *testing.T) {
	sm := &StackMachine{
		stack: []int{10, 20, 30},
	}

	sm.Add()
	sm.Multiply()
	sm.Dup()

	expectedStack := []int{500, 500, 30}
	for i := range expectedStack {
		if sm.stack[i] != expectedStack[i] {
			t.Errorf("Expected stack element at index %d to be %d, got %d", i, expectedStack[i], sm.stack[i])
		}
	}
}*/

// pop
func TestPopFromEmptyStack(t *testing.T) {
	sm := &StackMachine{}

	_, err := sm.Pop()
	if err == nil {
		t.Error("expected error when popping from empty stack")
	}
}
func TestPopFromNonEmptyStack(t *testing.T) {
	sm := &StackMachine{}
	sm.Push(1)
	sm.Push(2)
	sm.Push(3)

	_, err := sm.Pop()
	if err != nil {
		t.Error("unexpected error")
	}
}
func TestFirstElementAfterPop(t *testing.T) {
	sm := &StackMachine{}
	sm.Push(1)
	sm.Push(2)
	sm.Push(3)

	sm.Pop()

	top := sm.stack[len(sm.stack)-1]
	expected := 2
	if top != expected {
		t.Error("expected new top element")
	}
}
func TestPoppedValue(t *testing.T) {
	sm := &StackMachine{}
	sm.Push(1)
	sm.Push(2)
	sm.Push(3)

	popped, _ := sm.Pop()
	if popped != 3 {
		t.Error("expected popped value to be 3")
	}
}
func TestPoppedLength(t *testing.T) {
	sm := &StackMachine{}
	sm.Push(1)
	sm.Push(2)
	sm.Push(3)

	sm.Pop()

	if len(sm.stack) != 2 {
		t.Error("expected stack lenth after pop to be 3")
	}

}
func TestPoppedFromSingleElementStack(t *testing.T) {
	sm := &StackMachine{}
	sm.Push(50)

	popped, _ := sm.Pop()
	if popped != 50 {
		t.Error("expected popped value to be 42")
	}
}

// dup
func TestSucsessOfDuplicationOfStack(t *testing.T) {
	sm := &StackMachine{}
	sm.Push(1)
	sm.Push(2)
	sm.Push(3)

	err := sm.Dup()
	if err != nil {
		t.Error("duplication failed")
	}
}
func TestDulpicationFromEmptyStack(t *testing.T) {
	sm := &StackMachine{}

	err := sm.Dup()
	if err == nil {
		t.Error("expected error, cannot duplicate from empty stack")
	}
}
func TestTopElementAfterDuplication(t *testing.T) {
	sm := &StackMachine{}
	sm.Push(1)
	sm.Push(2)
	sm.Push(3)

	sm.Dup()
	top := sm.stack[len(sm.stack)-1]
	expected := 3

	if top != expected {
		t.Error("unexpected top element")
	}
}

// add
func TestAddingOnEmptyStack(t *testing.T) {
	sm := &StackMachine{
		stack: []int{},
	}

	err := sm.Add()
	expectedError := "too few elements to add"

	if err == nil || err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%v'", expectedError, err)
	}
}
func TestAddingLessThanTwoElements(t *testing.T) {
	sm := &StackMachine{
		stack: []int{10},
	}

	err := sm.Add()
	expectedError := "too few elements to add"

	if err == nil || err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%v'", expectedError, err)
	}
}
func TestAddingTwoElements(t *testing.T) {
	sm := &StackMachine{
		stack: []int{10, 3},
	}

	sm.Add()

	if sm.stack[0] != 13 {
		t.Error("unexpected addition")
	}
}
func TestAddingOutOfRange(t *testing.T) {
	sm := StackMachine{
		stack: []int{50001, 80},
	}

	err := sm.Add()
	expectedError := errors.New("integer must be between 0 and 50000 inclusive")

	if err != nil && err.Error() != expectedError.Error() {
		t.Errorf("expected error '%s', got '%v'", expectedError, err)
	}
}
func TestForAddingTwoInRangeNumsThatAddOutOfRange(t *testing.T) {
	sm := StackMachine{
		stack: []int{50000, 2},
	}

	err := sm.Add()

	expectedError := "integer must be between 0 and 50000 inclusive"

	if err == nil || err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%v'", expectedError, err)
	}
}

// subtract
func TestSubtractionOnEmptyStack(t *testing.T) {
	sm := &StackMachine{
		stack: []int{},
	}

	err := sm.Subtract()
	expectedError := "too few elements to subtract"

	if err == nil || err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%v'", expectedError, err)
	}
}
func TestSubtractingLessThanTwoElements(t *testing.T) {
	sm := &StackMachine{
		stack: []int{10},
	}

	err := sm.Subtract()
	expectedError := "too few elements to subtract"

	if err == nil || err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%v'", expectedError, err)
	}
}
func TestSubtractingTwoElements(t *testing.T) {
	sm := &StackMachine{
		stack: []int{3, 10},
	}

	sm.Subtract()

	if sm.stack[0] != 7 {
		t.Error("unexpected subtraction")
	}
}
func TestForSubtractingTwoInRangeNumsThatAddOutOfRange(t *testing.T) {
	sm := StackMachine{
		stack: []int{2, 0},
	}

	err := sm.Subtract()

	expectedError := "integer must be between 0 and 50000 inclusive"

	if err == nil || err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%v'", expectedError, err)
	}
}
func TestSubtractingOutOfRange(t *testing.T) {
	sm := StackMachine{
		stack: []int{50001, 50002},
	}

	err := sm.Subtract()
	expectedError := errors.New("integer must be between 0 and 50000 inclusive")

	if err != nil && err.Error() != expectedError.Error() {
		t.Errorf("expected error '%s', got '%v'", expectedError, err)
	}
}

// multilpy
func TestMultiplyEmptyStack(t *testing.T) {
	sm := &StackMachine{
		stack: []int{},
	}

	err := sm.Multiply()
	expectedError := "too few elements to multiply"

	if err == nil || err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%v'", expectedError, err)
	}
}
func TestMulitiplyingLessThanTwoElements(t *testing.T) {
	sm := &StackMachine{
		stack: []int{10},
	}

	err := sm.Multiply()
	expectedError := "too few elements to multiply"

	if err == nil || err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%v'", expectedError, err)
	}
}
func TestMultiplyingTwoElements(t *testing.T) {
	sm := &StackMachine{
		stack: []int{3, 10},
	}

	sm.Multiply()

	if sm.stack[0] != 30 {
		t.Error("unexpected multiplication")
	}
}
func TestForMultiplyingTwoInRangeNumsThatAddOutOfRange(t *testing.T) {
	sm := StackMachine{
		stack: []int{900, 60},
	}

	err := sm.Multiply()

	expectedError := "integer must be between 0 and 50000 inclusive"

	if err == nil || err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%v'", expectedError, err)
	}
}
func TestMultiplyingOutOfRange(t *testing.T) {
	sm := StackMachine{
		stack: []int{50001, 50002},
	}

	err := sm.Multiply()
	expectedError := errors.New("integer must be between 0 and 50000 inclusive")

	if err != nil && err.Error() != expectedError.Error() {
		t.Errorf("expected error '%s', got '%v'", expectedError, err)
	}
}

// sum
func TestSum(t *testing.T) {
	sm := &StackMachine{
		stack: []int{1, 2, 3, 4, 5},
	}

	err := sm.Sum()
	if err != nil {
		t.Error("unexpected error")
	}
}

func TestSumOfElements(t *testing.T) {
	sm := &StackMachine{
		stack: []int{1, 2, 3, 4, 5},
	}
	sm.Sum()
	expectedStack := []int{15}
	for i := range expectedStack {
		if sm.stack[i] != expectedStack[i] {
			t.Errorf("Expected stack element at index %d to be %d, got %d", i, expectedStack[i], sm.stack[i])
		}
	}
}
func TestSumEmptyStack(t *testing.T) {
	sm := &StackMachine{
		stack: []int{},
	}

	err := sm.Sum()
	expectedError := "empty stack"

	if err == nil || err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%v'", expectedError, err)
	}
}
func TestSumSingleElement(t *testing.T) {
	sm := &StackMachine{
		stack: []int{20},
	}

	err := sm.Sum()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

//clear

func TestStackMachineIsClear(t *testing.T) {
	sm := &StackMachine{
		stack: []int{1, 2, 3, 4, 5},
	}

	sm.Clear()

	if len(sm.stack) != 0 {
		t.Error("expected stack to be empty")
	}

}
func TestStackMachineIsClearSingleIineteger(t *testing.T) {
	sm := &StackMachine{
		stack: []int{56},
	}

	sm.Clear()

	if len(sm.stack) != 0 {
		t.Error("expected tack to be empty")
	}
}
func TestCheckToSeeIfStackIsEmptyAfterOperations(t *testing.T) {
	sm := &StackMachine{
		stack: []int{2, 3, 4},
	}

	sm.Add()
	sm.Multiply()

	sm.Clear()

	if len(sm.stack) != 0 {
		t.Error("expected stack to be empty")
	}
}

//isEmpty

func TestCheckIfEmpty(t *testing.T) {
	sm := &StackMachine{
		stack: []int{},
	}

	if !sm.IsEmpty() {
		t.Error("expected isEmpty to return true for an empty stack")
	}
}
func TestIsEmptyAfterWithOneElement(t *testing.T) {
	sm := &StackMachine{
		stack: []int{45},
	}

	if sm.IsEmpty() {
		t.Error("expected isEmpty to return false for a stack with one element")
	}
}
func TestIsEmptyAfterClearing(t *testing.T) {
	sm := &StackMachine{
		stack: []int{45},
	}

	sm.Clear()

	if !sm.IsEmpty() {
		t.Errorf("expected IsEmpty to return true after clearing the stack")
	}
}

// top
func TestTopEmptyStack(t *testing.T) {
	sm := &StackMachine{
		stack: []int{},
	}

	_, err := sm.Top()
	if err == nil {
		t.Error("expected error when calling top on an empty stack")
	}
}
func TestTopSingleElementStack(t *testing.T) {
	sm := &StackMachine{
		stack: []int{42},
	}

	_, err := sm.Top()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestAcceptanceTests(t *testing.T) {
	tests := []struct {
		name        string
		commands    string
		expected    int
		expectedErr error
	}{
		{name: "empty error", commands: "", expected: 0, expectedErr: errors.New("")},
		{name: "add overflow", commands: "50000 DUP +", expected: 0, expectedErr: errors.New("")},
		{name: "too few add", commands: "99 +", expected: 0, expectedErr: errors.New("")},
		{name: "too few minus", commands: "99 -", expected: 0, expectedErr: errors.New("")},
		{name: "too few multiply", commands: "99 *", expected: 0, expectedErr: errors.New("")},
		{name: "empty stack", commands: "99 CLEAR", expected: 0, expectedErr: errors.New("")},
		{name: "sum single value", commands: "99 SUM", expected: 99, expectedErr: nil},
		{name: "sum empty", commands: "SUM", expected: 0, expectedErr: errors.New("")},
		{name: "normal +*", commands: "5 6 + 2 *", expected: 22, expectedErr: nil},
		{name: "clear too few", commands: "1 2 3 4 + CLEAR 12 +", expected: 0, expectedErr: errors.New("")},
		{name: "normal after clear", commands: "1 CLEAR 2 3 +", expected: 5, expectedErr: nil},
		{name: "single integer", commands: "9876", expected: 9876, expectedErr: nil},
		{name: "invalid command", commands: "DOGBANANA", expected: 0, expectedErr: errors.New("")},
		{name: "normal +-*", commands: "5 9 DUP + + 43 - 3 *", expected: 60, expectedErr: nil},
		{name: "minus", commands: "2 5 -", expected: 3, expectedErr: nil},
		{name: "underflow minus", commands: "5 2 -", expected: 0, expectedErr: errors.New("")},
		{name: "at overflow limit", commands: "25000 DUP +", expected: 50000, expectedErr: nil},
		{name: "at overflow limit single value", commands: "50000 0 +", expected: 50000, expectedErr: nil},
		{name: "overflow plus", commands: "50000 1 +", expected: 0, expectedErr: errors.New("")},
		{name: "overflow single value", commands: "50001", expected: 0, expectedErr: errors.New("")},
		{name: "times zero at overflow limit", commands: "50000 0 *", expected: 0, expectedErr: nil},
		{name: "too few at first", commands: "1 2 3 4 5 + + + + * 999", expected: 0, expectedErr: errors.New("")},
		{name: "normal simple", commands: "1 2 - 99 +", expected: 100, expectedErr: nil},
		{name: "at overflow minus to zero", commands: "50000 50000 -", expected: 0, expectedErr: nil},
		{name: "clear empties stack", commands: "CLEAR", expected: 0, expectedErr: errors.New("")},
		{name: "normal sum", commands: "3 4 3 5 5 1 1 1 SUM", expected: 23, expectedErr: nil},
		{name: "sum after clear stack", commands: "3 4 3 5 CLEAR 5 1 1 1 SUM", expected: 8, expectedErr: nil},
		{name: "sum then too few", commands: "3 4 3 5 5 1 1 1 SUM -", expected: 0, expectedErr: errors.New("")},
		{name: "fibonacci", commands: "1 2 3 4 5 * * * *", expected: 120, expectedErr: nil},
	}

	for _, test := range tests {

		got, err := stackMachine(test.commands)

		if test.expectedErr != nil {
			if err == nil {
				t.Errorf("%s (%s) Expected error, but received nil", test.name, test.commands)
			} else if err.Error() != test.expectedErr.Error() {
				t.Errorf("%s (%s) got error %v, want %v", test.name, test.commands, err, test.expectedErr)
			}
		} else if got != test.expected {
			t.Errorf("%s (%s) got %v, want %v", test.name, test.commands, got, test.expected)
		}
	}
}
