package construct_binary_tree_from_inorder_and_postorder_traversal

import (
	"fmt"
)

const (
	
)

type TreeNode struct {
	leftNode *TreeNode
	rightNode *TreeNode
	value int
}




func BuildTree(preOrder []int, inOrder []int)(root *TreeNode){
	
	if len(preOrder) != len(inOrder) {
		panic("The number "+len(preOder) +" of preorder "+preOrder+" doesn't match the number "+ len(inOrder) +" of "+inOrder)
	}
	
	if len(preOrder) == 0{
		return nil
	}
	
	orderMap := OrderMap(inOrder)
	
	ValueOfRoot := preOrder[0]
	
	pos,ok := ordermap[ValueOfRoot]
	
	
	
	if !ok {
		panic("Failed to find the root"+ValueOfRoot +" in perorder")
	}
	
	leftSubTree_inOrder := inOrder[0:pos]
	
	leftSubTree_preOrder := preOrder[1: 2+ len(leftSubTree_inOrder)]
	
	rightSubTree_inOrder := make([]int,0)
	
	if pos != len(inOrder) - 1{
		rightSubTree_inOrder = inOrder[pos+1:len(inOrder)]
	}
			
	rightSubTree_preOrder := make([]int,0)
	
	if 2+ len(leftSubTree_inOrder) < len(preOrder){
		rightSubTree_preOrder = preOrder[2+ len(leftSubTree_inOrder):len(preOrder)]
	}
	
	root = &TreeNode{leftNode: BuildTree(leftSubTree_preOrder, leftSubTree_inOrder), rightNode: BuildTree(rightSubTree_preOrder, rightSubTree_inOrder), value: ValueOfRoot}
	
	return root
}


/**Map the order as [value]i
*/
func OrderMap(order []int) (orderMap map[int]int){
	
	orderMap = make(map[int]int, len(order))
	
	
	for i, v := order {
		orderMap[v]=i
	}
	
	return orderMap
}