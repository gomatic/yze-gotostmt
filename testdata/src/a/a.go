package a

// withGoto uses a goto and must be flagged.
func withGoto() int {
	i := 0
loop:
	if i < 10 {
		i++
		goto loop // want `goto is not permitted`
	}
	return i
}

// withBreak uses break (a non-goto branch statement) and must NOT be flagged.
func withBreak() {
	for {
		break
	}
}

// withContinue uses continue (a non-goto branch statement) and must NOT be flagged.
func withContinue() {
	for i := 0; i < 3; i++ {
		continue
	}
}

// withFallthrough uses fallthrough (a non-goto branch statement) and must NOT be flagged.
func withFallthrough(n int) int {
	switch n {
	case 0:
		fallthrough
	case 1:
		return 1
	}
	return 0
}

// withLabeledBreak uses a labeled break (a non-goto branch statement) and must NOT be flagged.
func withLabeledBreak() {
loop:
	for {
		break loop
	}
}
