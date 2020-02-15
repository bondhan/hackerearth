## Implementation for Binary Search Tree DFS and BFS

	`data := []int{11, 1, 3, 6, 2, 8, 4, 7, 10, 9, 0, 11}

	bst := b.Bst(5)
	for _, val := range data {
		bst.Insert(val)
	}

	res := bst.MapString(strconv.Itoa)
	fmt.Println(res)

	resBst := bst.MapStringBst()
	fmt.Println(resBst)`