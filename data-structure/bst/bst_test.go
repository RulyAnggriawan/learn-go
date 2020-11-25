package bst

import "testing"

func TestPrint(t *testing.T) {
	// example bst
	/*
				6
			3			7
		1		4	5		9
						8

	*/

	/*
		expected output
		===========================

		           |---[ 1 ]
		     |---[ 3 ]
		     |     |     |---[ 4 ]
		     |     |---[ 5 ]
		---[ 6 ]
		     |---[ 7 ]
		           |     |---[ 8 ]
		           |---[ 9 ]

		===========================
	*/

	bst := New()
	bst.Add(6)
	bst.Add(3)
	bst.Add(7)
	bst.Add(5)
	bst.Add(9)
	bst.Add(1)
	bst.Add(8)
	bst.Add(4)

	bst.Print()

	t.Fatalf("Print Tree fail")
}

func TestTraverseOnOrder(t *testing.T) {
	// example bst
	/*
				6
			3		7
		1		5		9

		expected output 1, 3,5, 6, 7, 9
	*/
	bst := New()
	bst.Add(6)
	bst.Add(3)
	bst.Add(7)
	bst.Add(5)
	bst.Add(9)
	bst.Add(1)

	result := bst.TraverseInOrder()
	expectedResult := []int{1, 3, 5, 6, 7, 9}

	if len(result) != len(expectedResult) {
		t.Fatalf("traverse on order fail, result and expected result length not match")
	}

	// t.Fatalf("traverse on order fail")
}
