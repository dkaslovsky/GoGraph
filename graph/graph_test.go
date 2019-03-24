package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestDirGraphAddEdge(t *testing.T) {
	dg := NewDirGraph("test")

	// test adding edge with default weight
	dg.AddEdge("x", "y")
	wgt, ok := dg.GetEdgeWeight("x", "y")
	assert.True(t, ok)
	assert.Equal(t, 1.0, wgt)
}

type DirGraphTestSuite struct {
	suite.Suite
	DG    *DirGraph
	Nodes []string
}

func (suite *DirGraphTestSuite) SetupTest() {
	suite.DG = NewDirGraph("test")
	suite.DG.AddEdge("a", "b", 1.5)
	suite.DG.AddEdge("a", "c", 2)
	suite.DG.AddEdge("b", "c", 3.3)
	suite.DG.AddEdge("c", "a", 7)
	suite.DG.AddEdge("c", "d", 1.1)
	suite.Nodes = []string{"a", "b", "c", "d"}
}

func (suite *DirGraphTestSuite) TestDirGraphGetNodes() {
	nodes := suite.DG.GetNodes()
	assert.Len(suite.T(), nodes, len(suite.Nodes))
	for _, node := range suite.Nodes {
		assert.Contains(suite.T(), nodes, node)
	}
}

type getNbrTest struct {
	node    string
	exists  bool
	expNbrs []string
}

var getOutNbrTest = []getNbrTest{
	{
		node:    "a",
		exists:  true,
		expNbrs: []string{"b", "c"},
	},
	{
		node:    "b",
		exists:  true,
		expNbrs: []string{"c"},
	},
	{
		node:    "c",
		exists:  true,
		expNbrs: []string{"a", "d"},
	},
	{
		node:   "d",
		exists: false,
	},
	{
		node:   "z",
		exists: false,
	},
}

func (suite *DirGraphTestSuite) TestDirGraphGetOutNeighbors() {
	for _, tt := range getOutNbrTest {
		nbrs, ok := suite.DG.GetOutNeighbors(tt.node)
		if !tt.exists {
			assert.False(suite.T(), ok)
			continue
		}
		assert.Len(suite.T(), nbrs, len(tt.expNbrs))
		for _, n := range tt.expNbrs {
			assert.Contains(suite.T(), nbrs, n)
		}
	}
}

var getInNbrTest = []getNbrTest{
	{
		node:    "a",
		exists:  true,
		expNbrs: []string{"c"},
	},
	{
		node:    "b",
		exists:  true,
		expNbrs: []string{"a"},
	},
	{
		node:    "c",
		exists:  true,
		expNbrs: []string{"a", "b"},
	},
	{
		node:    "d",
		exists:  true,
		expNbrs: []string{"c"},
	},
	{
		node:   "z",
		exists: false,
	},
}

func (suite *DirGraphTestSuite) TestDirGraphGetInNeighbors() {
	for _, tt := range getInNbrTest {
		nbrs, ok := suite.DG.GetInNeighbors(tt.node)
		if !tt.exists {
			assert.False(suite.T(), ok)
			continue
		}
		assert.Len(suite.T(), nbrs, len(tt.expNbrs))
		for _, n := range tt.expNbrs {
			assert.Contains(suite.T(), nbrs, n)
		}
	}
}

func TestDirGraphTestSuite(t *testing.T) {
	suite.Run(t, new(DirGraphTestSuite))
}
