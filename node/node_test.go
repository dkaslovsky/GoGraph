package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupNodeStack() *NodeStack {
	return &NodeStack{
		nodes: []Node{"x", "y", "z"},
	}
}

func TestNewNodeStack(t *testing.T) {
	t.Run("new NodeStack is empty", func(t *testing.T) {
		n := NewNodeStack()
		assert.Empty(t, n.nodes)
	})
}

func TestNodeStackPush(t *testing.T) {
	tests := map[string]struct {
		stack  *NodeStack
		toPush Node
	}{
		"push to empty stack": {
			stack:  NewNodeStack(),
			toPush: "a",
		},
		"push to nonempty stack": {
			stack:  setupNodeStack(),
			toPush: "a",
		},
		"push to stack already containing same element": {
			stack:  setupNodeStack(),
			toPush: "x",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			curStackLen := len(test.stack.nodes)
			test.stack.Push(test.toPush)
			assert.Contains(t, test.stack.nodes, test.toPush)
			assert.Equal(t, len(test.stack.nodes), curStackLen+1)
		})
	}
}

func TestNodeStackPop(t *testing.T) {
	tests := map[string]struct {
		stack     *NodeStack
		shouldErr bool
	}{
		"pop from empty stack should error": {
			stack:     NewNodeStack(),
			shouldErr: true,
		},
		"pop from nonempty stack": {
			stack:     setupNodeStack(),
			shouldErr: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			curStackLen := len(test.stack.nodes)
			n, err := test.stack.Pop()
			if test.shouldErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, n, Node("z"))
			assert.Equal(t, len(test.stack.nodes), curStackLen-1)
		})
	}
}

func TestNodeStackLen(t *testing.T) {
	tests := map[string]struct {
		stack    *NodeStack
		stackLen int
	}{
		"empty stack": {
			stack:    NewNodeStack(),
			stackLen: 0,
		},
		"nonempty stack": {
			stack:    setupNodeStack(),
			stackLen: 3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.stack.Len(), test.stackLen)
		})
	}
}

func setupNodeSet() *NodeSet {
	return &NodeSet{
		set: map[Node]struct{}{
			"x": struct{}{},
			"y": struct{}{},
		},
	}
}

func TestNewNodeSet(t *testing.T) {
	t.Run("new NodeSet is empty", func(t *testing.T) {
		n := NewNodeSet()
		assert.Empty(t, n.set)
	})
}

func TestNodeSetAdd(t *testing.T) {
	tests := map[string]struct {
		set   *NodeSet
		toAdd Node
	}{
		"add to empty set": {
			set:   NewNodeSet(),
			toAdd: "a",
		},
		"add to nonempty set": {
			set:   setupNodeSet(),
			toAdd: "a",
		},
		"upsert": {
			set:   setupNodeSet(),
			toAdd: "x",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.set.Add(test.toAdd)
			assert.Contains(t, test.set.set, test.toAdd)
		})
	}
}

func TestNodeSetContains(t *testing.T) {
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

	set := setupNodeSet()
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			isIn := set.Contains(test.element)
			assert.Equal(t, isIn, test.shouldContain)
		})
	}
}

func TestNodeSetLen(t *testing.T) {
	tests := map[string]struct {
		set    *NodeSet
		setLen int
	}{
		"empty set": {
			set:    NewNodeSet(),
			setLen: 0,
		},
		"nonempty set": {
			set:    setupNodeSet(),
			setLen: 2,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.set.Len(), test.setLen)
		})
	}
}

func TestNodeSetToSlice(t *testing.T) {
	tests := map[string]struct {
		set   *NodeSet
		slice []Node
	}{
		"empty set": {
			set:   NewNodeSet(),
			slice: []Node{},
		},
		"nonempty set": {
			set:   setupNodeSet(),
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
