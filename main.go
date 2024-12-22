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
	if p.hasDoubleValue() {
		return Pair
	}
	return HighCard
}

func (p *PokerHand) hasDoubleValue() bool {
	for i := 0; i < len(p.cards); i++ {
		for j := i + 1; j < len(p.cards); j++ {
			if p.cards[i][0] == p.cards[j][0] {
				return true
			}
		}
	}

	return false
}
