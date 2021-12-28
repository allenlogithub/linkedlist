package operation

import (
	"errors"

	"linkedlist"
)

type (
	DLL struct {
		linkedlist.DoublyLinkedList
	}

	OperationError struct {
		Error   error
		Message string
	}
)

func (l *DLL) out() (string, *OperationError) {
	s, ptr := []byte{}, l.Tail
	for {
		if ptr == nil {
			return string(s), nil
		}
		if val, ok := ptr.Value.(byte); ok {
			s = append(s, val)
		} else {
			return "", &OperationError{
				errors.New("TypeError"),
				"Type should be byte",
			}
		}
		ptr = ptr.Pre
	}
}
