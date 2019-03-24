package io

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type edgeAdder interface {
	AddEdge(string, string, ...float64)
}

func ReadFromFile(filepath string, ea edgeAdder) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
			ea.AddEdge(from, to)
		} else {
			weight, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				return err
			}
			ea.AddEdge(from, to, weight)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
