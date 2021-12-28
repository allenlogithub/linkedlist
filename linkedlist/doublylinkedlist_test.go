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
	for i := 0; i < l.Length; i++ {
		values = append(values, n.Value)
		n = n.Next
	}

	return values
}

func (l *DoublyLinkedList) revTraverse() []interface{} {
	var values []interface{}
	n := l.Tail
	for i := l.Length - 1; i >= 0; i-- {
		values = append(values, n.Value)
		n = n.Pre
	}

	return values
}

func (l *DoublyLinkedList) getLength() int {
	return l.Length
}

func (l *DoublyLinkedList) checkDLL(t *testing.T, forward []interface{}, backward []interface{}, Length int, funcName string) {
	if !reflect.DeepEqual(l.traverse(), forward) {
		t.Errorf("DLL.%s failed", funcName)
	}
	if !reflect.DeepEqual(l.revTraverse(), backward) {
		t.Errorf("DLL.%s failed", funcName)
	}
	if l.getLength() != Length {
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

	// pos == l.Length check
	l.PushAt(2, 2)
	l.checkDLL(t, []interface{}{0, 1, 2}, []interface{}{2, 1, 0}, 3, "PushAt")

	// pos > l.Length/2 check
	l.PushAt(1, 0.5)
	l.checkDLL(t, []interface{}{0, 0.5, 1, 2}, []interface{}{2, 1, 0.5, 0}, 4, "PushAt")

	// pos <= l.Length/2 check
	l.PushAt(3, 1.5)
	l.checkDLL(t, []interface{}{0, 0.5, 1, 1.5, 2}, []interface{}{2, 1.5, 1, 0.5, 0}, 5, "PushAt")
}

func TestDLLPop(t *testing.T) {
	l := DoublyLinkedList{}

	// error check
	if err := l.Pop(1); err != nil {
		if err.Error.Error() != "IndexError" {
			t.Errorf("DLL.Pop failed")
		}
	}

	// error check
	l.Push(0)
	if err := l.Pop(1); err != nil {
		if err.Error.Error() != "ElementNotFound" {
			t.Errorf("DLL.Pop failed")
		}
	}

	// error check
	l.Push(1)
	l.Push(2)
	if err := l.Pop(3); err != nil {
		if err.Error.Error() != "ElementNotFound" {
			t.Errorf("DLL.Pop failed")
		}
	}

	// l.Head.value == v
	l.Pop(0)
	l.checkDLL(t, []interface{}{1, 2}, []interface{}{2, 1}, 2, "Pop")

	// *Node.value == v
	l.Pop(2)
	l.checkDLL(t, []interface{}{1}, []interface{}{1}, 1, "Pop")
}

func TestDLLPopAt(t *testing.T) {
	l := DoublyLinkedList{}

	// l.Length == 0
	if err := l.PopAt(1); err != nil {
		if err.Error.Error() != "IndexError" {
			t.Errorf("DLL.PopAt failed")
		}
	}

	// pos >- l.Length
	l.Push(0)
	if err := l.PopAt(1); err != nil {
		if err.Error.Error() != "IndexError" {
			t.Errorf("DLL.PopAt failed")
		}
	}

	// pos == 0
	l.Push(1)
	l.PopAt(0)
	l.checkDLL(t, []interface{}{1}, []interface{}{1}, 1, "PopAt")

	// pos*2 > l.Length
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Push(5)
	l.PopAt(3)
	l.checkDLL(t, []interface{}{1, 2, 3, 5}, []interface{}{5, 3, 2, 1}, 4, "PopAt")

	// pos*2 <= l.Length
	l.PopAt(1)
	l.checkDLL(t, []interface{}{1, 3, 5}, []interface{}{5, 3, 1}, 3, "PopAt")

	// pos == l.Length - 1
	l.PopAt(2)
	l.checkDLL(t, []interface{}{1, 3}, []interface{}{3, 1}, 2, "PopAt")
}

func TestDLLInsertBefore(t *testing.T) {
	l := DoublyLinkedList{}

	// error check
	if err := l.InsertBefore(l.Head, 2); err != nil {
		if err.Error.Error() != "IndexError" {
			t.Errorf("DLL.InsertBefore failed")
		}
	}

	// error check
	l.Push(0)
	if err := l.InsertBefore(l.Head.Pre, 2); err != nil {
		if err.Error.Error() != "TypeError" {
			t.Errorf("DLL.InsertAfter failed")
		}
	}

	// pos != 0 check
	l.Push(1)
	l.InsertBefore(l.Head.Next, 2)
	l.checkDLL(t, []interface{}{0, 2, 1}, []interface{}{1, 2, 0}, 3, "InsertBefore")

	// pos == 0 check
	l.InsertBefore(l.Head, 3)
	l.checkDLL(t, []interface{}{3, 0, 2, 1}, []interface{}{1, 2, 0, 3}, 4, "InsertBefore")
}

func TestDLLInsertAfter(t *testing.T) {
	l := DoublyLinkedList{}

	// error check
	if err := l.InsertAfter(l.Head, 2); err != nil {
		if err.Error.Error() != "IndexError" {
			t.Errorf("DLL.InsertAfter failed")
		}
	}

	// error check
	l.Push(0)
	if err := l.InsertAfter(l.Head, 2); err != nil {
		if err.Error.Error() != "TypeError" {
			t.Errorf("DLL.InsertAfter failed")
		}
	}

	// pos != 0 check
	l.Push(1)
	l.InsertAfter(l.Head.Next, 3)
	l.checkDLL(t, []interface{}{0, 2, 3, 1}, []interface{}{1, 3, 2, 0}, 4, "InsertAfter")

	// pos == Length - 1
	l.InsertAfter(l.Head.Next.Next.Next, 4)
	l.checkDLL(t, []interface{}{0, 2, 3, 1, 4}, []interface{}{4, 1, 3, 2, 0}, 5, "InsertAfter")
}

func TestDLLRemove(t *testing.T) {
	l := DoublyLinkedList{}
	var emptyDoublyLinkedList []interface{}

	// error check
	if err := l.Remove(l.Head); err != nil {
		if err.Error.Error() != "ElementNotFound" {
			t.Errorf("DLL.Remove failed")
		}
	}

	// check
	l.Push(1)
	l.Remove(l.Head)
	l.checkDLL(t, emptyDoublyLinkedList, emptyDoublyLinkedList, 0, "Remove")
	l.Push(0)
	l.Push(1)
	l.Remove(l.Head.Next)
	l.checkDLL(t, []interface{}{0}, []interface{}{0}, 1, "Remove")
	l.Push(2)
	l.Remove(l.Head)
	l.checkDLL(t, []interface{}{2}, []interface{}{2}, 1, "Remove")
}
