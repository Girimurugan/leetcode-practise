package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTStack struct {
	items []*TreeNode
}

// Push Function

func (s *BSTStack) push(value *TreeNode) {
	s.items = append(s.items, value)
}

// Pop Function
func (s *BSTStack) pop() *TreeNode {

	length := len(s.items) - 1

	if length >= 0 {
		popedValued := s.items[length]
		s.items = s.items[:length]

		return popedValued
	} else {
		return nil
	}
}
func kthSmallest(root *TreeNode, k int) int {

	s := &BSTStack{}
	for true {
		for root != nil {

			s.push(root)
			root = root.Left

		}
		root = s.pop()
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}

	return -1
}

func main() {

}
