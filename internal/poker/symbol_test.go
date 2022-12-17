package poker

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const totalUnicodeCharacters = 149186

type symbolTestSuite struct {
	suite.Suite
}

func TestSymbolTestSuite(t *testing.T) {
	suite.Run(t, new(symbolTestSuite))
}

func (s symbolTestSuite) TestHasValidSymbols() {
	for _, cards := range []string{
		"23456789TJQKA",
		"",
		"2223334445556667778889999TTTTJJJQQQQKKKKAAAA",
		"A",
	} {
		s.NoError(hasValidSymbols(cards))
	}
}

func (s symbolTestSuite) TestHasValidSymbols_ReturnsErrorOnInvalidSymbol() {
	for _, test := range []struct {
		name, input, expectedErrorChar string
	}{
		{"space", " ", " "},
		{"zero", "0", "0"},
		{"one", "1", "1"},
		{"first invalid letter", "234BC67TJ", "B"},
		{"non alphanumeric", "666.JJJ", "."},
		{"lowercase", "TJjQ", "j"},
		{"unicode character", "23📡TT", "📡"},
	} {
		s.Run(test.name, func() {
			s.EqualError(hasValidSymbols(test.input), "invalid symbol "+test.expectedErrorChar)
		})
	}
}

func (s symbolTestSuite) TestHasValidSymbols_ReturnsErrorOnGeneratedInvalidSymbol() {
	for i := 0; i < totalUnicodeCharacters; i++ {
		symbol := string(rune(i))
		if strings.Index("23456789TJQKA", symbol) > -1 {
			continue
		}
		s.EqualError(hasValidSymbols(symbol), "invalid symbol "+symbol)
	}
}

func (s symbolTestSuite) TestSymbolSort() {
	for _, test := range []struct {
		name, input, expected string
	}{
		{"all reversed", "AKQJT98765432", "23456789TJQKA"},
		{"all same", "AAAA", "AAAA"},
		{"unordered", "2A5J8", "258JA"},
		{"fullhouse", "Q2Q2Q", "22QQQ"},
		{"invalid characters to the end", "213", "231"},
	} {
		s.Run(test.name, func() {
			input := []rune(test.input)
			sort.Sort(symbolSort(input))
			s.Equal(test.expected, string(input))
		})

	}
}
