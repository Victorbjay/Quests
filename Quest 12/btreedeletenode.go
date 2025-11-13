package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node != nil {
		if node.Left == nil {
			root = BTreeTransplant(root, node, node.Right)
		} else if node.Right == nil {
			root = BTreeTransplant(root, node, node.Left)
		} else {
			y := BTreeMin(node.Right)
			if y != nil && y.Parent != node {
				root = BTreeTransplant(root, y, y.Right)

				y.Right = node.Right
				y.Right.Parent = y
			}
			root = BTreeTransplant(root, node, y)
			y.Left = node.Left
			y.Left.Parent = y
		}
	}
	return root
}
