package main

// 2-10, J, Q, K, A
// H-hearts, D-diamonds, S-spares, C-clubs

type PokerHand struct {
	cards []string
}

func NewPokerHand(cards []string) *PokerHand {
	return &PokerHand{
		cards: cards,
	}
}
