package main

import "testing"

func TestInitialPokerHand(t *testing.T) {
	_ = NewPokerHand([]string{"2C", "9S", "8H", "4D", "5H"})
}
