package linkedlist

import (
	"errors"
	"fmt"
)

type (
	Node struct {
		value interface{}
		next  *Node
	}

	LinkedList struct {
		head   *Node
		length int
	}

	LinkedlistError struct {
		Error   error
		Message string
	}
)

// add a node with value:v
func (l *LinkedList) Push(v interface{}) *LinkedlistError {
	n := Node{}
	n.value = v
	if l.length == 0 {
		l.head = &n
		l.length++
		return nil
	}
	ptr := l.head
	for i := 0; i < l.length; i++ {
		if ptr.next != nil {
			ptr = ptr.next
		} else {
			n.value = v
			ptr.next = &n
			l.length++
			return nil
		}
	}
	return &LinkedlistError{
		errors.New("Unknown"),
		"Sth wrong in linkedlist.Push",
	}
}

// remove the first node with value:v
func (l *LinkedList) Pop(v interface{}) *LinkedlistError {
	prePtr, ptr := l.head, l.head
	if l.length == 0 {
		return &LinkedlistError{
			errors.New("IndexError"),
			"Length should larger than 0",
		}
	}
	if l.head.value == v {
		l.head = l.head.next
		l.length--
		return nil
	}
	for {
		switch val := ptr.value; val {
		case v:
			prePtr.next = ptr.next
			ptr = nil
			l.length--
			return nil
		default:
			if ptr.next == nil {
				return &LinkedlistError{
					errors.New("ElementNotFound"),
					"Element not found",
				}
			}
			prePtr, ptr = ptr, ptr.next
		}
	}
}

// add a node with value:v at given position
func (l *LinkedList) PushAt(pos int, v interface{}) *LinkedlistError {
	if pos == l.length {
		if err := l.Push(v); err != nil {
			return err
		}
		return nil
	} else if pos > l.length {
		return &LinkedlistError{
			errors.New("IndexError"),
			"Insert position should not larger than the last index",
		}
	}
	ptr, newNode := l.head, Node{}
	if pos == 0 {
		newNode.value = v
		newNode.next = ptr
		l.head = &newNode
		l.length++
		return nil
	}
	for i := 1; i < l.length; i++ {
		if i == pos {
			newNode.value = v
			newNode.next = ptr.next
			ptr.next = &newNode
			l.length++
			return nil
		}
	}

	return &LinkedlistError{
		errors.New("Unknown"),
		"Sth wrong in linkedlist.Push",
	}
}

// remove the element located at Position:pos
func (l *LinkedList) PopAt(pos int) *LinkedlistError {
	if l.length == 0 {
		return &LinkedlistError{
			errors.New("IndexError"),
			"Length should larger than 0",
		}
	}
	if pos > l.length-1 {
		return &LinkedlistError{
			errors.New("IndexError"),
			"Pop position should not larger than the last index",
		}
	}
	if pos == 0 {
		l.head = l.head.next
		l.length--
		return nil
	}
	prePtr, ptr := l.head, l.head.next
	for i := 1; i < l.length; i++ {
		if i == pos {
			prePtr.next = ptr.next
			l.length--
			return nil
		}
		prePtr, ptr = ptr, ptr.next
	}

	return &LinkedlistError{
		errors.New("Unknown"),
		"Sth wrong in linkedlist.PopAt",
	}
}

// show the data of the linked list
func (l *LinkedList) Show() {
	fmt.Println("LinkedList length:", l.length)
	var data []interface{}
	node := l.head
	for {
		if node == nil {
			fmt.Println("Data:", data)
			return
		}
		data = append(data, node.value)
		node = node.next
	}
}
