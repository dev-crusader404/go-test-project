package main

import "fmt"

type List[T comparable] struct {
	head *Nodes[T]
}

type Nodes[T comparable] struct {
	data T
	next *Nodes[T]
}

func (a *List[T]) Push(v T) {
	if a.head == nil {
		a.head = &Nodes[T]{data: v}
		// a.tail = a.head
	} else {
		currentNode := a.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = &Nodes[T]{data: v}
	}
}

func (a *List[T]) GetAll() []T {
	var res []T
	for i := a.head; i != nil; i = i.next {
		res = append(res, i.data)
	}
	return res
}

func (a *List[T]) delete(key T) {
	if a.head == nil {
		return
	}
	searchFrom := a.head
	prev := &Nodes[T]{}
	if searchFrom.data == key {
		a.head = searchFrom.next
		return
	} else {
		for searchFrom.next != nil && searchFrom.data != key {
			prev = searchFrom
			searchFrom = searchFrom.next
		}
		if searchFrom.data == key {
			prev.next = searchFrom.next
			searchFrom = nil
		} else {
			fmt.Printf("\nNot found: %+v\n", key)
		}
	}
}

func runLinkedList() {

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	lst.Push(-34)
	lst.Push(42)
	fmt.Println("list:", lst.GetAll())

	k := List[string]{}
	k.Push("ram")
	k.Push("sam")
	k.Push("zam")
	k.Push("tam")
	k.Push("ham")
	k.Push("pam")
	fmt.Println("list:", k.GetAll())
	k.delete("zam")
	fmt.Println("list:", k.GetAll())

	lst.delete(10)
	lst.delete(-34)
	lst.Push(12)
	fmt.Println("list:", lst.GetAll())
	lst.delete(12)
	fmt.Println("list:", lst.GetAll())
}
