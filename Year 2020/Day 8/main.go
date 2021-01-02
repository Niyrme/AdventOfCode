package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	d, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Error reading input file!\n%s", err.Error())
	}

	lines := strings.Split(string(d), "\n")

	var (
		acc, line int          = 0, 0
		visited   map[int]bool = map[int]bool{}
	)

	for true {
		if line == len(lines)-1 {
			break
		}

		content := strings.Split(lines[line], " ")

		instruction := strings.TrimSpace(content[0])
		valueOperator := strings.TrimSpace(content[1])

		if visited[line] {
			break
		} else {
			visited[line] = true
		}

		if instruction == "acc" {
			value, err := strconv.Atoi(valueOperator[1:])
			if err != nil {
				log.Fatalf("Error converting string to int: %v\n%s", valueOperator[1:], err.Error())
			}

			switch valueOperator[0] {
			case '+':
				acc += value
			case '-':
				acc -= value
			default:
				log.Fatalf("Wrong operator: %v (line %v)", valueOperator[0], line+1)
			}
		} else if instruction == "jmp" {
			value, err := strconv.Atoi(valueOperator[1:])
			if err != nil {
				log.Fatalf("Error converting string to int: %v\n%s", valueOperator[1:], err.Error())
			}
			switch valueOperator[0] {
			case '+':
				line += value
			case '-':
				line -= value
			default:
				log.Fatalf("Wrong operator: %v (line %v)", string(valueOperator[0]), line+1)
			}

			continue
		} else if instruction == "nop" {
			// Just do nothing
		} else {
			log.Fatalf("Wrong instruction found on line %v: %s", line, instruction)
		}

		line++
	}

	log.Printf("acc value *before* entering the loop: %v", acc)
}
