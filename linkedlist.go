package linkedlist

import (
	"fmt"
)

type (
	Node struct {
		value int
		next *Node
	}

	LinkedList struct {
		head *Node
		length int
	}
)

// add a node with value:v
func (l *LinkedList) Push(v int) {
	n := Node{}
	n.value = v
	if l.length == 0 {
		l.head = &n
		l.length ++
		return
	}
	ptr := l.head
	for i := 0; i < l.length; i ++ {
		if ptr.next != nil {
			ptr = ptr.next
		} else {
			n.value = v
			ptr.next = &n
			l.length ++
			return
		}		
	}
}

// remove the first node with value:v
func (l *LinkedList) Pop(v int) {
	prePtr, ptr := l.head, l.head
	if l.length == 0 {
		fmt.Println("data length is 0")
		return
	}
	if l.head.value == v {
		l.head = l.head.next
		l.length --
		return
	}
	for {
		switch val := ptr.value; val {
		case v:
			prePtr.next = ptr.next
			ptr = nil
			l.length --
			return
		default:
			if ptr.next == nil {
				fmt.Printf("pop failed, data %d not found\n", v)
				return
			}
			prePtr, ptr = ptr, ptr.next
		}
	}
}

// add a node with value:v at given position
func (l *LinkedList) PushAt(pos, v int) {
	if pos == l.length {
		l.push(v)
		return
	} else if pos > l.length {
		fmt.Println("posistion not allowed")
		return
	}
	ptr, newNode := l.head, Node{}
	if pos == 0 {
		newNode.value = v
		newNode.next = ptr
		l.head = &newNode
		l.length ++
		return
	}
	for i := 1; i < l.length; i ++ {
		if i == pos {
			newNode.value = v
			newNode.next = ptr.next
			ptr.next = &newNode
			l.length ++
			return
		}
	}
}

// show the data of the linked list
func (l *LinkedList) Show() {
	fmt.Println("l.length:", l.length)
	var data []int
	node := l.head
	for {
		if node == nil {
			fmt.Println("data:", data)
			return
		}
		data = append(data, node.value)
		node = node.next
	}	
}
