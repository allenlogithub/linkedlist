package operation

// multiplication of two int, classical method, big O squared
// input: string
// output: DLL
func DLLIntMultiply(a, b *DLL) (*DLL, *OperationError) {
	ptrA, l, m, carry, sum, isPositive := a.Head, DLL{}, make(map[int]int), 0, 0, 1

	if a.Tail.Value.(int) == 45 {
		a.Remove(a.Tail)
		isPositive ^= 0
	} else {
		isPositive ^= 1
	}
	if b.Tail.Value.(int) == 45 {
		b.Remove(b.Tail)
		isPositive ^= 0
	} else {
		isPositive ^= 1
	}

	for i := 0; i < a.Length; i++ {
		if ptrA != nil {
			ptrB := b.Head
			for j := 0; j < b.Length; j++ {
				if val, ok := m[i+j]; ok {
					m[i+j] = val + ptrB.Value.(int)*ptrA.Value.(int)
				} else {
					m[i+j] = ptrB.Value.(int) * ptrA.Value.(int)
				}
				ptrB = ptrB.Next
			}
		}
		ptrA = ptrA.Next
	}
	for i := 0; i < a.Length+b.Length; i++ {
		if i == 0 {
			l.Push(m[i] % 10)
		} else {
			sum = m[i]%10 + m[i-1]/10 + carry
			if sum >= 10 {
				carry = sum / 10
				sum = sum % 10
			} else {
				carry = 0
			}
			l.Push(sum)
		}
	}
	if sum == 0 {
		l.Remove(l.Tail)
	}

	if isPositive == 0 {
		l.Push(45)
	}

	return &l, nil
}
