/*
	https://stackoverflow.com/questions/16935965/how-to-run-test-cases-in-a-specified-file
	go clean -testcache && go test -v linkedlist
	go clean -testcache && go test -v linkedlist/doublylinkedlist_test.go linkedlist/doublylinkedlist.go
*/

package linkedlist

import (
	"reflect"
	"testing"
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
	for i := l.length - 1; i >= 0; i-- {
		values = append(values, n.Value)
		n = n.Pre
	}

	return values
}

func (l *DoublyLinkedList) getLength() int {
	return l.length
}

func (l *DoublyLinkedList) checkDLL(t *testing.T, forward []interface{}, backward []interface{}, length int, funcName string) {
	if !reflect.DeepEqual(l.traverse(), forward) {
		t.Errorf("DLL.%s failed", funcName)
	}
	if !reflect.DeepEqual(l.revTraverse(), backward) {
		t.Errorf("DLL.%s failed", funcName)
	}
	if l.getLength() != length {
		t.Errorf("DLL.%s failed", funcName)
	}
}

func TestDLLPush(t *testing.T) {
	l := DoublyLinkedList{}

	l.Push(1)
	l.Push("a")
	l.Push(1.0)
	l.checkDLL(t, []interface{}{1, "a", 1.0}, []interface{}{1.0, "a", 1}, 3, "Push")
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
	l.checkDLL(t, []interface{}{0, 1}, []interface{}{1, 0}, 2, "PushAt")

	// pos == l.length check
	l.PushAt(2, 2)
	l.checkDLL(t, []interface{}{0, 1, 2}, []interface{}{2, 1, 0}, 3, "PushAt")

	// pos > l.length/2 check
	l.PushAt(1, 0.5)
	l.checkDLL(t, []interface{}{0, 0.5, 1, 2}, []interface{}{2, 1, 0.5, 0}, 4, "PushAt")

	// pos <= l.length/2 check
	l.PushAt(3, 1.5)
	l.checkDLL(t, []interface{}{0, 0.5, 1, 1.5, 2}, []interface{}{2, 1.5, 1, 0.5, 0}, 5, "PushAt")
}
