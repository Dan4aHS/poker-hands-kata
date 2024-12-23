package main

import "testing"

func TestHighCard(t *testing.T) {
	ph := NewPokerHand([]string{"2C", "9S", "8H", "5H", "6H"})
	assertString(t, HighCard, ph.GetCombo())
}

func TestPair(t *testing.T) {
	ph := NewPokerHand([]string{"2C", "8H", "5H", "6H", "5D"})
	assertString(t, Pair, ph.GetCombo())
}

func TestTwoPairs(t *testing.T) {
	ph := NewPokerHand([]string{"8C", "8H", "5H", "6H", "5D"})
	assertString(t, TwoPairs, ph.GetCombo())
}

func TestThreeOfAKind(t *testing.T) {
	ph := NewPokerHand([]string{"5H", "5C", "5D", "8C", "9S"})
	assertString(t, ThreeOfAKind, ph.GetCombo())
}

func TestFlush(t *testing.T) {
	ph := NewPokerHand([]string{"8C", "9C", "5C", "6C", "KC"})
	assertString(t, Flush, ph.GetCombo())
}

func TestFourOfAKind(t *testing.T) {
	ph := NewPokerHand([]string{"5H", "5C", "5S", "5D", "8C"})
	assertString(t, FourOfAKind, ph.GetCombo())
}

func TestStraight(t *testing.T) {
	ph := NewPokerHand([]string{"5H", "6C", "7S", "8D", "9C"})
	assertString(t, Straight, ph.GetCombo())
}

func TestStraightFromAceToFive(t *testing.T) {
	ph := NewPokerHand([]string{"AH", "2C", "3S", "4D", "5C"})
	assertString(t, Straight, ph.GetCombo())
}

func assertString(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
