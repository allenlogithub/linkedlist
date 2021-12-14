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
				prePtr := ptr.pre
				n.next = ptr
				n.pre = prePtr
				ptr.pre = &n
				prePtr.next = &n
				l.length++
				return nil
			}
			ptr = ptr.pre
		}
	} else if pos*2 <= (l.length) {
		ptr := l.head
		for i := 0; i < l.length; i++ {
			if i == pos {
				nxtPtr := ptr.next
				n.next = nxtPtr
				n.pre = ptr
				nxtPtr.pre = &n
				ptr.next = &n
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
