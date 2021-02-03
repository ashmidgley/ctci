package main

func isOneEditAway(a, b string) bool {
	aLen := len(a)
	bLen := len(b)

	if aLen+1 == bLen {
		// Insertion.
		return isInsertion(a, b)
	}
	if aLen == bLen+1 {
		// Removal.
		return isInsertion(b, a)
	}
	if aLen == bLen {
		// Replacement.
		return isReplacement(a, b)
	}
	return false
}

func isInsertion(a, b string) bool {
	extraHit := false
	for i, j := 0, 0; i < len(a) && j < len(b); {
		if a[i] != b[j] {
			if extraHit {
				return false
			}
			extraHit = true
			j++
		} else {
			i++
			j++
		}
	}
	return true
}

func isReplacement(a, b string) bool {
	replacementHit := false
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if replacementHit {
				return false
			}
			replacementHit = true
		}
	}
	return true
}
