package piscine

func ListLast(l *List) interface{} {
	if l == nil || l.Head == nil {
		return nil
	}

	// If Tail is properly maintained, we can return it directly
	if l.Tail != nil {
		return l.Tail.Data
	}

	// Fallback: traverse from Head to find the last node
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	return current.Data
}
