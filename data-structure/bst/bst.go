package bst

import (
	"fmt"

	. "example.com/node"
)

type Bst struct {
	root *Node
}

func New() Bst {
	var bst Bst
	return bst
}

func (t *Bst) Add(value int) {
	if t.root == nil {
		t.root = &Node{value, nil, nil}
	} else {
		insertNode(t.root, value)
	}
}

func insertNode(n *Node, newValue int) {
	//insert left
	if n.Value > newValue {
		if n.Left == nil {
			n.Left = &Node{newValue, nil, nil}
		} else {
			insertNode(n.Left, newValue)
		}
	} else if n.Value < newValue { //insert right
		if n.Right == nil {
			n.Right = &Node{newValue, nil, nil}
		} else {
			insertNode(n.Right, newValue)
		}
	}
}

//TaverseInOrder return slice of item
func (t *Bst) TraverseInOrder() []int {
	var traverseItems []int
	if t.root != nil {
		traverseInOrder(t.root, &traverseItems)
	}
	return traverseItems
}

func traverseInOrder(n *Node, traverseItems *[]int) {
	if n != nil {
		traverseInOrder(n.Left, traverseItems)
		*traverseItems = append(*traverseItems, n.Value)
		traverseInOrder(n.Right, traverseItems)
	}
}

func (t *Bst) Print() {
	fmt.Printf("===========================\n\n\n")
	printNode(t.root, 0, 0, false)
	fmt.Printf("\n\n===========================\n")

}

// branchDir => branch direction
// 1 : right
// 0 : root
// -1 : left
func printNode(n *Node, level int, branchDir int, change bool) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			if change && (i == (level - 1)) {
				format += "|"
				format += "     "
			} else {
				format += "      "
			}

		}

		format += "|---"
		level++

		nodeRightDir, nodeLeftDir := 1, 1
		changeRight, changeLeft := true, true
		if branchDir == 1 {
			nodeRightDir = branchDir
			nodeLeftDir = branchDir * -1
			changeRight = false
			changeLeft = true
		} else if branchDir == -1 {
			nodeRightDir = branchDir * -1
			nodeLeftDir = branchDir
			changeRight = true
			changeLeft = false
		} else if branchDir == 0 {
			nodeRightDir = 1
			nodeLeftDir = -1
			changeRight = false
			changeLeft = false
		}

		printNode(n.Right, level, nodeRightDir, changeRight)
		fmt.Printf(format+"[ %d ]\n", n.Value)
		printNode(n.Left, level, nodeLeftDir, changeLeft)
	}

}
