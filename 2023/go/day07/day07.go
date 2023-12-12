package day07

import (
	"cmp"
	"log"
	"slices"
	"strings"

	"github.com/kingkero/adventofcode/2023/go/util"
)

// Get strength for a single hand and the value map in original order.
// - high card: 1
// - one pair: 2
// - two pair: 3
// - three of a kind: 4
// - full house: 5
// - four of a kind: 6
// - five of a kind: 7
func getStrength(hand string) (int, []int) {
	cardToValue := map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	cards := util.Map(strings.Split(hand, ""), func(card string) int {
		return cardToValue[card]
	})
	cardsSorted := slices.Clone(cards)
	slices.Sort(cardsSorted)

	length, pairs, threeOfAKinds, fourOfAKind, fiveOfAKind := 1, 0, 0, 0, 0

	prev := cardsSorted[0]
	for i := 1; i < len(cardsSorted); i++ {
		if cardsSorted[i] == prev {
			length++
			prev = cardsSorted[i]
			continue
		}

		if length == 2 {
			pairs++
		} else if length == 3 {
			threeOfAKinds++
		} else if length == 4 {
			fourOfAKind++
		} else if length == 5 {
			fiveOfAKind++
		}

		prev = cardsSorted[i]
		length = 1
	}

	if length > 1 {
		if length == 2 {
			pairs++
		} else if length == 3 {
			threeOfAKinds++
		} else if length == 4 {
			fourOfAKind++
		} else if length == 5 {
			fiveOfAKind++
		}
	}

	if fiveOfAKind > 0 {
		return 7, cards
	}
	if fourOfAKind > 0 {
		return 6, cards
	}
	if pairs > 0 && threeOfAKinds > 0 {
		return 5, cards
	}
	if threeOfAKinds > 0 {
		return 4, cards
	}
	if pairs > 1 {
		return 3, cards
	}
	if pairs > 0 {
		return 2, cards
	}

	return 1, cards
}

func getStrengthPart02(hand string) (int, []int) {
	cardToValue := map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 1,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	cards := util.Map(strings.Split(hand, ""), func(card string) int {
		return cardToValue[card]
	})
	cardsSorted := slices.Clone(cards)
	slices.Sort(cardsSorted)

	jokers, length, pairs, threeOfAKinds, fourOfAKind, fiveOfAKind := 0, 1, 0, 0, 0, 0
	for i := 0; i < len(cardsSorted) && cardsSorted[i] == cardToValue["J"]; i++ {
		jokers++
	}
	if jokers == 5 {
		return 7, cards
	}
	cardsSorted = cardsSorted[jokers:]
	prev := cardsSorted[0]
	for i := 1; i < len(cardsSorted); i++ {
		if cardsSorted[i] == prev {
			length++
			prev = cardsSorted[i]
			continue
		}

		if length == 2 {
			pairs++
		} else if length == 3 {
			threeOfAKinds++
		} else if length == 4 {
			fourOfAKind++
		} else if length == 5 {
			fiveOfAKind++
		}

		prev = cardsSorted[i]
		length = 1
	}

	if length > 1 {
		if length == 2 {
			pairs++
		} else if length == 3 {
			threeOfAKinds++
		} else if length == 4 {
			fourOfAKind++
		} else if length == 5 {
			fiveOfAKind++
		}
	}

	for ; jokers > 0; jokers-- {
		if fourOfAKind > 0 {
			fiveOfAKind++
			fourOfAKind--
		} else if threeOfAKinds > 0 {
			fourOfAKind++
			threeOfAKinds--
		} else if pairs > 0 {
			threeOfAKinds++
			pairs--
		} else {
			pairs++
		}
	}

	if fiveOfAKind > 0 {
		return 7, cards
	}
	if fourOfAKind > 0 {
		return 6, cards
	}
	if pairs > 0 && threeOfAKinds > 0 {
		return 5, cards
	}
	if threeOfAKinds > 0 {
		return 4, cards
	}
	if pairs > 1 {
		return 3, cards
	}
	if pairs > 0 {
		return 2, cards
	}

	return 1, cards
}

type WeightedHand struct {
	sorted []int
	weight int
	bid    int
}

func sortHands(a WeightedHand, b WeightedHand) int {
	if a.weight != b.weight {
		return cmp.Compare(a.weight, b.weight)
	}

	for i := range a.sorted {
		if a.sorted[i] != b.sorted[i] {
			return cmp.Compare(a.sorted[i], b.sorted[i])
		}
	}

	return 0
}

func part01(lines []string) int {
	var hands []WeightedHand
	for _, line := range lines {
		parts := strings.Split(line, " ")

		hand := parts[0]
		bid := util.ParseInt(parts[1])

		strength, handValues := getStrength(hand)
		var tmp = WeightedHand{handValues, strength, bid}
		hands = append(hands, tmp)
	}

	slices.SortFunc(hands, sortHands)

	result := 0
	for i, hand := range hands {
		result += hand.bid * (i + 1)
	}

	return result
}

func part02(lines []string) int {
	var hands []WeightedHand
	for _, line := range lines {
		parts := strings.Split(line, " ")

		hand := parts[0]
		bid := util.ParseInt(parts[1])

		strength, handValues := getStrengthPart02(hand)
		var tmp = WeightedHand{handValues, strength, bid}
		hands = append(hands, tmp)
	}

	slices.SortFunc(hands, sortHands)

	result := 0
	for i, hand := range hands {
		result += hand.bid * (i + 1)
	}

	return result
}

func Solve(input string) (int, int) {
	lines, err := util.ReadFile(input)
	if err != nil {
		log.Fatal("Could not open file "+input, err)
	}

	return part01(lines), part02(lines)
}
