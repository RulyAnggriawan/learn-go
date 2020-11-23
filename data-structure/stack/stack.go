package stack

import . "example.com/node"

type Stack struct {
	values []Node
}

func (s *Stack) Add(newNode Node) {
	s.values = append(s.values, newNode)
}

func (s *Stack) Size() int {
	return len(s.values)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Peek() Node {
	return s.values[s.Size()-1]
}

func (s *Stack) Pop() (item Node) {
	stackLength := len(s.values)
	item = s.values[stackLength-1]
	s.values = s.values[:stackLength-1]
	return
}

//New return stack
func New() Stack {
	var stack Stack
	return stack
}

// func main() {
// 	fmt.Println("stack test")
// 	var stack Stack
// 	fmt.Printf("stack => %#v \n", stack)

// 	stack.add(4)
// 	stack.add(6)
// 	stack.add(2)
// 	fmt.Printf("stack => %#v \n", stack)

// 	x := stack.pop()
// 	fmt.Printf("stack => %#v \n", stack)
// 	fmt.Println("item :", x)

// 	stack.add(300)
// 	fmt.Printf("stack => %#v \n", stack)
// 	fmt.Printf("size => %#v \n", stack.size())
// 	fmt.Printf("isEmpty => %#v \n", stack.isEmpty())

// 	stack.pop()
// 	stack.pop()
// 	stack.pop()
// 	fmt.Printf("isEmpty => %#v \n", stack.isEmpty())

// }
