package crdt

import (
	"strings"
)

type StringAtom struct {
	Val string
}

func (s StringAtom) Compare(s2 StringAtom) int {
	return strings.Compare(s.Val, s2.Val)
}

type Atom interface {
	Compare(a Atom) int
}
