package main

import "testing"

func TestHighCard(t *testing.T) {
	ph := NewPokerHand([]string{"2C", "9S", "8H", "5H", "6H"})
	assertString(t, HighCard, ph.GetCombo())
}

func assertString(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
