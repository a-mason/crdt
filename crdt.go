package crdt

import (
	"fmt"
)

type StringValue struct {
	Val string
}

type DataType interface {
	String() (string)
}

type Node struct {
	ID string
	Parent *Node
	Children []*Node
	Val *StringValue
}

type Tree struct {
	ID string
	Root Node
}

func (sv *StringValue) String() (string) {
	return fmt.Sprintf("Value: %s", sv.Val);
}

func (sv *StringValue) Add(s string) (bool) {
	sv.Val += s
	return true;
}

func (vs *StringValue) Snapshot() (string) {
	return vs.Val;
}

func mapN(vs []*Node, f func(*Node) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func NString(n *Node) (string) {
	var ps string
	if n.Parent == nil {
		ps = "none"
	} else {
		ps = n.Parent.ID
	}
	return fmt.Sprintf("Node[ID: %s, Value: %s, Parent: %s, Children: %s]", n.ID, n.Val, ps, mapN(n.Children, NString));
}

func (t *Tree) String() (string) {
	return fmt.Sprintf("ID: %s, Structure: %s", t.ID, NString(&t.Root));
}

func InitTree() (*Tree) {
	fmt.Println("You are creating a CRDTree")
	sv := StringValue{"string value"}
	node := Node{"node1", nil, []*Node{}, &sv}
	tree := Tree{"abc", node}
	return &tree
}
