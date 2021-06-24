package internal

// HeadSuffix splits a text between its start and a suffix of length 1 character
func HeadSuffix(s string) (head, suffix string) {
	var (
		l = len(s)
	)

	switch l {
	case 0:
		return "", ""
	case 1:
		return "", s
	}

	head = s[:l-1]
	suffix = s[l-1:]

	return head, suffix
}
