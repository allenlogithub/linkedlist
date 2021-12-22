package linkedlist

import (
	"errors"
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

// add a node before a node
func (l *DoublyCircularLinkedList) InsertBefore(n *DCNode, v interface{}) *DoublyCircularLinkedListError {
	if l.Length == 0 {
		return &DoublyCircularLinkedListError{
			errors.New("IndexError"),
			"Length should larger than 0",
		}
	} else if n == nil {
		return &DoublyCircularLinkedListError{
			errors.New("TypeError"),
			"Node can't be nil",
		}
	}
	newNode := DCNode{}
	newNode.Value = v
	newNode.Next = n
	newNode.Pre = n.Pre
	n.Pre.Next = &newNode
	n.Pre = &newNode

	return nil
}

func (l *DoublyCircularLinkedList) InsertAfter(n *DCNode, v interface{}) *DoublyCircularLinkedListError {
	if l.Length == 0 {
		return &DoublyCircularLinkedListError{
			errors.New("IndexError"),
			"Length should larger than 0",
		}
	} else if n == nil {
		return &DoublyCircularLinkedListError{
			errors.New("TypeError"),
			"Node can't be nil",
		}
	}
	newNode := DCNode{}
	newNode.Value = v
	newNode.Next = n.Next
	newNode.Pre = n
	n.Next.Pre = &newNode
	n.Next = &newNode	

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
