package pokerlarvis

import (
	"fmt"
	"iceye/internal/poker"
)

type (
	PokerLarvis struct {
		compareRules []poker.CompareRuleFunc
	}
	Winner string
)

const (
	NoWinner Winner = ""
	Tie      Winner = "Tie"
	Hand1    Winner = "Hand 1"
	Hand2    Winner = "Hand 2"
)

// New creates and configures a Poker Larvis game
func New() PokerLarvis {
	orderedRules := []poker.CompareRuleFunc{
		poker.ComparePokers,
		poker.CompareFullHouses,
		poker.CompareTriples,
		poker.CompareDoublePairs,
		poker.CompareSinglePairs,
		poker.CompareSingles,
	}

	return PokerLarvis{
		compareRules: orderedRules,
	}
}

// Game starts a poker game and accepts two arguments for each players cards
func (p PokerLarvis) Game(hand1, hand2 string) (Winner, error) {
	c1, err := poker.Parse(hand1)
	if err != nil {
		return NoWinner, fmt.Errorf("hand 1 error: %w", err)
	}

	c2, err := poker.Parse(hand2)
	if err != nil {
		return NoWinner, fmt.Errorf("hand 2 error: %w", err)
	}

	winner, err := c1.Compare(*c2, p.compareRules)
	if err != nil {
		return NoWinner, fmt.Errorf("game error: %w", err)
	}

	switch winner {
	case 1:
		return Hand1, nil
	case -1:
		return Hand2, nil
	case 0:
		return Tie, nil
	default:
		return NoWinner, fmt.Errorf("unknown game result: %d", winner)
	}
}
