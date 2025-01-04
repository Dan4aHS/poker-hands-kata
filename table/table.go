package table

import "poker-hands-kata/card"

type PokerTable struct {
	cards []card.Card
	state State
}

func NewPokerTable() *PokerTable {
	return &PokerTable{
		cards: make([]card.Card, 0, 5),
		state: StateEmpty,
	}
}

func (t *PokerTable) GetCurrentCards() []card.Card {
	return t.cards
}
