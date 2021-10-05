package main

// Definition for a binary tree node.
 type TreeNode struct {
     Val int
    Left *TreeNode
    Right *TreeNode
 }
 
 func maxDepth(root *TreeNode) int {
    
	if root == nil{
		return 0
	}

	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)

	if leftHeight > rightHeight{
		return leftHeight+1
	}else{
		return	rightHeight+1
	} 


}

func main(){

}