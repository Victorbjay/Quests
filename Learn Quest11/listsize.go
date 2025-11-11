package piscine

func ListSize(l *List) int {
	if l == nil || l.Head == nil {
		return 0
	}

	count := 0
	current := l.Head

	for current != nil {
		count++
		current = current.Next
	}

	return count
}
