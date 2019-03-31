package graph

import (
	"bufio"
	"strconv"
	"strings"
)

type reader interface {
	Read([]byte) (int, error)
}

func (g *Graph) addFromReader(r reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}

		from := parts[0]
		to := parts[1]
		if from == "" || to == "" {
			continue
		}

		if len(parts) == 2 {
			g.AddEdge(from, to)
		} else {
			weight, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				return err
			}
			g.AddEdge(from, to, weight)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (dg *DirGraph) addFromReader(r reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}

		from := parts[0]
		to := parts[1]
		if from == "" || to == "" {
			continue
		}

		if len(parts) == 2 {
			dg.AddEdge(from, to)
		} else {
			weight, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				return err
			}
			dg.AddEdge(from, to, weight)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
