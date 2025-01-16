package treebasics

func InorderTraversal(A *treeNode) []int {
	if A == nil {
		return nil
	}

	left := InorderTraversal(A.left)
	right := InorderTraversal(A.right)

	B := make([]int, 0)
	B = append(B, left...)
	B = append(B, A.data)
	B = append(B, right...)
	return B
}
