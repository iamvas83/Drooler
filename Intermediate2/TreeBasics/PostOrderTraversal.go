package treebasics

func PostOrderTarversal(A *treeNode) []int {
	if A == nil {
		return nil
	}

	left := PostOrderTarversal(A.left)
	right := PostOrderTarversal(A.right)

	B := make([]int, 0)

	B = append(B, left...)
	B = append(B, right...)
	B = append(B, A.data)
	return B
}
