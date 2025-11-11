package piscine

func ListPushFront(l *List, data interface{}) {
	// Create the new node
	newNode := &NodeL{Data: data}

	// Link new node to current head
	newNode.Next = l.Head

	// Update Head to point to new node
	l.Head = newNode

	// Special case: if list was empty, also update Tail
	if l.Tail == nil {
		l.Tail = newNode
	}
}
