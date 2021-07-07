package crdt

import (
	"fmt"
)

func ExampleTree() {
	fmt.Println("You are creating a CRDTree")
	sv := StringValue{Val: "string value"}
	node := Node{ID: "root", Parent: nil, Children: []*Node{}, Val: &sv}
	child := Node{ID: "child1", Parent: &node, Children: []*Node{}, Val: &sv}
	node.Children = append(node.Children, &child);
	tree := Tree{ID: "abc", Root: node}
	fmt.Println(tree.String())
	// Output:
  // You are creating a CRDTree
	// ID: abc, Structure: Node[ID: root, Value: Value: string value, Parent: none, Children: [Node[ID: child1, Value: Value: string value, Parent: root, Children: []]]]
}
