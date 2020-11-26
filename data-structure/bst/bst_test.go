package bst

import (
	"testing"
)

var bst Bst

func TestPrint(t *testing.T) {
	/*

		===========================


		            |---[ 9 ]
		            |     |---[ 8 ]
		      |---[ 7 ]
		|---[ 6 ]
		                  |---[ 5 ]
		      |     |---[ 4 ]
		            |     |---[ 3 ]
		      |---[ 2 ]
		            |---[ 1 ]


		===========================


	*/

	bst = New()
	bst.Add(6)
	bst.Add(2)
	bst.Add(7)
	bst.Add(4)
	bst.Add(9)
	bst.Add(1)
	bst.Add(8)
	bst.Add(3)
	bst.Add(5)

	bst.Print()
	//t.Fatalf("Print Tree fail")
}

func TestTraverseOnOrder(t *testing.T) {
	/*
		expected output 1,2,3,4,5,6,7,8,9
	*/

	result := bst.TraverseInOrder()
	expectedResult := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if len(result) != len(expectedResult) {
		t.Fatalf("result and expected result length not match. got %v, expected %v", len(result), len(expectedResult))
	}

	// t.Fatalf("traverse on order fail")
	if !isSameSlice(result, expectedResult) {
		t.Fatalf("result and expected slice not match.\n got %#v, expected %#v", result, expectedResult)
	}
}

func isSameSlice(firstSlice, secondSlice []int) bool {
	for i := range firstSlice {
		if firstSlice[i] != secondSlice[i] {
			return false
		}
	}
	return true
}
