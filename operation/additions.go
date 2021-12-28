package operation

import (
	"errors"
	"strings"
)

// sum of two natural int
// input: string
// output: DLL
func IntAdd(a, b string) (*DLL, *OperationError) {
	l := DLL{}

	if len(a) < len(b) {
		a, b = b, a
	}

	if (a[0] == 48 && len(a) > 1) || (b[0] == 48 && len(b) > 1) {
		return &l, &OperationError{
			errors.New("InputError"),
			"First char of the input shouldn't be 0",
		}
	}

	b = strings.Repeat("0", len(a)-len(b)) + b

	carry := 0
	for i := len(a) - 1; i >= 0; i-- {
		a, b := int(a[i]-48), int(b[i]-48)
		sum := a + b + carry
		if sum > 9 {
			sum = sum % 10
			carry = 1
		} else {
			carry = 0
		}
		if err := l.Push(sum); err != nil {
			return &l, &OperationError{
				err.Error,
				err.Message,
			}
		}
	}
	if carry == 1 {
		if err := l.Push(1); err != nil {
			return &l, &OperationError{
				err.Error,
				err.Message,
			}
		}
	}

	return &l, nil
}

// sum of two natural int
// input: DLL
// output: DLL
func DLLIntAdd(a, b *DLL) (*DLL, *OperationError) {
 ptrA, ptrB, carry, l := a.Head, b.Head, 0, DLL{}    
 for {
     sum := ptrA.Value.(int) + ptrB.Value.(int) + carry
     if sum > 9 {
         sum = sum % 10
         carry = 1
     } else {
         carry = 0
     }
     if err := l.Push(sum); err != nil {
         return &l, &OperationError{
             err.Error,
             err.Message,
         }
     }
     ptrA = ptrA.Next
     ptrB = ptrB.Next
     if ptrA == nil {
         return &l, nil
     }
 }
}
