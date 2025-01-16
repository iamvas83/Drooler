package treebasics

func PreOrderTraversal(A *treeNode) []int {
	if A == nil {
		return nil
	}

	left := PreOrderTraversal(A.left)
	right := PreOrderTraversal(A.right)

	B := make([]int, 0)

	B = append(B, left...)
	B = append(B, A.data)
	B = append(B, right...)

	return B
}
