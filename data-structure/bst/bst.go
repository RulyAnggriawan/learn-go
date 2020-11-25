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
	fmt.Println("===========================")
	printNode(t.root, 0)
	fmt.Println("===========================")
}

func printNode(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		// if level > 0 {
		// 	format += "|---"
		// }
		level++
		printNode(n.Left, level)
		fmt.Printf(format+"%d\n", n.Value)
		printNode(n.Right, level)
	}

}
