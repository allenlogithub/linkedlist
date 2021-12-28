package operation

import (
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

// DLL.head.Value = 3, DLL.head.Next.Value = 2, DLL.head.Next.Next.Value = 1
// -> "123"
// input: DLL
// output: string
func (l *DLL) ToString() string {
	s, ptr := []byte{}, l.Tail
	for {
		if ptr == nil {
			return string(s)
		}
		s = append(s, byte(ptr.Value.(int)+48))
		ptr = ptr.Pre
	}
}

// "123"
// -> DLL.head.Value = 3, DLL.head.Next.Value = 2, DLL.head.Next.Next.Value = 1
func ToDLL(s string) (*DLL, *OperationError) {
	l := DLL{}
	for i := len(s) - 1; i >= 0; i-- {
		if err := l.Push(int(s[i] - 48)); err != nil {
			return nil, &OperationError{
				err.Error,
				err.Message,
			}
		}
	}

	return &l, nil
}
