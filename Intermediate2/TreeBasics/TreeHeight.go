package treebasics

func TreeHeight(A *treeNode) int {
	if A == nil {
		return 0
	}

	return 1 + max(TreeHeight(A.left), TreeHeight(A.right))
}
