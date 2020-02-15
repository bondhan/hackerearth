package main

import (
	"fmt"
	"strconv"

	b "github.com/bondhan/hackerearth/binary-search-tree-bfs/bstlib"
)

func main() {

	data := []int{11, 1, 3, 6, 2, 8, 4, 7, 10, 9, 0, 11}

	bst := b.Bst(5)
	for _, val := range data {
		bst.Insert(val)
	}

	res := bst.MapString(strconv.Itoa)
	fmt.Println(res) //[0 1 2 3 4 5 6 7 8 9 10 11 11]

	resBst := bst.MapStringBst()
	fmt.Println(resBst) //[5 1 11 0 3 6 2 4 8 7 10 9 11]

	fmt.Println("is valid bst:", bst.IsValid(b.MIN_NUM, b.MAX_NUM))

}
