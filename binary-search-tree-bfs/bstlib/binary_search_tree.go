package bstlib

import "container/list"

// SearchTreeData ...
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// // SearchTreeData ...
// type searchTreeData interface {
// 	Insert(int)
// 	MapString(func(int) string) []string
// 	MapInt(func(int) int) []int
// }

// Bst ...
func Bst(number int) *SearchTreeData {
	return &SearchTreeData{
		left:  nil,
		data:  number,
		right: nil,
	}
}

// Insert ...
func (s *SearchTreeData) Insert(number int) {
	// loop through all the branch from the left to the right
	if number <= s.data {
		if s.left == nil {
			s.left = &SearchTreeData{nil, number, nil}
			return
		}
		s.left.Insert(number)
	} else if number >= s.data {
		if s.right == nil {
			s.right = &SearchTreeData{nil, number, nil}
			return
		}
		s.right.Insert(number)
	}
}

// MapString ...
func (s *SearchTreeData) MapString(f func(int) string) []string {
	var res []string
	if s.left != nil {
		res = append(res, s.left.MapString(f)...)
	}

	res = append(res, f(s.data))

	if s.right != nil {
		res = append(res, s.right.MapString(f)...)
	}

	return res
}

// MapInt ...
func (s *SearchTreeData) MapInt(f func(int) int) []int {
	var res []int
	if s.left != nil {
		res = append(res, s.left.MapInt(f)...)
	}

	res = append(res, f(s.data))

	if s.right != nil {
		res = append(res, s.right.MapInt(f)...)
	}

	return res
}

func checkSimple() []int {
	var res []int

	return res
}

// MapStringBst ...
func (s *SearchTreeData) MapStringBst() []int {
	visited := make(map[*SearchTreeData]bool)
	tocheck := list.New()

	var res []int

	tocheck.PushBack(s)

	for tocheck.Len() > 0 {
		e := tocheck.Front()
		s := e.Value.(*SearchTreeData)
		tocheck.Remove(e)

		if s != nil {
			if visited[s] == false {
				res = append(res, s.data)
				visited[s] = true
			}
		}

		if s.left != nil {
			if visited[s.left] == false {
				res = append(res, s.left.data)
				tocheck.PushBack(s.left)
				visited[s.left] = true
			}
		}

		if s.right != nil {
			if visited[s.right] == false {
				res = append(res, s.right.data)
				tocheck.PushBack(s.right)
				visited[s.right] = true
			}
		}
	}

	return res
}
