package day10

import (
	"math/bits"
	"strconv"
	"strings"

	"github.com/kingkero/adventofcode/2025/go/util"
)

func Part01(input []string) string {
	totalPresses := 0

	for _, line := range input {
		parts := strings.Split(line, " ")

		// Parse goal state from [.##.] format
		goal := 0
		for pos, light := range parts[0][1 : len(parts[0])-1] {
			if light == '#' {
				goal |= 1 << pos
			}
		}

		// Parse buttons - collect all parts in parentheses (stop at curly braces)
		var buttonParts []string
		for _, part := range parts[1:] {
			if strings.HasPrefix(part, "{") {
				break
			}
			buttonParts = append(buttonParts, part)
		}

		// Convert each button definition to a bitmap
		buttons := make([]int, len(buttonParts))
		for buttonIndex, definition := range buttonParts {
			for _, pos := range util.Map(strings.Split(strings.Trim(definition, "()"), ","), util.ParseInt) {
				buttons[buttonIndex] |= 1 << pos
			}
		}

		// Find minimum button presses using brute force
		// Since pressing a button twice cancels out, each button is pressed 0 or 1 times
		// Try all 2^n combinations and find minimum that reaches goal
		minPresses := len(buttons) + 1
		numCombinations := 1 << len(buttons)

		for combo := 0; combo < numCombinations; combo++ {
			// Calculate resulting state by XORing selected buttons
			state := 0
			for i := 0; i < len(buttons); i++ {
				if combo&(1<<i) != 0 {
					state ^= buttons[i]
				}
			}

			// Check if this combination reaches the goal
			if state == goal {
				presses := bits.OnesCount(uint(combo))
				if presses < minPresses {
					minPresses = presses
				}
			}
		}

		totalPresses += minPresses
	}

	return strconv.Itoa(totalPresses)
}

func Part02(input []string) string {
	totalPresses := 0

	for _, line := range input {
		parts := strings.Split(line, " ")

		// Parse buttons - collect all parts in parentheses (stop at curly braces)
		var buttonParts []string
		var targetPart string
		for _, part := range parts[1:] {
			if strings.HasPrefix(part, "{") {
				targetPart = part
				break
			}
			buttonParts = append(buttonParts, part)
		}

		// Parse button definitions as lists of counter indices they affect
		buttons := make([][]int, len(buttonParts))
		for i, definition := range buttonParts {
			buttons[i] = util.Map(strings.Split(strings.Trim(definition, "()"), ","), util.ParseInt)
		}

		// Parse target joltage values
		targets := util.Map(strings.Split(strings.Trim(targetPart, "{}"), ","), util.ParseInt)

		// Solve using Gaussian elimination
		totalPresses += solveWithGaussian(buttons, targets)
	}

	return strconv.Itoa(totalPresses)
}

// Rational number for exact arithmetic
type rational struct {
	num, den int64
}

func gcd(a, b int64) int64 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func newRational(num, den int64) rational {
	if den == 0 {
		return rational{0, 0} // Invalid
	}
	if den < 0 {
		num, den = -num, -den
	}
	g := gcd(num, den)
	if g != 0 {
		num, den = num/g, den/g
	}
	return rational{num, den}
}

func (r rational) add(other rational) rational {
	return newRational(r.num*other.den+other.num*r.den, r.den*other.den)
}

func (r rational) sub(other rational) rational {
	return newRational(r.num*other.den-other.num*r.den, r.den*other.den)
}

func (r rational) mul(other rational) rational {
	return newRational(r.num*other.num, r.den*other.den)
}

func (r rational) div(other rational) rational {
	return newRational(r.num*other.den, r.den*other.num)
}

func (r rational) isZero() bool {
	return r.num == 0
}

func (r rational) toInt() (int, bool) {
	if r.den == 0 || r.num%r.den != 0 {
		return 0, false
	}
	return int(r.num / r.den), true
}

func solveWithGaussian(buttons [][]int, targets []int) int {
	m := len(targets) // Number of counters (equations)
	n := len(buttons) // Number of buttons (variables)

	// Build augmented matrix [A | b] using rationals
	// A[i][j] = 1 if button j affects counter i
	matrix := make([][]rational, m)
	for i := range matrix {
		matrix[i] = make([]rational, n+1)
		for j := range matrix[i] {
			matrix[i][j] = newRational(0, 1)
		}
		matrix[i][n] = newRational(int64(targets[i]), 1) // RHS
	}
	for j, button := range buttons {
		for _, counter := range button {
			matrix[counter][j] = newRational(1, 1)
		}
	}

	// Gaussian elimination with partial pivoting to RREF
	pivotRows := make([]int, n) // Which row has the pivot for each column
	for i := range pivotRows {
		pivotRows[i] = -1
	}

	row := 0
	for col := 0; col < n && row < m; col++ {
		// Find pivot row (non-zero in this column)
		pivotRow := -1
		for i := row; i < m; i++ {
			if !matrix[i][col].isZero() {
				pivotRow = i
				break
			}
		}
		if pivotRow == -1 {
			continue // No pivot in this column, free variable
		}

		// Swap rows
		matrix[row], matrix[pivotRow] = matrix[pivotRow], matrix[row]

		// Scale pivot row
		scale := matrix[row][col]
		for j := col; j <= n; j++ {
			matrix[row][j] = matrix[row][j].div(scale)
		}

		// Eliminate column in other rows
		for i := 0; i < m; i++ {
			if i != row && !matrix[i][col].isZero() {
				factor := matrix[i][col]
				for j := col; j <= n; j++ {
					matrix[i][j] = matrix[i][j].sub(factor.mul(matrix[row][j]))
				}
			}
		}

		pivotRows[col] = row
		row++
	}

	// Identify pivot and free variables
	var pivotVars, freeVars []int
	for col := 0; col < n; col++ {
		if pivotRows[col] >= 0 {
			pivotVars = append(pivotVars, col)
		} else {
			freeVars = append(freeVars, col)
		}
	}

	// Compute max values for free variables (bounded by min target of affected counters)
	maxFree := make([]int, len(freeVars))
	for i, fv := range freeVars {
		maxFree[i] = 1 << 30
		for _, counter := range buttons[fv] {
			if targets[counter] < maxFree[i] {
				maxFree[i] = targets[counter]
			}
		}
	}

	// Search over free variables to find minimum sum solution
	best := 1 << 30
	freeValues := make([]int, len(freeVars))

	var searchFree func(idx int)
	searchFree = func(idx int) {
		if idx == len(freeVars) {
			// Compute pivot variable values based on free variable values
			solution := make([]int, n)
			for i, fv := range freeVars {
				solution[fv] = freeValues[i]
			}

			valid := true
			total := 0
			for i, fv := range freeVars {
				total += freeValues[i]
				if total >= best {
					return
				}
				_ = fv
			}

			for _, pv := range pivotVars {
				r := pivotRows[pv]
				// value = RHS - sum of (coef * freeVar) for all free vars
				val := matrix[r][n]
				for i, fv := range freeVars {
					val = val.sub(matrix[r][fv].mul(newRational(int64(freeValues[i]), 1)))
				}
				intVal, ok := val.toInt()
				if !ok || intVal < 0 {
					valid = false
					break
				}
				solution[pv] = intVal
				total += intVal
				if total >= best {
					return
				}
			}

			if valid && total < best {
				best = total
			}
			return
		}

		for v := 0; v <= maxFree[idx]; v++ {
			freeValues[idx] = v
			searchFree(idx + 1)
		}
	}
	searchFree(0)

	return best
}
