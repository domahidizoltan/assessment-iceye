package poker

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type combinationsTestSuite struct {
	suite.Suite
}

func TestCombinationsSuite(t *testing.T) {
	suite.Run(t, new(combinationsTestSuite))
}

func (s combinationsTestSuite) TestParse_ReturnError() {
	for _, test := range []struct {
		name, cards, expectedErrorMsg string
	}{
		{"less cards", "2345", "must have 5 cards"},
		{"more cards", "234567", "must have 5 cards"},
		{"invalid card", "23Y56", "invalid symbol: Y"},
		{"five of a kind", "22222", "can have at most 4 of a kind"},
	} {
		s.Run(test.name, func() {
			_, err := Parse(test.cards)
			s.Error(err, test.expectedErrorMsg)
		})
	}
}

func (s combinationsTestSuite) TestParse() {
	for _, test := range []struct {
		name, cards string
		q, t        *rune
		p, s        []rune
	}{
		{"poker", "22228", ptr('2'), nil, nil, []rune("8")},
		{"full house", "22288", nil, ptr('2'), []rune("8"), nil},
		{"triple", "22289", nil, ptr('2'), nil, []rune("98")},
		{"double pair", "22889", nil, nil, []rune("82"), []rune("9")},
		{"single pair", "2289A", nil, nil, []rune("2"), []rune("A98")},
		{"all different", "29A5J", nil, nil, nil, []rune("AJ952")},
	} {
		s.Run(test.name, func() {
			comb, err := Parse(test.cards)
			s.NoError(err)

			if test.q == nil {
				s.Nil(comb.quadruple)
			} else {
				s.Equal(string(*test.q), string(*comb.quadruple))
			}

			if test.t == nil {
				s.Nil(comb.triple)
			} else {
				s.Equal(string(*test.t), string(*comb.triple))
			}

			s.Equal(string(test.p), string(comb.pairs))
			s.Equal(string(test.s), string(comb.singles))
		})
	}
}

func (s combinationsTestSuite) TestCompareSame() {
	c1, _ := Parse("2K4A6")
	c2, _ := Parse("A2K64")
	res, err := c1.Compare(*c2, nil)
	s.NoError(err)
	s.Equal(0, res)
}

func (s combinationsTestSuite) TestCompareReturnsError_WhenNoComparatorsGiven() {
	callOrder := []string{}
	firstRule := func(c1 Combinations, c2 Combinations) int {
		callOrder = append(callOrder, "first")
		return 0
	}
	secondRule := func(c1 Combinations, c2 Combinations) int {
		callOrder = append(callOrder, "second")
		return -1
	}

	c1, _ := Parse("23456")
	c2, _ := Parse("45678")
	res, err := c1.Compare(*c2, []CompareRuleFunc{firstRule, secondRule})

	s.NoError(err)
	s.Equal(-1, res)
	s.Equal(callOrder, []string{"first", "second"})
}

func (s combinationsTestSuite) TestCompareCallRulesInOrder() {
	c1, _ := Parse("23456")
	c2, _ := Parse("45678")
	_, err := c1.Compare(*c2, nil)
	s.Error(err, "could not compare combinations")
}

func ptr(r rune) *rune {
	return &r
}
