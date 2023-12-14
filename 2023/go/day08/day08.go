package day08

import (
	"log"
	"regexp"
	"strings"

	"github.com/kingkero/adventofcode/2023/go/util"
)

type Node struct {
	data      string
	left      *Node
	leftData  string
	right     *Node
	rightData string
	isEndNote bool
}

func (node *Node) step(direction string) *Node {
	if direction == "L" {
		return node.left
	}

	return node.right
}

func getMinimumSteps(instructions []string, nodes map[string]*Node) int {
	checkNode := nodes["AAA"]
	steps := 0
	for {
		checkNode = checkNode.step(instructions[steps%len(instructions)])

		if checkNode.isEndNote {
			return steps + 1
		}

		steps++
	}
}

func getRules()

func getMinimumStepsParallel(instructions []string, nodes map[string]*Node, startNodes []*Node) int {
	return 0
}

func part01(lines []string) int {
	rules := lines[2:]

	lineParser := regexp.MustCompile("([A-Z]{3}) = \\(([A-Z]{3}), ([A-Z]{3})\\)")

	nodes := make(map[string]*Node)

	for _, line := range rules {
		matches := lineParser.FindAllStringSubmatch(line, -1)

		value := matches[0][1]
		left := matches[0][2]
		right := matches[0][3]

		node := &Node{value, nil, left, nil, right, value == "ZZZ"}
		nodes[value] = node
	}

	for _, node := range nodes {
		node.left = nodes[node.leftData]
		node.right = nodes[node.rightData]
	}

	return getMinimumSteps(strings.Split(lines[0], ""), nodes)
}

func part02(lines []string) int {
	rules := lines[2:]

	lineParser := regexp.MustCompile("([A-Z]{3}) = \\(([A-Z]{3}), ([A-Z]{3})\\)")

	nodes := make(map[string]*Node)
	var startNodes []*Node

	for _, line := range rules {
		matches := lineParser.FindAllStringSubmatch(line, -1)

		value := matches[0][1]
		left := matches[0][2]
		right := matches[0][3]

		node := &Node{value, nil, left, nil, right, value[2:] == "Z"}
		nodes[value] = node

		if value[2:] == "A" {
			startNodes = append(startNodes, node)
		}
	}

	for _, node := range nodes {
		node.left = nodes[node.leftData]
		node.right = nodes[node.rightData]
	}

	return getMinimumStepsParallel(strings.Split(lines[0], ""), nodes, startNodes)
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	return part01(lines), part02(lines)
}
