package node

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupStack() *Stack {
	itemX := &stackItem{data: "x"}
	itemY := &stackItem{data: "y", next: itemX}
	itemZ := &stackItem{data: "z", next: itemY}
	return &Stack{
		lock: &sync.Mutex{},
		last: itemZ,
		len:  3,
	}
}

func TestNewStack(t *testing.T) {
	t.Run("new Stack is empty", func(t *testing.T) {
		n := NewStack()
		assert.Nil(t, n.last)
		assert.Zero(t, n.len)
	})
}

func TestStackPush(t *testing.T) {
	tests := map[string]struct {
		stack  *Stack
		toPush Node
	}{
		"push to empty stack": {
			stack:  NewStack(),
			toPush: "a",
		},
		"push to nonempty stack": {
			stack:  setupStack(),
			toPush: "a",
		},
		"push to stack already containing same element": {
			stack:  setupStack(),
			toPush: "x",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			curStackLen := test.stack.len
			test.stack.Push(test.toPush)
			assert.Equal(t, test.stack.last.data, test.toPush)
			assert.Equal(t, test.stack.len, curStackLen+1)
		})
	}
}

func TestStackPop(t *testing.T) {
	tests := map[string]struct {
		stack     *Stack
		shouldErr bool
	}{
		"pop from empty stack should error": {
			stack:     NewStack(),
			shouldErr: true,
		},
		"pop from nonempty stack": {
			stack:     setupStack(),
			shouldErr: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			curStackLen := test.stack.len
			n, err := test.stack.Pop()
			if test.shouldErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, n, Node("z"))
			assert.Equal(t, test.stack.len, curStackLen-1)
		})
	}
}

func TestStackLen(t *testing.T) {
	tests := map[string]struct {
		stack    *Stack
		stackLen int
	}{
		"empty stack": {
			stack:    NewStack(),
			stackLen: 0,
		},
		"nonempty stack": {
			stack:    setupStack(),
			stackLen: 3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.stack.Len(), test.stackLen)
		})
	}
}

func setupSet() *Set {
	return &Set{
		items: map[Node]struct{}{
			"x": struct{}{},
			"y": struct{}{},
		},
		lock: &sync.Mutex{},
	}
}

func TestNewSet(t *testing.T) {
	t.Run("new Set is empty", func(t *testing.T) {
		n := NewSet()
		assert.Empty(t, n.items)
	})
}

func TestSetAdd(t *testing.T) {
	tests := map[string]struct {
		set   *Set
		toAdd Node
	}{
		"add to empty set": {
			set:   NewSet(),
			toAdd: "a",
		},
		"add to nonempty set": {
			set:   setupSet(),
			toAdd: "a",
		},
		"upsert": {
			set:   setupSet(),
			toAdd: "x",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.set.Add(test.toAdd)
			assert.Contains(t, test.set.items, test.toAdd)
		})
	}
}

func TestSetContains(t *testing.T) {
	tests := map[string]struct {
		element       Node
		shouldContain bool
	}{
		"set contains element": {
			element:       "x",
			shouldContain: true,
		},
		"set does not contain element": {
			element:       "a",
			shouldContain: false,
		},
	}

	set := setupSet()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			isIn := set.Contains(test.element)
			assert.Equal(t, isIn, test.shouldContain)
		})
	}
}

func TestSetLen(t *testing.T) {
	tests := map[string]struct {
		set    *Set
		setLen int
	}{
		"empty set": {
			set:    NewSet(),
			setLen: 0,
		},
		"nonempty set": {
			set:    setupSet(),
			setLen: 2,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.set.Len(), test.setLen)
		})
	}
}

func TestSetToSlice(t *testing.T) {
	tests := map[string]struct {
		set   *Set
		slice []Node
	}{
		"empty set": {
			set:   NewSet(),
			slice: []Node{},
		},
		"nonempty set": {
			set:   setupSet(),
			slice: []Node{"x", "y"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			setAsSlice := test.set.ToSlice()
			assert.ElementsMatch(t, setAsSlice, test.slice)
		})
	}
}
