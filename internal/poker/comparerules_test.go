package poker

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

const anyCards = "23456"

type (
	compareRulesTestSuite struct {
		suite.Suite
	}
	scenario struct {
		name, cards1, cards2 string
		fn                   CompareRuleFunc
		expected             int
	}
)

func TestCompareRulesSuite(t *testing.T) {
	suite.Run(t, new(compareRulesTestSuite))
}

func (s compareRulesTestSuite) TestCompareRules() {
	for _, test := range []scenario{
		{"no pokers", anyCards, anyCards, ComparePokers, 0},
		{"same pokers", "44449", "44448", ComparePokers, 0},
		{"only hand1 poker", "44449", anyCards, ComparePokers, 1},
		{"only hand2 poker", anyCards, "44449", ComparePokers, -1},
		{"hand1 poker wins", "44449", "3333A", ComparePokers, 1},
		{"hand2 poker wins", "3333A", "44449", ComparePokers, -1},

		{"no full houses", anyCards, anyCards, CompareFullHouses, 0},
		{"same full houses", "44433", "44433", CompareFullHouses, 0},
		{"only hand1 full house", "44433", anyCards, CompareFullHouses, 1},
		{"only hand2 full house", anyCards, "44433", CompareFullHouses, -1},
		{"hand1 full house wins", "44433", "333AA", CompareFullHouses, 1},
		{"hand2 full house wins", "333AA", "44433", CompareFullHouses, -1},

		{"no triple", anyCards, anyCards, CompareTriples, 0},
		{"same triple", "44432", "44432", CompareTriples, 0},
		{"only hand1 triple", "44432", anyCards, CompareTriples, 1},
		{"only hand2 triple", anyCards, "44432", CompareTriples, -1},
		{"hand1 triple wins", "44432", "333AJ", CompareTriples, 1},
		{"hand2 triple wins", "333AJ", "44432", CompareTriples, -1},

		{"no double pair", "44567", "55678", CompareDoublePairs, 0},
		{"same double pair", "44332", "44332", CompareDoublePairs, 0},
		{"only hand1 double pair", "44332", anyCards, CompareDoublePairs, 1},
		{"only hand2 double pair", anyCards, "44332", CompareDoublePairs, -1},
		{"hand1 double pair wins", "44332", "4422J", CompareDoublePairs, 1},
		{"hand2 double pair wins", "4422J", "44332", CompareDoublePairs, -1},

		{"no single pair", anyCards, anyCards, CompareSinglePairs, 0},
		{"same single pair", "44532", "44532", CompareSinglePairs, 0},
		{"only hand1 single pair", "44532", anyCards, CompareSinglePairs, 1},
		{"only hand2 single pair", anyCards, "44532", CompareSinglePairs, -1},
		{"hand1 single pair wins", "44532", "3342J", CompareSinglePairs, 1},
		{"hand2 single pair wins", "3342J", "44532", CompareSinglePairs, -1},

		{"no singles", "33322", "44433", CompareSingles, 0},
		{"hand1 singles wins", "56789", "46789", CompareSingles, 1},
		{"hand2 singles wins", "46789", "56789", CompareSingles, -1},
	} {
		s.Run(test.name, func() {
			c1, err1 := Parse(test.cards1)
			c2, err2 := Parse(test.cards2)
			s.NoError(err1)
			s.NoError(err2)
			res := test.fn(*c1, *c2)
			s.Equal(test.expected, res)
		})
	}
}
