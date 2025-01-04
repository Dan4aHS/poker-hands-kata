package deck

import (
	"errors"
	"math/rand/v2"
	"poker-hands-kata/card"
)

type Deck struct {
	current []card.Card
}

func New() *Deck {
	return &Deck{
		current: newDeckOfCards(),
	}
}

func (d *Deck) PickCard() (card.Card, error) {
	if len(d.current) == 0 {
		return card.Card{}, errors.New("deck is empty")
	}

	picked := d.current[0]
	d.current = d.current[1:]

	return picked, nil
}

func newDeckOfCards() []card.Card {
	deck := make([]card.Card, 0)

	for c := range card.CardIterator() {
		deck = append(deck, c)
	}

	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}
