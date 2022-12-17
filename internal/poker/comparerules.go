package poker

import (
	"strings"
)

// CompareRuleFunc acts as an interface to the comparator rules which are used by the game
type CompareRuleFunc = func(c1 Combinations, c2 Combinations) int

// ComparePokers checks four of a kind
//   example: 7777K
func ComparePokers(c1 Combinations, c2 Combinations) int {
	hasPoker := func(c Combinations) bool {
		return c.quadruple != nil
	}

	if isMissingCombination(c1, c2, hasPoker) {
		return 0
	}

	if res := compareCombinations(c1, c2, hasPoker); res != 0 {
		return res
	}

	return compareRune(c1.quadruple, c2.quadruple)
}

// CompareFullHouses checks full house
//   example: 777KK
func CompareFullHouses(c1 Combinations, c2 Combinations) int {
	hasFullHouse := func(c Combinations) bool {
		return c.triple != nil && len(c.pairs) == 1
	}

	if isMissingCombination(c1, c2, hasFullHouse) {
		return 0
	}

	if res := compareCombinations(c1, c2, hasFullHouse); res != 0 {
		return res
	}

	if res := compareRune(c1.triple, c2.triple); res != 0 {
		return res
	}

	pair1 := firsPairOrNil(c1)
	pair2 := firsPairOrNil(c2)
	return compareRune(pair1, pair2)
}

// CompareTriples checks triples
//   example: 77723
func CompareTriples(c1 Combinations, c2 Combinations) int {
	hasTriple := func(c Combinations) bool {
		return c.triple != nil
	}

	if isMissingCombination(c1, c2, hasTriple) {
		return 0
	}

	if res := compareCombinations(c1, c2, hasTriple); res != 0 {
		return res
	}

	return compareRune(c1.triple, c2.triple)
}

// CompareDoublePairs checks double pairs
//   example: 77223
func CompareDoublePairs(c1 Combinations, c2 Combinations) int {
	hasDoublePair := func(c Combinations) bool {
		return len(c.pairs) == 2
	}

	if isMissingCombination(c1, c2, hasDoublePair) {
		return 0
	}

	if res := compareCombinations(c1, c2, hasDoublePair); res != 0 {
		return res
	}

	if res := compareRune(&c1.pairs[0], &c2.pairs[0]); res != 0 {
		return res
	}

	return compareRune(&c1.pairs[1], &c2.pairs[1])
}

// CompareSinglePairs checks single pair
//   example: 77234
func CompareSinglePairs(c1 Combinations, c2 Combinations) int {
	hasSinglePair := func(c Combinations) bool {
		return len(c.pairs) == 1
	}

	if isMissingCombination(c1, c2, hasSinglePair) {
		return 0
	}

	if res := compareCombinations(c1, c2, hasSinglePair); res != 0 {
		return res
	}

	pair1 := firsPairOrNil(c1)
	pair2 := firsPairOrNil(c2)
	return compareRune(pair1, pair2)
}

// CompareSingles checks single cards
// example: 23456
func CompareSingles(c1 Combinations, c2 Combinations) int {
	hasSingles := func(c Combinations) bool {
		return len(c.singles) != 0
	}

	if isMissingCombination(c1, c2, hasSingles) {
		return 0
	}

	len1 := len(c1.singles)
	len2 := len(c2.singles)

	l := len1
	if len2 > len1 {
		l = len2
	}

	for i := 0; i < l; i++ {
		if res := compareRune(&c1.singles[i], &c2.singles[i]); res != 0 {
			return res
		}
	}
	return 0
}

func isMissingCombination(c1 Combinations, c2 Combinations, checkFunc func(c Combinations) bool) bool {
	return !checkFunc(c1) && !checkFunc(c2)
}

func compareCombinations(c1 Combinations, c2 Combinations, checkFunc func(c Combinations) bool) int {
	if checkFunc(c1) && !checkFunc(c2) {
		return 1
	}

	if !checkFunc(c1) && checkFunc(c2) {
		return -1
	}

	return 0
}

func compareRune(r1, r2 *rune) int {
	zero := rune(0)
	if r1 == nil {
		r1 = &zero
	}
	if r2 == nil {
		r2 = &zero
	}

	idx1 := strings.Index(symbols, string(*r1))
	idx2 := strings.Index(symbols, string(*r2))

	switch {
	case idx1 > idx2:
		return 1
	case idx1 < idx2:
		return -1
	default:
		return 0
	}
}

func firsPairOrNil(c Combinations) *rune {
	var pair *rune
	if len(c.pairs) > 0 {
		pair = &c.pairs[0]
	}
	return pair
}
