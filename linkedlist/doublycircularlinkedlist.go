package linkedlist

import (
	// "errors"
	"fmt"
)

type (
	DCNode struct {
		Value interface{}
		Next *DCNode
		Pre *DCNode
	}

	DoublyCircularLinkedList struct {
		Head *DCNode
		Length int
	}

	DoublyCircularLinkedListError struct {
		Error error
		Message string
	}
)

// add a node at the end
func (l *DoublyCircularLinkedList) Push(v interface{}) *DoublyCircularLinkedListError {
	n := DCNode{}
	n.Value = v
	if l.Length == 0 {
		n.Next = &n
		n.Pre = &n
		l.Head = &n
		l.Length ++
		return nil
	}
	ptr := l.Head.Pre
	n.Next = l.Head
	n.Pre = ptr
	ptr.Next = &n
	l.Head.Pre = &n
	l.Length++

	return nil
}

func (l *DoublyCircularLinkedList) Show(fromHead bool, steps int) {
	fmt.Println("DoublyCircularLinkedList length:", l.Length)
	var data []interface{}
	count := 0

	if fromHead {
		node := l.Head
		for {
			if count == steps {
				fmt.Println("Data:", data)
				return
			}

			data = append(data, node.Value)
			node = node.Next
			count++
		}
	}
	node := l.Head.Pre
	for {
		if count == steps {
			fmt.Println("Data:", data)
			return
		}
		data = append(data, node.Value)
		node = node.Pre
		count++
	}
}
