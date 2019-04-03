package graph

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func (g *Graph) addFromReader(r io.ReadCloser) error {
	defer r.Close()
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}

		src := parts[0]
		tgt := parts[1]
		if src == "" || tgt == "" {
			continue
		}

		if len(parts) == 2 {
			g.AddEdge(src, tgt)
			continue
		}

		weight, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return err
		}
		g.AddEdge(src, tgt, weight)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
