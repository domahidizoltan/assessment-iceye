package poker

import (
	"errors"
	"sort"
	"strings"
)

const numCards = 5

// Combinations holds the tokenized representation of the actual cards. 
// Every token is turned into a single character and it is sorted descending by it's value
// examples:
//   4444K -> quadruple: 4, singles: K
//   444KK -> triple: 4, pairs: K
//   44KK6 -> pairs: K4, singles 6
//   44456 -> triple 4, singles: 65
type Combinations struct {
	symbols   string
	quadruple *rune
	triple    *rune
	pairs     []rune
	singles   []rune
}

// Parse validates and turns string to Combinations
func Parse(cards string) (*Combinations, error) {
	sanitized := strings.ToUpper(strings.TrimSpace(cards))
	if len(sanitized) != numCards {
		return nil, errors.New("must have 5 cards")
	}

	if err := hasValidSymbols(sanitized); err != nil {
		return nil, err
	}

	combinations, err := getCombinations(sanitized)
	if err != nil {
		return nil, err
	}

	return combinations, nil
}

func getCombinations(symbols string) (*Combinations, error) {
	reversedSymbols := []rune(symbols)
	sort.Sort(sort.Reverse(symbolSort(reversedSymbols)))

	combinations := Combinations{
		symbols: string(reversedSymbols),
		pairs:   make([]rune, 0, 2),
		singles: make([]rune, 0, numCards),
	}

	occurrence := map[rune]int{}
	for _, symbol := range reversedSymbols {
		if _, ok := occurrence[symbol]; !ok {
			occurrence[symbol] = 0
		}
		occurrence[symbol]++
	}

	for _, symbol := range reversedSymbols {
		count := occurrence[symbol]
		switch count {
		case 5:
			return nil, errors.New("can have at most 4 of a kind")
		case 4:
			s := symbol
			combinations.quadruple = &s
		case 3:
			s := symbol
			combinations.triple = &s
		case 2:
			combinations.pairs = append(combinations.pairs, symbol)
		case 1:
			combinations.singles = append(combinations.singles, symbol)
		}
		delete(occurrence, symbol)
	}

	return &combinations, nil
}

// Compare runs the combinations in order against the comparators and returns the comparison value of the Combinations
// As every other a comparator it will return:
//   1 when the current Combination is bigger
//   -1 when the other Combination is bigger
//   0 when the Combinations are equal
func (c Combinations) Compare(other Combinations, compareRules []CompareRuleFunc) (int, error) {
	if c.symbols == other.symbols {
		return 0, nil
	}

	for _, compare := range compareRules {
		if res := compare(c, other); res != 0 {
			return res, nil
		}
	}

	return 0, errors.New("could not compare combinations")
}
