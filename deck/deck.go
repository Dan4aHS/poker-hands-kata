package deck

import "poker-hands-kata/card"

type Deck struct {
	current []card.Card
}

func New() *Deck {
	return &Deck{
		current: newDeckOfCards(),
	}
}

func (d *Deck) PickCard() (card.Card, error) {
	return d.current[0], nil
}

func newDeckOfCards() []card.Card {
	deck := make([]card.Card, 0)

	for c := range card.CardIterator() {
		deck = append(deck, c)
	}

	return deck
}
