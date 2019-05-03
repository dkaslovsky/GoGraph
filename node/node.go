package node

import (
	"errors"
	"sync"
)

// Node is a vertex of a graph
type Node string

// Stack is a LIFO of nodes
type Stack struct {
	nodes []Node
	lock  sync.Mutex
}

// NewStack returns a pointer to an empty Stack
func NewStack() *Stack {
	return &Stack{}
}

// Push adds a node to the stack
func (s *Stack) Push(n Node) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.nodes = append(s.nodes, n)
}

// Pop removes and returns the most recently added node from the stack
func (s *Stack) Pop() (n Node, err error) {
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
func (s *Stack) Len() int {
	return len(s.nodes)
}

// Set is an unordered unique collection of nodes
type Set struct {
	items map[Node]struct{}
	lock  sync.Mutex
}

// NewSet returns a pointer to an empty Set
func NewSet() *Set {
	return &Set{
		items: map[Node]struct{}{},
	}
}

// Add adds a node to the set
func (s *Set) Add(elem Node) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items[elem] = struct{}{}
}

// Contains returns a bool indicating if the set contains a specified node
func (s *Set) Contains(elem Node) bool {
	_, ok := s.items[elem]
	return ok
}

// Len returns the number of nodes in the set
func (s *Set) Len() int {
	return len(s.items)
}

// ToSlice returns a slice of all nodes in the set
func (s *Set) ToSlice() []Node {
	sl := make([]Node, s.Len())
	i := 0
	for elem := range s.items {
		sl[i] = elem
		i++
	}
	return sl
}
