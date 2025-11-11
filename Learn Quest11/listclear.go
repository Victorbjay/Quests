package piscine

func ListClear(l *List) {
	// Clear all nodes by breaking the chain and letting GC handle memory
	current := l.Head
	for current != nil {
		next := current.Next
		current.Next = nil  // help garbage collector
		current.Data = nil  // optional: clear data
		current = next
	}
	
	// Reset both Head and Tail to nil
	l.Head = nil
	l.Tail = nil
}
