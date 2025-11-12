package piscine 

func ListMerge(l1 *List, l2 *List) {
	// If l1 is nil or empty, just point l1 to l2
	if l1 == nil {
		return
	}
	if l2 == nil || l2.Head == nil {
		return
	}
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}

	// Link the tail of l1 to the head of l2
	l1.Tail.Next = l2.Head

	// Update l1's Tail to be l2's Tail
	l1.Tail = l2.Tail

	// Optional: clear l2 to avoid double usage (good practice)
	l2.Head = nil
	l2.Tail = nil
}
