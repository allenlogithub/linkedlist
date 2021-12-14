package linkedlist

import (
	// "errors"
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

func (l *DoublyLinkedList) Push(v interface{}) *LinkedlistError {
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
