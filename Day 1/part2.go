package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func part2() int {
	f, err := os.Open(".input.txt")
	if err != nil {
		log.Fatalf("Error reading input file! %s", err.Error())
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	var t []int

	for s.Scan() {
		d, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatalf("Error converting string to int! failed object: %v\n%s", s.Text(), err.Error())
		}
		t = append(t, d)
	}

	for _, i := range t {
		for _, j := range t {
			for _, k := range t {
				if i+j+k == 2020 {
					return i * j * k
				}
			}
		}
	}

	return -1
}