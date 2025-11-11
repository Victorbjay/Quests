package piscine

// Apply function f to every node in the list
func ListForEach(l *List, f func(*NodeL)) {
	if l == nil {
		return
	}

	current := l.Head
	for current != nil {
		f(current)
		current = current.Next
	}
}

// Function to add 2 (to int) or "2" (to string)
func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

// Function to subtract 3 (from int) or append "-3" (to string)
func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
