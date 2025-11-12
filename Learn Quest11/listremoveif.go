package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil {
		return
	}

	// Handle cases where Head needs to be removed
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	// If list is now empty, update Tail
	if l.Head == nil {
		l.Tail = nil
		return
	}

	// Traverse the list and remove matching nodes
	current := l.Head
	for current != nil && current.Next != nil {
		if current.Next.Data == data_ref {
			// Remove next node
			current.Next = current.Next.Next
			// If we removed the tail, update Tail
			if current.Next == nil {
				l.Tail = current
			}
		} else {
			current = current.Next
		}
	}
}
