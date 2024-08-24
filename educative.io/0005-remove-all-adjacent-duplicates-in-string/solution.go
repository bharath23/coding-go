package educative0005

func removeDuplicatesV0(s string) string {
	stack := ""
	size := 0
	for i := range s {
		if size == 0 || stack[size-1] != s[i] {
			stack += string(s[i])
			size++
		} else {
			stack = stack[:size-1]
			size--
		}

	}
	return stack
}

func removeDuplicatesV1(s string) string {
	out := []byte(s)
	lidx := 1
	ridx := 1
	for ridx < len(s) {
		if lidx != 0 && out[ridx] == out[lidx-1] {
			lidx--
		} else {
			out[lidx] = out[ridx]
			lidx++
		}
		ridx++
	}

	return string(out[:lidx])
}
