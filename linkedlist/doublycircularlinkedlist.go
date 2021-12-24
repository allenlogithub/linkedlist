package linkedlist

import (
	"errors"
	"fmt"
)

type (
	DCNode struct {
		Value interface{}
		Next  *DCNode
		Pre   *DCNode
	}

	DoublyCircularLinkedList struct {
		Head   *DCNode
		Length int
	}

	DoublyCircularLinkedListError struct {
		Error   error
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
		l.Length++
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
	l.Length++

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
	l.Length++

	return nil
}

// remove a node by a given element value
func (l *DoublyCircularLinkedList) Pop(v interface{}) *DoublyCircularLinkedListError {
	ptr := l.Head
	for i := 0; i < l.Length; i++ {
		if ptr.Value == v {
			l.Remove(ptr)
			return nil
		}
		ptr = ptr.Next
	}

	return &DoublyCircularLinkedListError{
		errors.New("ElementNotFound"),
		"Node not found",
	}
}

// remove a node
func (l *DoublyCircularLinkedList) Remove(n *DCNode) *DoublyCircularLinkedListError {
	if n == nil {
		return &DoublyCircularLinkedListError{
			errors.New("ElementNotFound"),
			"Node shouldn't be nil",
		}
	}
	switch len := l.Length; len {
	case 0:
		return &DoublyCircularLinkedListError{
			errors.New("ElementNotFound"),
			"Length of DCLL is 0",
		}
	case 1:
		if l.Head == n {
			l.Head = nil
			l.Length--
			return nil
		}
		return &DoublyCircularLinkedListError{
			errors.New("ElementNotFound"),
			"Node not found",
		}
	default:
		if n == l.Head {
			l.Head = n.Next
		}
		n.Pre.Next = n.Next
		n.Next.Pre = n.Pre
		l.Length--
		return nil
	}
}

func (l *DoublyCircularLinkedList) Show(fromHead bool, steps int) {
	fmt.Println("DoublyCircularLinkedList length:", l.Length)
	var data []interface{}
	if l.Length == 0 {
		fmt.Println("Data:", data)
		return
	}
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
