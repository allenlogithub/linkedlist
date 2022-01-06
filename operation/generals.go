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
		if ptr.Value == 45 {
			s = append(s, byte(ptr.Value.(int)))
		} else {
			s = append(s, byte(ptr.Value.(int)+48))
		}
		ptr = ptr.Pre
	}
}

// "123"
// -> DLL.head.Value = 3, DLL.head.Next.Value = 2, DLL.head.Next.Next.Value = 1
func ToDLL(s string) (*DLL, *OperationError) {
	l := DLL{}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 45 {
			if err := l.Push(int(s[i])); err != nil {
				return nil, &OperationError{
					err.Error,
					err.Message,
				}
			}
		} else {
			if err := l.Push(int(s[i] - 48)); err != nil {
				return nil, &OperationError{
					err.Error,
					err.Message,
				}
			}
		}
	}

	return &l, nil
}

func CopyDLL(a *DLL) (*DLL, *OperationError) {
	l, ptr := DLL{}, a.Head
	for {
		if ptr == nil {
			return &l, nil
		}
		if err := l.Push(ptr.Value); err != nil {
			return &l, &OperationError{
				err.Error,
				err.Message,
			}
		}
		ptr = ptr.Next
	}
}
