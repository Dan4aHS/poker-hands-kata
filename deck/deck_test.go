package deck

import "testing"

func TestNewDeckOfCards(t *testing.T) {
	cards := newDeckOfCards()

	if len(cards) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(cards))
	}
}

func TestDeckPickCard(t *testing.T) {
	deck := New()

	newCard, _ := deck.PickCard()

	if newCard.Suit() == "" {
		t.Error("Expected a non empty card")
	}
}

func TestDeckPickTwoCards(t *testing.T) {
	deck := New()

	first, _ := deck.PickCard()
	second, _ := deck.PickCard()

	if first == second {
		t.Error("Expected a different card")
	}
}

func TestDeckPickFromEmptyDeck(t *testing.T) {
	deck := New()

	for range 52 {
		_, _ = deck.PickCard()
	}

	if _, err := deck.PickCard(); err == nil {
		t.Error("Expected an error")
	}
}
