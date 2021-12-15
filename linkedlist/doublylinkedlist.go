package linkedlist

import (
	"errors"
	"fmt"
)

type (
	DNode struct {
		value interface{}
		next  *DNode
		pre   *DNode
	}

	DoublyLinkedList struct {
		head   *DNode
		tail   *DNode
		length int
	}

	DoublyLinkedlistError struct {
		Error   error
		Message string
	}
)

// add a node at the tail
func (l *DoublyLinkedList) Push(v interface{}) *DoublyLinkedlistError {
	n := DNode{}
	n.value = v
	if l.length == 0 {
		l.head = &n
		l.tail = &n
		l.length++
		return nil
	}

	ptr := l.tail
	n.value = v
	n.next = nil
	n.pre = ptr
	ptr.next = &n
	l.tail = &n
	l.length++

	return nil
}

// add a node at the tail
// adding a node from the head or the tail is based on pos (if pos > l.length/2: from tail)
func (l *DoublyLinkedList) PushAt(pos int, v interface{}) *DoublyLinkedlistError {
	n := DNode{}
	n.value = v
	if pos == l.length {
		l.Push(v)
		return nil
	} else if pos > l.length {
		return &DoublyLinkedlistError{
			errors.New("IndexError"),
			"Insert position should not larger than the last index",
		}
	} else if pos == 0 {
		ptr := l.head
		n.next = ptr
		n.pre = nil
		ptr.pre = &n
		l.head = &n
		l.length++
		return nil
	} else if pos*2 > (l.length) {
		ptr := l.tail
		for i := l.length - 1; i >= 0; i-- {
			if i == pos {
				n.next = ptr
				n.pre = ptr.pre
				ptr.pre = &n
				ptr.pre.pre.next = &n
				l.length++
				return nil
			}
			ptr = ptr.pre
		}
	} else if pos*2 <= l.length {
		ptr := l.head.next
		for i := 1; i < l.length; i++ {
			if i == pos {
				n.next = ptr
				n.pre = ptr.pre
				ptr.pre = &n
				ptr.pre.pre.next = &n
				l.length++
				return nil
			}
			ptr = ptr.next
		}
	}

	return &DoublyLinkedlistError{
		errors.New("Unknown"),
		"Sth wrong in linkedlist.PushAt",
	}
}

// pop the value:v from DLL if v exists
func (l *DoublyLinkedList) Pop(v interface{}) *DoublyLinkedlistError {
	if l.length == 0 {
		return &DoublyLinkedlistError{
			errors.New("IndexError"),
			"Length should larger than 0",
		}
	}
	if l.head.value == v {
		l.head.next.pre = nil
		l.head = l.head.next
		l.length--
		return nil
	}
	ptr := l.head.next
	for i := 1; i < l.length; i++ {
		if ptr.value == v {
			if ptr.next == nil {
				ptr.pre.next = nil
				l.tail = ptr.pre
				l.length--
				return nil
			}
			ptr.pre.next = ptr.next
			ptr.next.pre = ptr.pre
			l.length--
			return nil
		}
		ptr = ptr.next
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
	}
	if pos >= l.length {
		return &DoublyLinkedlistError{
			errors.New("IndexError"),
			"Pop position should not larger than the last index",
		}
	} else if pos == 0 {
		l.head.next.pre = nil
		l.head = l.head.next
		l.length--
		return nil
	} else if pos == l.length-1 {
		ptr := l.tail
		ptr.pre.next = nil
		l.tail = ptr.pre
		l.length--
		return nil
	} else if pos*2 > l.length {
		ptr := l.tail.pre
		for i := l.length - 2; i > 0; i-- {
			if i == pos {
				ptr.next.pre = ptr.pre
				ptr.pre.next = ptr.next
				l.length--
				return nil
			}
			ptr = ptr.pre
		}
	} else if pos*2 <= l.length {
		ptr := l.head.next
		for i := 1; i < l.length; i++ {
			if i == pos {
				ptr.next.pre = ptr.pre
				ptr.pre.next = ptr.next
				l.length--
				return nil
			}
			ptr = ptr.next
		}
	}

	return &DoublyLinkedlistError{
		errors.New("Unknown"),
		"Sth wrong in linkedlist.PopAt",
	}
}

// show the data of the linked list
// if fromHead is true, will traverse from head
// if fromHead is false, will traverse from tail
func (l *DoublyLinkedList) Show(fromHead bool) {
	fmt.Println("DoublyLinkedList length:", l.length)
	var data []interface{}
	if fromHead {
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
	node := l.tail
	for {
		if node == nil {
			fmt.Println("Data:", data)
			return
		}

		data = append(data, node.value)
		node = node.pre
	}
}
