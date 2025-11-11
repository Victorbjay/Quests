package piscine

func ListReverse(l *List) {
	if l == nil || l.Head == nil || l.Head.Next == nil {
		return // empty or single element → nothing to reverse
	}

	// Reverse the linked list using three pointers
	var prev *NodeL = nil
	current := l.Head
	next := current.Next

	for current != nil {
		next = current.Next // save next
		current.Next = prev // reverse the link
		prev = current      // move prev forward
		current = next      // move current forward
	}

	// After reversal:
	// - prev is now the new head (last node originally)
	// - original head becomes the new tail
	l.Tail = l.Head
	l.Head = prev
}
