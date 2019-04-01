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
