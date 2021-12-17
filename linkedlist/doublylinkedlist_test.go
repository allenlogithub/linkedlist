/*
	go clean -testcache && go test -v linkedlist
	go clean -testcache && go test -v linkedlist/doublylinkedlist_test.go linkedlist/doublylinkedlist.go
*/

package linkedlist

import (
	"testing"
	"reflect"
)

func (l *DoublyLinkedList) traverse() []interface{} {
	var values []interface{}
	n := l.Head
	for i := 0; i < l.length; i++ {
		values = append(values, n.Value)
		n = n.Next
	}

	return values
}

func (l *DoublyLinkedList) revTraverse() []interface{} {
	var values []interface{}
	n := l.Tail
	for i := l.length - 1; i >= 0; i -- {
		values = append(values, n.Value)
		n = n.Pre
	}

	return values
}

func (l *DoublyLinkedList) getLength() int {
	return l.length
}

func TestDLLPush(t *testing.T) {
	l := DoublyLinkedList{}

	l.Push(1)
	l.Push("a")
	l.Push(1.0)
	if !reflect.DeepEqual(l.traverse(), []interface{}{1, "a", 1.0}) {
		t.Errorf("DLL.Push failed")
	}
	if !reflect.DeepEqual(l.revTraverse(), []interface{}{1.0, "a", 1}) {
		t.Errorf("DLL.Push failed")
	}
	if l.getLength() != 3 {
		t.Errorf("DLL.Push failed")
	}
}
