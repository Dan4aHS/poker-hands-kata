package main

import (
	"sort"
)

// 2-10, J, Q, K, A
// H-hearts, D-diamonds, S-spares, C-clubs

// Combos:
const (
	HighCard      = "high card"
	Pair          = "pair"
	TwoPairs      = "two pairs"
	ThreeOfAKind  = "three of a kind"
	Straight      = "straight"
	Flush         = "flush"
	FullHouse     = "full house"
	FourOfAKind   = "four of a kind"
	StraightFlush = "straight flush"
	RoyalFlush    = "royal flush"
)

type PokerHand struct {
	cards []Card
}

func NewPokerHand(cards []string) *PokerHand {
	parsedCards := make([]Card, 0, len(cards))
	for _, c := range cards {
		parsedCards = append(parsedCards, NewCard(c))
	}

	return &PokerHand{
		cards: parsedCards,
	}
}

func (p *PokerHand) GetCombo() string {
	if v, ok := p.pairCombo(); ok {
		return v
	}

	if v, ok := p.noPairCombo(); ok {
		return v
	}

	return HighCard
}

func (p *PokerHand) noPairCombo() (string, bool) {
	switch {
	case p.hasRoyalFlush():
		return RoyalFlush, true
	case p.hasStraightFlush():
		return StraightFlush, true
	case p.hasFlush():
		return Flush, true
	case p.hasStraight():
		return Straight, true
	}

	return "", false
}

func (p *PokerHand) hasStraightFlush() bool {
	return p.hasFlush() && p.hasStraight()
}

func (p *PokerHand) pairCombo() (string, bool) {
	switch {
	case p.countPairs() == 1:
		return Pair, true
	case p.countPairs() == 2:
		return TwoPairs, true
	case p.countPairs() == 3:
		return ThreeOfAKind, true
	case p.countPairs() == 4:
		return FullHouse, true
	case p.countPairs() == 6:
		return FourOfAKind, true
	}

	return "", false
}

func (p *PokerHand) countPairs() int {
	counter := 0
	for i := 0; i < len(p.cards); i++ {
		for j := i + 1; j < len(p.cards); j++ {
			if p.cards[i].value == p.cards[j].value {
				counter++
			}
		}
	}

	return counter
}

func (p *PokerHand) hasFlush() bool {
	sample := p.cards[0].suit
	for i := 1; i < len(p.cards); i++ {
		if p.cards[i].suit != sample {
			return false
		}
	}

	return true
}

func (p *PokerHand) hasStraight() bool {
	values, ok := p.uniqueValues()
	if !ok {
		return false
	}

	for i := 0; i < len(values)-1; i++ {
		if values[i+1]-values[i] != 1 {
			if i == 3 && values[i+1] == 14 && values[0] == 2 {
				return true
			}

			return false
		}
	}

	return true
}

func (p *PokerHand) uniqueValues() ([]int, bool) {
	uniqueMap := make(map[int]struct{})
	for _, card := range p.cards {
		uniqueMap[card.value] = struct{}{}
	}

	if len(uniqueMap) != 5 {
		return nil, false
	}

	values := make([]int, 0, len(uniqueMap))
	for k := range uniqueMap {
		values = append(values, k)
	}

	sort.Ints(values)
	return values, true
}

func (p *PokerHand) hasRoyalFlush() bool {
	return p.hasStraightFlush() && p.hasAceAndTen()
}

func (p *PokerHand) hasAceAndTen() bool {
	var hasAce, hasTen bool
	for _, card := range p.cards {
		if card.value == valueMap['A'] {
			hasAce = true
		}
		if card.value == valueMap['T'] {
			hasTen = true
		}
	}

	return hasAce && hasTen
}
