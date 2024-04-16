package main

import "fmt"

type Stack[T any] struct {
	data []T
}

func (a *Stack[T]) Push(V T) {
	a.data = append(a.data, V)
}

func (a *Stack[T]) Pop() (T, bool) {
	if len(a.data) == 0 {
		var r T
		return r, false
	}

	lastIndex := len(a.data) - 1
	popped := a.data[lastIndex]
	a.data = a.data[:lastIndex]

	return popped, true
}

func (a *Stack[T]) Size() int {
	return len(a.data)
}

func runStack() {
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println(intStack.Pop())  // Output: 3, true
	fmt.Println(intStack.Size()) // Output: 2

	intStack.Push(5)           // Output: 3, true
	fmt.Println(intStack.data) // Output: 2

	stringStack := Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")

	fmt.Println(stringStack.Pop()) // Output: world, true
	fmt.Println(stringStack.data)  // Output: 1
}
