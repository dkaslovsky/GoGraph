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
	last *stackItem
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

	toPush := &stackItem{data: node}
	if s.last == nil {
		s.last = toPush
	} else {
		toPush.next = s.last
		s.last = toPush
	}
	s.len++
}

// Pop removes and returns the most recently added node from the stack
func (s *Stack) Pop() (Node, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.last == nil {
		return "", errors.New("cannot pop from empty stack")
	}

	curLast := s.last
	s.last = curLast.next
	s.len--

	val := curLast.data
	curLast = nil // prevent memory grow (likely not needed due to GC)
	return val, nil
}

// Len returns the number of nodes in the stack
func (s *Stack) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.len
}

type queueItem struct {
	data Node
	next *queueItem
}

// Queue is a FIFO of nodes
type Queue struct {
	lock  *sync.Mutex
	first *queueItem
	last  *queueItem
	len   int
}

// NewQueue creates an empty Queue
func NewQueue() *Queue {
	return &Queue{lock: &sync.Mutex{}}
}

// Push adds a node to the queue
func (q *Queue) Push(node Node) {
	q.lock.Lock()
	defer q.lock.Unlock()

	toPush := &queueItem{data: node}
	if q.last == nil {
		q.last = toPush
		q.first = toPush
	} else {
		q.last.next = toPush
		q.last = toPush
	}
	q.len++
}

// Pop removes the first node in the queue
func (q *Queue) Pop() (Node, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.first == nil {
		return "", errors.New("cannot pop from empty queue")
	}

	curFirst := q.first
	q.first = curFirst.next
	if q.first == nil {
		q.last = nil
	}
	q.len--

	val := curFirst.data
	curFirst = nil // prevent memory grow (likely not needed due to GC)
	return val, nil
}

// Len returns the number of nodes in the queue
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.len
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
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.items[elem]
	return ok
}

// Len returns the number of nodes in the set
func (s *Set) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return len(s.items)
}

// ToSlice returns a slice of all nodes in the set
func (s *Set) ToSlice() []Node {
	s.lock.Lock()
	defer s.lock.Unlock()

	sl := make([]Node, len(s.items))
	i := 0
	for elem := range s.items {
		sl[i] = elem
		i++
	}
	return sl
}
