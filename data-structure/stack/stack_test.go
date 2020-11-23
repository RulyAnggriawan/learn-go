package stack

import (
	"testing"

	. "example.com/node"
)

func TestNew(t *testing.T) {
	s := New()
	if !isStack(s) {
		t.Fatalf("new() not Stack")
	}

}

func isStack(i interface{}) bool {
	_, ok := i.(Stack)
	return ok
}

func TestPush(t *testing.T) {
	s := New()
	s.Add(Node{3, nil, nil})
	if s.Size() == 0 {
		t.Fatalf("push fail")
	}
}
