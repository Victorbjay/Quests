package piscine

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int:
		return node.Data.(int) > 0
	case float32:
		return node.Data.(float32) > 0
	case float64:
		return node.Data.(float64) > 0
	case byte:
		return node.Data.(byte) > 0
	default:
		return false
	}
}

// Returns true if node.Data is NOT a number (i.e. alphanumeric/string)
func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

// Apply function f to a node only if cond(node) returns true
func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	if l == nil {
		return
	}

	current := l.Head
	for current != nil {
		if cond(current) {
			f(current)
		}
		current = current.Next
	}
}
