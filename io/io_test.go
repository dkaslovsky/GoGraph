package io

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ea struct {
	from    []string
	to      []string
	weights []float64
}

func (e *ea) AddEdge(from string, to string, weight ...float64) {
	e.from = append(e.from, from)
	e.to = append(e.to, to)

	w := 1.0
	if len(weight) > 0 {
		w = weight[0]
	}
	e.weights = append(e.weights, w)
}

func TestLoadGraphFromFile(t *testing.T) {
	e := &ea{}
	LoadGraphFromFile("../graph.golden", e)

	expectedFrom := []string{"a", "a", "b", "c", "c", "c"}
	expectedTo := []string{"b", "c", "c", "a", "d", "d"}
	expectedWeights := []float64{1.5, 2.0, 3.3, 7.0, 1.0, 1.1}

	assert.Equal(t, expectedFrom, e.from)
	assert.Equal(t, expectedTo, e.to)
	assert.Equal(t, expectedWeights, e.weights)
}
