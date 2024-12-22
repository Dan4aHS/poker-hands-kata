package main

import "testing"

func TestInitialPokerHand(t *testing.T) {
	_ = NewPokerHand([]string{"2C", "9S", "8H", "4D", "5H"})
}

func TestGetCombo(t *testing.T) {
	ph := NewPokerHand([]string{"2C", "9S", "8H", "4D", "5H"})
	ph.GetCombo()
}
