/*
	go clean -testcache && go test -cover linkedlist
	go test -v for verbose
*/

package linkedlist

import (
	"testing"
	"reflect"
)

func (l *LinkedList) traverse() []interface{} {
	var values []interface{}
	n := l.head
	for i := 0; i < l.length; i++ {
		values = append(values, n.value)
		n = n.next
	}

	return values
}

func TestPush(t *testing.T) {
	l := LinkedList{}

	l.Push(1)
	l.Push("a")
	l.Push(1.0)
	if !reflect.DeepEqual(l.traverse(), []interface{}{1, "a", 1.0}) {
		t.Errorf("linkedlist.Push failed")
	}
}

func TestPop(t *testing.T) {
	l := LinkedList{}

	err := l.Pop(1)
	if !(err.Error.Error() == "IndexError") {
		t.Errorf("linkedlist.Pop failed")
	}

	l.Push(1)
	l.Push(2)
	l.Push(3)
	err = l.Pop(4)
	if !(err.Error.Error() == "ElementNotFound") {
		t.Errorf("linkedlist.Pop failed")
	}

	l.Pop(2)
	if !reflect.DeepEqual(l.traverse(), []interface{}{1, 3}) {
		t.Errorf("linkedlist.Push failed")
	}

	l.Pop(1)
	l.Pop(3)
	var empty []interface{}
	if !reflect.DeepEqual(l.traverse(), empty) {
		t.Errorf("linkedlist.Pop failed")
	}
}

func TestPushAt(t *testing.T) {
	l := LinkedList{}

	err := l.PushAt(1, 2)	
	if !(err.Error.Error() == "IndexError") {
		t.Errorf("linkedlist.PushAt failed")
	}

	l.Push(1)
	l.Push(3)
	l.PushAt(1, 2)
	l.PushAt(3, 4)
	if !reflect.DeepEqual(l.traverse(), []interface{}{1, 2, 3, 4}) {
		t.Errorf("linkedlist.PushAt failed")
	}	
}

func TestPopAt(t *testing.T) {
	l := LinkedList{}

	err := l.PopAt(0)
	if !(err.Error.Error() == "IndexError") {
		t.Errorf("linkedlist.PopAt failed")
	}

	l.Push(1)
	err = l.PopAt(1)
	if !(err.Error.Error() == "IndexError") {
		t.Errorf("linkedlist.PopAt failed")
	}
	l.Push(2)
	l.Push(3)
	l.PopAt(0)
	l.PopAt(1)
	if !reflect.DeepEqual(l.traverse(), []interface{}{2}) {
		t.Errorf("linkedlist.PopAt failed")
	}

	l.PopAt(0)
	var empty []interface{}
	if !reflect.DeepEqual(l.traverse(), empty) {
		t.Errorf("linkedlist.PopAt failed")
	}
}
