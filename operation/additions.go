package operation

// sum of two natural numbers
// input: DLL
// output: DLL
func DLLIntAdd(a, b *DLL) (*DLL, *OperationError) {
	hasSwitched := false
	if a.Length < b.Length {
		a, b, hasSwitched = b, a, true
	}
	ptrA, ptrB, carry, l := a.Head, b.Head, 0, DLL{}

	for i := 0; i < b.Length; i++ {
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
		ptrA, ptrB = ptrA.Next, ptrB.Next
	}
	if ptrA != nil {
		for {
			if ptrA == nil {
				break
			}
			sum := ptrA.Value.(int) + carry
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

	if hasSwitched {
		a, b = b, a
	}

	return &l, nil
}
