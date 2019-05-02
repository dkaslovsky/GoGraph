package node

import (
	"errors"
	"sync"
)

// Node is a node of a graph
type Node string

// NodeStack is a stack (LIFO) of nodes
type NodeStack struct {
	nodes []Node
	lock  sync.Mutex
}

// NewNodeStack returns a pointer to an empty NodeStack
func NewNodeStack() *NodeStack {
	return &NodeStack{}
}

// Push adds a node to the stack
func (s *NodeStack) Push(n Node) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.nodes = append(s.nodes, n)
}

// Pop removes and returns the most recently added node from the stack
func (s *NodeStack) Pop() (n Node, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	sLen := len(s.nodes)
	if sLen == 0 {
		return n, errors.New("cannot pop from empty stack")
	}

	n = s.nodes[sLen-1]
	s.nodes[sLen-1] = "" // prevent memory grow
	s.nodes = s.nodes[:sLen-1]
	return n, nil
}

// Len returns the number of nodes in the stack
func (s *NodeStack) Len() int {
	return len(s.nodes)
}

// NodeSet is a set of nodes
type NodeSet struct {
	set  map[Node]struct{}
	lock sync.Mutex
}

// NewNodeSet returns a pointer to an empty NodeSet
func NewNodeSet() *NodeSet {
	return &NodeSet{
		set: map[Node]struct{}{},
	}
}

// Add adds a node to the set
func (s *NodeSet) Add(elem Node) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.set[elem] = struct{}{}
}

// Contains returns a bool indicating if the set contains a specified node
func (s *NodeSet) Contains(elem Node) bool {
	_, ok := s.set[elem]
	return ok
}

// Len returns the number of nodes in the set
func (s *NodeSet) Len() int {
	return len(s.set)
}

// ToSlice returns a slice of all nodes in the set
func (s *NodeSet) ToSlice() []Node {
	sl := make([]Node, s.Len())
	i := 0
	for elem := range s.set {
		sl[i] = elem
		i++
	}
	return sl
}
