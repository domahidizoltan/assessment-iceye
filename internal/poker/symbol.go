package poker

import (
	"errors"
	"strings"
)

const (
	symbols = "23456789TJQKA"
)

func hasValidSymbols(cards string) error {
	for _, c := range cards {
		if symbolIndex(c) == -1 {
			return errors.New("invalid symbol " + string(c))
		}
	}
	return nil
}

func symbolIndex(symbol rune) int {
	return strings.Index(symbols, string(symbol))
}

type symbolSort []rune

func (s symbolSort) Len() int {
	return len(s)
}

func (s symbolSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s symbolSort) Less(i, j int) bool {
	last := len(symbols)

	si := symbolIndex(s[i])
	if si < 0 {
		si = last
	}

	sj := symbolIndex(s[j])
	if sj < 0 {
		sj = last
	}

	return si < sj
}
