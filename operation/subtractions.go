package operation

// subtraction of two int, a - b
// input: *DLL, *DLL
// output: *DLL, *OperationError
func DLLIntSubtract(a, b *DLL) (*DLL, *OperationError) {
	// return of (a, b) are positive after the following condictions
	// if a > 0 and b < 0: return a+b => DLLIntAdd(a+b)
	// if a > 0 and b > 0: return a-b
	// if a < 0 and b < 0: return b-a => -(a-b)
	// if a < 0 and b > 0: return -(a+b) => DLLIntAdd(a+b) * (-1)

	// get isNeg and hasSwitched information
	isNegA, isNegB, isNegGlobal, hasSwitched := false, false, false, false
	if a.Tail.Value.(int) == 45 {
		isNegA = true
		a.Remove(a.Tail)
	}
	if b.Tail.Value.(int) == 45 {
		isNegB = true
		b.Remove(b.Tail)
	}

	// use addition method if (a, b) are both positive/ negative
	if isNegA && !isNegB {
		if l, err := DLLIntAdd(a, b); err == nil {
			l.Push(45)
			a.Push(45)
			return l, nil
		} else {
			return nil, &OperationError{
				err.Error,
				err.Message,
			}
		}
	} else if !isNegA && isNegB {
		if l, err := DLLIntAdd(a, b); err == nil {
			b.Push(45)
			return l, nil
		} else {
			return nil, &OperationError{
				err.Error,
				err.Message,
			}
		}
	}

	// check absolute value of (a, b), if |a| < |b|, then swap()
	if a.Length < b.Length {
		a, b, isNegGlobal, hasSwitched = b, a, true, true
	} else if a.Length == b.Length {
		ptrA, ptrB := a.Tail, b.Tail
		for {
			if ptrA == nil {
				break
			} else if ptrA.Value.(int) > ptrB.Value.(int) {
				break
			} else if ptrA.Value.(int) < ptrB.Value.(int) {
				a, b, isNegGlobal, hasSwitched = b, a, true, true
				break
			} else { // ptrA.Value.(int) == ptrB.Value.(int)
				ptrA, ptrB = ptrA.Pre, ptrB.Pre
			}
		}
	}

	// minus
	ptrA, ptrB, l, borrow := a.Head, b.Head, DLL{}, 0
	for {
		if ptrB == nil {
			break
		}
		sum := ptrA.Value.(int) - ptrB.Value.(int) + borrow
		borrow = 0 // reset
		if sum < 0 {
			if ptrA.Next == nil {
				break
			}
			borrow -= 1 // borrow one from the next digit
			sum += 10
		}
		l.Push(sum)
		ptrA, ptrB = ptrA.Next, ptrB.Next
	}

	// push the element of DLL.a if ptrA is not the last node (ptrA != nil)
	// need to check whether we have borrow (borrow != 0) as well
	if ptrA != nil {
		for {
			if ptrA == nil {
				break
			}
			sum := ptrA.Value.(int) + borrow
			if sum >= 0 {
				borrow = 0
			}
			l.Push(sum)
			ptrA = ptrA.Next
		}
	}

	// Pop the node with value == 0, eg, 0001 -> 1
	if l.Tail.Value.(int) == 0 {
		ptr := l.Tail
		for {
			if ptr == nil || ptr.Value.(int) != 0 {
				break
			} else if ptr.Value.(int) == 0 {
				l.Remove(ptr)
			}
			ptr = ptr.Pre
		}
	}

	// push minus symbol (-) if required
	if !isNegA && !isNegB {
		if isNegGlobal {
			l.Push(45)
		}
	} else if isNegA && isNegB {
		if !isNegGlobal {
			if l.Length != 0 { // eg. -103 - (-103)
				l.Push(45)
			}
		}
	}

	if hasSwitched {
		a, b = b, a
	}
	if isNegA {
		if err := a.Push(45); err != nil {
			return nil, &OperationError{
				err.Error,
				err.Message,
			}
		}
	}
	if isNegB {
		if err := b.Push(45); err != nil {
			return nil, &OperationError{
				err.Error,
				err.Message,
			}
		}
	}

	return &l, nil
}
