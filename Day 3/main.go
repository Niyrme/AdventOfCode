package main

import (
	"bufio"
	"log"
	"os"
)

type slope struct {
	x, y int
}

var (
	f     *os.File
	s     *bufio.Scanner
	lines []string
)

func main() {
	var (
		slopes []slope = []slope{
			{x: 1, y: 1},
			{x: 3, y: 1},
			{x: 5, y: 1},
			{x: 7, y: 1},
			{x: 1, y: 2},
		}

		slopeResultProduct int = 1
	)

	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("Error opening input.txt!\n%s", err.Error())
	}
	defer f.Close()
	s = bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	for _, slope := range slopes {
		trees := countTrees(slope)
		slopeResultProduct *= trees
		log.Printf("Found %v trees in slope (%v)", trees, slope)
	}

	log.Printf("slopeResultProduct is: %v", slopeResultProduct)
}

func countTrees(slo slope) int {
	var (
		treeEncounters int = 0
		posX, posY     int = 0, 0
	)

	for true {
		posX += slo.x
		posY += slo.y

		if posY >= len(lines) {
			break
		}

		posX %= len(lines[posY])

		if lines[posY][posX] == '#' {
			treeEncounters++
		}
	}

	return treeEncounters
}
