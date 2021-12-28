package operation

import (
	"errors"
	"strings"
)

func IntAdd(a, b string) (string, *OperationError) {
	l := DLL{}

	if len(a) < len(b) {
		a, b = b, a
	}

	if a[0] == 48 || b[0] == 48 {
		return "", &OperationError{
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
		if err := l.Push(byte(sum + 48)); err != nil {
			return "", &OperationError{
				err.Error,
				err.Message,
			}
		}
	}
	if carry == 1 {
		if err := l.Push(byte(49)); err != nil {
			return "", &OperationError{
				err.Error,
				err.Message,
			}
		}
	}

	if val, err := l.out(); err == nil {
		return val, nil
	} else {
		return "", err
	}
}
