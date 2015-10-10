package add_two_numbers_002

import (
	"fmt"
)

type linkedNode struct {
	value int
	next  *linkedNode
}

type linkedList struct {
	first *linkedNode
}

func (l *linkedList) printList() {

	node := l.first

	for node != nil {
		fmt.Println(node)

		node = node.next
	}

	fmt.Println("=============================")

	return
}

// https://leetcode.com/problems/add-two-numbers/
func AddTwoNumbers(listA, listB linkedList) linkedList {

	var listC linkedList = linkedList{&linkedNode{-1, nil}}

	// the pointer for list C
	var currentC *linkedNode = listC.first

	nodeA := listA.first

	nodeB := listB.first

	vCNext := 0

	for nodeA != nil || nodeB != nil {

		vA := 0
		vB := 0

		if nodeA != nil {
			vA = nodeA.value
		}

		if nodeB != nil {
			vB = nodeB.value
		}

		valueC := (vA + vB + vCNext) % 10

		vCNext = (vA + vB + vCNext) / 10

		if listC.first.value == -1 {
			(*currentC).value = valueC
		} else {
			temp := &linkedNode{value: valueC, next: nil}
			(*currentC).next = temp

			currentC = temp
		}

		nodeA = nodeA.next

		nodeB = nodeB.next

	}

	return listC
}
