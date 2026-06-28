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
