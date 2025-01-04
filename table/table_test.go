package table

import "testing"

func TestNewTable(t *testing.T) {
	pt := NewPokerTable()

	if pt.state != StateEmpty {
		t.Errorf("NewPokerTable() should return empty state")
	}
}

func TestGetCardsOnEmptyTable(t *testing.T) {
	pt := NewPokerTable()
	_ = pt.GetCurrentCards()
}
