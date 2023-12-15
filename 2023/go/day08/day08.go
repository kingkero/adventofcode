package day08

import (
	"log"
	"math"
	"regexp"
	"slices"
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

func getMinimumSteps(instructions []string, nodes map[string]*Node, checkNode *Node) int {
	steps := 0
	for {
		checkNode = checkNode.step(instructions[steps%len(instructions)])

		if checkNode.isEndNote {
			return steps + 1
		}

		steps++
	}
}

func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func getPrimeFactors(value int, primes []int) map[int]int {
	factors := make(map[int]int)
	for _, prime := range primes {
		if prime > value {
			break
		}

		for value%prime == 0 {
			value /= prime
			factors[prime]++
		}
	}
	return factors
}

func getMinimumStepsParallel(instructions []string, nodes map[string]*Node, checkNodes []*Node) int {
	steps := make([]int, len(checkNodes))
	// once we know how many steps for each node
	for i, node := range checkNodes {
		steps[i] = getMinimumSteps(instructions, nodes, node)
	}
	slices.Sort(steps)

	uniqueSteps := slices.Compact(steps)
	primes := sieveOfEratosthenes(slices.Max(uniqueSteps))

	// collect their prime factors to find least common multiple of all exits
	primeFactors := make(map[int]int)
	for _, step := range uniqueSteps {
		for prime, occurences := range getPrimeFactors(step, primes) {
			if existingFactor, ok := primeFactors[prime]; !ok || existingFactor < occurences {
				primeFactors[prime] = occurences
			}
		}
	}

	result := 1.0
	for prime, occurences := range primeFactors {
		result *= math.Pow(float64(prime), float64(occurences))
	}

	return int(result)
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

	return getMinimumSteps(strings.Split(lines[0], ""), nodes, nodes["AAA"])
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
