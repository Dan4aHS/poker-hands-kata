package main

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
	cards []string
}

func NewPokerHand(cards []string) *PokerHand {
	return &PokerHand{
		cards: cards,
	}
}

func (p *PokerHand) GetCombo() string {
	switch {
	case p.countPairs() == 1:
		return Pair
	case p.countPairs() == 2:
		return TwoPairs
	case p.countPairs() == 3:
		return ThreeOfAKind
	case p.countPairs() == 6:
		return FourOfAKind
	case p.hasFlush():
		return Flush
	default:
		return HighCard
	}
}

func (p *PokerHand) countPairs() int {
	counter := 0
	for i := 0; i < len(p.cards); i++ {
		for j := i + 1; j < len(p.cards); j++ {
			if p.cards[i][0] == p.cards[j][0] {
				counter++
			}
		}
	}

	return counter
}

func (p *PokerHand) hasFlush() bool {
	sample := p.cards[0][1]
	for i := 1; i < len(p.cards); i++ {
		if p.cards[i][1] != sample {
			return false
		}
	}

	return true
}
