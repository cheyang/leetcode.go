package maximum_depth_of_binary_tree_104

import (
	"math"
)

type BinaryTreeNode struct{
	
	value int
	left *BinaryTreeNode
	right *BinaryTreeNode
	
}


func MaxDepth(root *BinaryTreeNode) int{
	
	if root == nil ||(root.left == nil && root.right == nil){
		return 0
	}else{
		return math.Max(MaxDepth(root.left) + 1, MaxDepth(root.right) + 1 )
		
	}
}