package node

import (
	"errors"
	"sync"
)

// Node is a node of a graph
type Node string

type stackItem struct {
	data Node
	next *stackItem
}

// Stack is a LIFO of nodes
type Stack struct {
	lock *sync.Mutex
	top  *stackItem
	len  int
}

// NewStack returns a pointer to an empty Stack
func NewStack() *Stack {
	return &Stack{lock: &sync.Mutex{}}
}

// Push adds a node to the stack
func (s *Stack) Push(node Node) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.len++
	toPush := &stackItem{data: node}
	if s.top == nil {
		s.top = toPush
		return
	}
	toPush.next = s.top
	s.top = toPush
}

// Pop removes and returns the most recently added node from the stack
func (s *Stack) Pop() (Node, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.top == nil {
		return "", errors.New("cannot pop from empty stack")
	}

	s.len--
	curTop := s.top
	val := curTop.data
	s.top = curTop.next
	// prevent memory grow (likely not needed due to GC)
	curTop = nil

	return val, nil
}

// Len returns the number of nodes in the stack
func (s *Stack) Len() int {
	return s.len
}

// Set is an unordered unique collection of nodes
type Set struct {
	lock  *sync.Mutex
	items map[Node]struct{}
}

// NewSet returns a pointer to an empty Set
func NewSet() *Set {
	return &Set{
		items: map[Node]struct{}{},
		lock:  &sync.Mutex{},
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
