package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	// If position is negative, return nil
	if pos < 0 {
		return nil
	}

	current := l
	index := 0

	for current != nil {
		if index == pos {
			return current
		}
		index++
		current = current.Next
	}

	// If we reach here, pos is out of bounds
	return nil
}
