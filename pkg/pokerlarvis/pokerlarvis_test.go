package pokerlarvis

import (
	"fmt"
	"iceye/internal/poker"
	"testing"

	"github.com/stretchr/testify/suite"
)

type (
	pokerLarvisTestSuite struct {
		suite.Suite
	}
)

func TestPokerLarvisSuite(t *testing.T) {
	suite.Run(t, new(pokerLarvisTestSuite))
}

func (s pokerLarvisTestSuite) TestPokerLarvis() {
	poker := New()

	for idx, test := range []struct {
		hand1, hand2, expectedWinner string
	}{
		{"AAAQQ", "QQAAA", "Tie"}, //1
		{"53QQ2", "Q53Q2", "Tie"},
		{"53888", "88385", "Tie"},
		{"QQAAA", "AAAQQ", "Tie"},
		{"Q53Q2", "53QQ2", "Tie"},
		{"88385", "53888", "Tie"},
		{"AAAQQ", "QQQAA", "Hand 1"},
		{"Q53Q4", "53QQ2", "Hand 1"},
		{"53888", "88375", "Hand 1"},
		{"33337", "QQAAA", "Hand 1"}, //10
		{"22333", "AAA58", "Hand 1"},
		{"33389", "AAKK4", "Hand 1"},
		{"44223", "AA892", "Hand 1"},
		{"22456", "AKQJT", "Hand 1"},
		{"99977", "77799", "Hand 1"},
		{"99922", "88866", "Hand 1"},
		{"9922A", "9922K", "Hand 1"},
		{"99975", "99965", "Hand 1"},
		{"99975", "99974", "Hand 1"},
		{"99752", "99652", "Hand 1"}, //20
		{"99752", "99742", "Hand 1"},
		{"99753", "99752", "Hand 1"},
		{"QQQAA", "AAAQQ", "Hand 2"},
		{"53QQ2", "Q53Q4", "Hand 2"},
		{"88375", "53888", "Hand 2"},
		{"QQAAA", "33337", "Hand 2"},
		{"AAA58", "22333", "Hand 2"},
		{"AAKK4", "33389", "Hand 2"},
		{"AA892", "44223", "Hand 2"},
		{"AKQJT", "22456", "Hand 2"}, //30
		{"77799", "99977", "Hand 2"},
		{"88866", "99922", "Hand 2"},
		{"9922K", "9922A", "Hand 2"},
		{"99965", "99975", "Hand 2"},
		{"99974", "99975", "Hand 2"},
		{"99652", "99752", "Hand 2"},
		{"99742", "99752", "Hand 2"},
		{"99752", "99753", "Hand 2"},
	} {
		s.Run(fmt.Sprintf("scenario %d: %s %s", idx+1, test.hand1, test.hand2), func() {
			actualWinner, err := poker.Game(test.hand1, test.hand2)
			s.NoError(err)
			s.Equal(Winner(test.expectedWinner), actualWinner)
		})
	}
}

func (s pokerLarvisTestSuite) TestGameReturnsError() {
	invalidRule := func(c1 poker.Combinations, c2 poker.Combinations) int {
		return -2
	}

	poker := PokerLarvis{
		compareRules: []poker.CompareRuleFunc{invalidRule},
	}

	for _, test := range []struct {
		name, hand1, hand2, expectedError string
	}{
		{"hand 1 parse error", "12345", "23456", "hand 1 error: invalid symbol 1"},
		{"hand 2 parse error", "23456", "12345", "hand 2 error: invalid symbol 1"},
		{"comparator error", "23457", "23456", "unknown game result: -2"},
	} {
		s.Run(test.name, func() {
			winner, err := poker.Game(test.hand1, test.hand2)
			s.ErrorContains(err, test.expectedError)
			s.Equal(NoWinner, winner)
		})
	}
}
