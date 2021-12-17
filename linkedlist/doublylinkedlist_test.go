/*
	https://stackoverflow.com/questions/16935965/how-to-run-test-cases-in-a-specified-file
	go clean -testcache && go test -v linkedlist
	go clean -testcache && go test -v linkedlist/doublylinkedlist_test.go linkedlist/doublylinkedlist.go
*/

package linkedlist

import (
	"testing"
	"reflect"
	// "fmt"
	// "errors"
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

func TestDLLPushAt(t *testing.T) {
	l := DoublyLinkedList{}

	// error check
	if err := l.PushAt(1, 3); err != nil {		
		if err.Error.Error() != "IndexError" {
			t.Errorf("DLL.PushAt failed")
		}
	}

	// error check
	l.Push(1)
	if err := l.PushAt(2, 3); err != nil {		
		if err.Error.Error() != "IndexError" {
			t.Errorf("DLL.PushAt failed")
		}
	}

	// pos == 0 check
	l.PushAt(0, 0)
	if !reflect.DeepEqual(l.traverse(), []interface{}{0, 1}) {
		t.Errorf("DLL.PushAt failed")
	}
	if !reflect.DeepEqual(l.revTraverse(), []interface{}{1, 0}) {
		t.Errorf("DLL.PushAt failed")
	}
	if l.getLength() != 2 {
		t.Errorf("DLL.PushAt failed")
	}

	// pos == l.length check
	l.PushAt(2, 2)
	if !reflect.DeepEqual(l.traverse(), []interface{}{0, 1, 2}) {
		t.Errorf("DLL.PushAt failed")
	}
	if !reflect.DeepEqual(l.revTraverse(), []interface{}{2, 1, 0}) {
		t.Errorf("DLL.PushAt failed")
	}
	if l.getLength() != 3 {
		t.Errorf("DLL.PushAt failed")
	}

	// pos > l.length/2 check
	l.PushAt(1, 0.5)
	if !reflect.DeepEqual(l.traverse(), []interface{}{0, 0.5, 1, 2}) {
		t.Errorf("DLL.PushAt failed")
	}
	if !reflect.DeepEqual(l.revTraverse(), []interface{}{2, 1, 0.5, 0}) {
		t.Errorf("DLL.PushAt failed")
	}
	if l.getLength() != 4 {
		t.Errorf("DLL.PushAt failed")
	}

	// pos <= l.length/2 check
	l.PushAt(3, 1.5)
	if !reflect.DeepEqual(l.traverse(), []interface{}{0, 0.5, 1, 1.5, 2}) {
		t.Errorf("DLL.PushAt failed")
	}
	if !reflect.DeepEqual(l.revTraverse(), []interface{}{2, 1.5, 1, 0.5, 0}) {
		t.Errorf("DLL.PushAt failed")
	}
	if l.getLength() != 5 {
		t.Errorf("DLL.PushAt failed")
	}
}
