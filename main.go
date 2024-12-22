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
	return HighCard
}
