package main

import "fmt"

type Queue struct {
	head, tail *Node
	size       int
}

type Node struct {
	next, prev *Node
	value      int
}

func (a *Queue) Push(v int) {
	if a.tail == nil {
		a.tail = &Node{
			value: v,
		}
		a.head = a.tail
	} else if a.head == a.tail {
		a.head = &Node{
			value: v,
		}
		a.head.next = a.tail
		a.tail.prev = a.head
	} else {
		temp := a.head
		a.head = &Node{
			value: v,
		}
		a.head.next = temp
		temp.prev = a.head
	}
	a.size++
}

func (a *Queue) isEmpty() bool {
	return a.size == 0
}

func (a *Queue) Pop() (int, bool) {
	if a.isEmpty() {
		return -1, false
	} else if a.head == a.tail {
		p := a.head.value
		a.head = nil
		a.tail = nil
		a.size--
		return p, true
	}
	popped := a.tail.value
	a.tail = a.tail.prev
	a.tail.next = nil
	a.size--
	return popped, true
}

func (a *Queue) GetAll() []int {
	result := make([]int, 0)

	for i := a.head; i != nil; i = i.next {
		result = append(result, i.value)
	}
	return result
}

func runQueue() {
	o := Queue{}
	o.Push(5)
	o.Push(-45)
	o.Push(432)
	o.Push(57)
	o.Push(86)
	o.Push(-12)
	fmt.Println(o.GetAll())
	fmt.Println("Size: ", o.size)
	fmt.Println(o.Pop())
	fmt.Println(o.GetAll())
	fmt.Println(o.Pop())
	fmt.Println(o.Pop())
	fmt.Println(o.Pop())
	fmt.Println(o.isEmpty())
	fmt.Println("Size: ", o.size)
	fmt.Println(o.Pop())
	fmt.Println(o.Pop())
	fmt.Println(o.Pop())
}
