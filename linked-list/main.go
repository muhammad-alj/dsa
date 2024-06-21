package main

import "fmt"

type node[T any] struct {
	data T
	next *node[T]
}
type linkedList[T any] struct {
	head   *node[T]
	length int
}

func (l linkedList[T]) print() {
	iterator := l.head
	for iterator != nil {
		fmt.Println(iterator)
		iterator = iterator.next
	}
}

func (l *linkedList[T]) append(item T) {
	l.length++

	newNode := &node[T]{item, nil}

	if l.head == nil {
		l.head = newNode
		return
	}

	// Very criptic, but way cooler
	for iteratorPtr := l.head; iteratorPtr != nil; iteratorPtr = iteratorPtr.next {
		if iteratorPtr.next == nil {
			iteratorPtr.next = newNode
			break
		}
	}
}

func (l *linkedList[T]) prepend(item T) {
	l.length++
	newNode := &node[T]{item, nil}

	if l.head == nil {
		l.head = newNode
		return
	}

	newNode.next = l.head
	l.head = newNode
	l.length++
}

func main() {
	var people linkedList[string]

	people.prepend("mohamamd")
	people.prepend("abdo")
	people.prepend("some random ass nigger")
	people.print()
}
