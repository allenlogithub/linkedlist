package linkedlist

import (
	"errors"
	"fmt"
)

type (
	DNode struct {
		Value interface{}
		Next  *DNode
		Pre   *DNode
	}

	DoublyLinkedList struct {
		Head   *DNode
		Tail   *DNode
		length int
	}

	DoublyLinkedlistError struct {
		Error   error
		Message string
	}
)

// add a node at the Tail
func (l *DoublyLinkedList) Push(v interface{}) *DoublyLinkedlistError {
	n := DNode{}
	n.Value = v
	if l.length == 0 {
		l.Head = &n
		l.Tail = &n
		l.length++
		return nil
	}

	ptr := l.Tail
	n.Value = v
	n.Next = nil
	n.Pre = ptr
	ptr.Next = &n
	l.Tail = &n
	l.length++

	return nil
}

// add a node at the Tail
// adding a node from the Head or the Tail is based on pos (if pos > l.length/2: from Tail)
func (l *DoublyLinkedList) PushAt(pos int, v interface{}) *DoublyLinkedlistError {
	n := DNode{}
	n.Value = v
	if pos == l.length {
		l.Push(v)
		return nil
	} else if pos > l.length {
		return &DoublyLinkedlistError{
			errors.New("IndexError"),
			"Insert position should not larger than the last index",
		}
	} else if pos == 0 {
		ptr := l.Head
		n.Next = ptr
		n.Pre = nil
		ptr.Pre = &n
		l.Head = &n
		l.length++
		return nil
	} else if pos*2 > (l.length) {
		ptr := l.Tail
		for i := l.length - 1; i >= 0; i-- {
			if i == pos {
				n.Next = ptr
				n.Pre = ptr.Pre
				ptr.Pre = &n
				ptr.Pre.Pre.Next = &n
				l.length++
				return nil
			}
			ptr = ptr.Pre
		}
	} else if pos*2 <= l.length {
		ptr := l.Head.Next
		for i := 1; i < l.length; i++ {
			if i == pos {
				n.Next = ptr
				n.Pre = ptr.Pre
				ptr.Pre = &n
				ptr.Pre.Pre.Next = &n
				l.length++
				return nil
			}
			ptr = ptr.Next
		}
	}

	return &DoublyLinkedlistError{
		errors.New("Unknown"),
		"Sth wrong in linkedlist.PushAt",
	}
}

// pop the Value:v from DLL if v exists
func (l *DoublyLinkedList) Pop(v interface{}) *DoublyLinkedlistError {
	if l.length == 0 {
		return &DoublyLinkedlistError{
			errors.New("IndexError"),
			"Length should larger than 0",
		}
	}
	ptr := l.Head
	for {
		if ptr.Value == v {
			l.Remove(ptr)
			return nil
		}
		ptr = ptr.Next
		if ptr == nil {
			return &DoublyLinkedlistError{
				errors.New("ElementNotFound"),
				"Element not found",
			}
		}
	}

	return &DoublyLinkedlistError{
		errors.New("Unknown"),
		"Sth wrong in linkedlist.Pop",
	}
}

// pop an element at a given position
func (l *DoublyLinkedList) PopAt(pos int) *DoublyLinkedlistError {	
	if l.length == 0 {
		return &DoublyLinkedlistError{
			errors.New("IndexError"),
			"Length should larger than 0",
		}
	} else if pos >= l.length {
		return &DoublyLinkedlistError{
			errors.New("IndexError"),
			"Pop position should not larger than the last index",
		}
	}
	if pos*2 > l.length {
		ptr := l.Tail
		for i := l.length - 1; i >= 0; i-- {
			if i == pos {
				l.Remove(ptr)
				return nil
			}
			ptr = ptr.Pre
		}
	} else {
		ptr := l.Head
		for i := 0; i < l.length; i++ {
			if i == pos {
				l.Remove(ptr)
				return nil
			}
			ptr = ptr.Next
		}
	}

	return &DoublyLinkedlistError{
		errors.New("Unknown"),
		"Sth wrong in linkedlist.PopAt",
	}
}

// insert an element before a node
func (l *DoublyLinkedList) InsertBefore(n *DNode, v interface{}) *DoublyLinkedlistError {
	if l.length == 0 {
		return &DoublyLinkedlistError{
			errors.New("IndexError"),
			"Length should larger than 0",
		}
	}
	newNode := DNode{}
	newNode.Value = v
	if n.Pre == nil {
		newNode.Pre = nil
		newNode.Next = n
		n.Pre = &newNode
		l.Head = &newNode
		l.length++
		return nil
	}
	newNode.Pre = n.Pre
	newNode.Next = n
	n.Pre.Next = &newNode
	n.Pre = &newNode
	l.length++
	return nil
}

// insert an element after a node
func (l *DoublyLinkedList) InsertAfter(n *DNode, v interface{}) *DoublyLinkedlistError {
	if l.length == 0 {
		return &DoublyLinkedlistError{
			errors.New("IndexError"),
			"Length should larger than 0",
		}
	} else if n == nil {
		return &DoublyLinkedlistError{
			errors.New("TypeError"),
			"Node can't be nil",
		}
	}
	newNode := DNode{}
	newNode.Value = v
	if n.Next == nil {
		newNode.Next = nil
		newNode.Pre = n
		n.Next = &newNode
		l.Tail = &newNode
		l.length++
		return nil
	}
	newNode.Pre = n
	newNode.Next = n.Next
	n.Next.Pre = &newNode
	n.Next = &newNode
	l.length++
	return nil
}

// remove an element
// no RemoveBefore/ RemoveAfter required, just use Remove(node.pre)/ Remove(node.next)
func (l *DoublyLinkedList) Remove(n *DNode) *DoublyLinkedlistError {
	if n == nil {
		return &DoublyLinkedlistError{
			errors.New("ElementNotFound"),
			"Element shouldn't be nil",
		}
	}
	if n.Pre == nil {
		l.Head = n.Next
	} else {
		n.Pre.Next = n.Next
	}
	if n.Next == nil {
		l.Tail = n.Pre
	} else {
		n.Next.Pre = n.Pre
	}
	l.length--

	return nil
}

// show the data of the linked list
// if fromHead is true, will traverse from Head
// if fromHead is false, will traverse from Tail
func (l *DoublyLinkedList) Show(fromHead bool) {
	fmt.Println("DoublyLinkedList length:", l.length)
	var data []interface{}
	if fromHead {
		node := l.Head
		for {
			if node == nil {
				fmt.Println("Data:", data)
				return
			}

			data = append(data, node.Value)
			node = node.Next
		}
	}
	node := l.Tail
	for {
		if node == nil {
			fmt.Println("Data:", data)
			return
		}

		data = append(data, node.Value)
		node = node.Pre
	}
}
