package operation

// subtraction of two int, a - b
// input: DLL
// output: DLL
func DLLIntSubtract(a, b *DLL) (*DLL, *OperationError) {
	// return of (a, b) are positive after the following condictions
	// if a > 0 and b < 0: return a+b => DLLIntAdd(a+b)
	// if a > 0 and b > 0: return a-b
	// if a < 0 and b < 0: return b-a => -(a-b)
	// if a < 0 and b > 0: return -(a+b) => DLLIntAdd(a+b) * (-1)

	// get isNeg information
	isNegA, isNegB, isNegGlobal := false, false, false
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
			return l, nil
		} else {
			return nil, &OperationError{
				err.Error,
				err.Message,
			}
		}
	} else if !isNegA && isNegB {
		if l, err := DLLIntAdd(a, b); err == nil {
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
		a, b, isNegGlobal = b, a, true
	} else if a.Length == b.Length {
		ptrA, ptrB := a.Head, b.Head
		for {
			if ptrA == nil {
				break
			} else if ptrA.Value.(int) < ptrB.Value.(int) {
				a, b, isNegGlobal = b, a, true
				break
			}
			ptrA, ptrB = ptrA.Next, ptrB.Next
		}
	}

	// minus
	ptrA, ptrB, l := a.Head, b.Head, DLL{}
	for {
		if ptrB == nil {
			break
		}
		sum := ptrA.Value.(int) - ptrB.Value.(int)
		if sum < 0 {
			if ptrA.Next == nil {
				break
			}
			ptrA.Next.Value = ptrA.Next.Value.(int) - 1
			sum += 10
		}
		l.Push(sum)
		ptrA, ptrB = ptrA.Next, ptrB.Next
	}

	// push the element of DLL.a if ptrA is not the last node
	if ptrA != nil {
		for {
			if ptrA == nil {
				break
			}
			l.Push(ptrA.Value)
			ptrA = ptrA.Next
		}
	}

	// Pop the node with value == 0, eg, 0001 -> 1
	if l.Tail.Value.(int) == 0 {
		ptr := l.Tail
		for {
			if ptr.Value.(int) != 0 || ptr == nil {
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
			l.Push(45)
		}
	}

	return &l, nil
}
