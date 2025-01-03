package poker_hand

import "testing"

func TestPokerHands(t *testing.T) {
	tests := []struct {
		name  string
		hand  []string
		combo string
	}{
		{
			"CanHighCard",
			[]string{"2C", "9S", "8H", "5H", "6H"},
			HighCard,
		},
		{
			"CanPair",
			[]string{"2C", "8H", "5H", "6H", "5D"},
			Pair,
		},
		{
			"CanTwoPairs",
			[]string{"8C", "8H", "5H", "6H", "5D"},
			TwoPairs,
		},
		{
			"CanThreeOfAKind",
			[]string{"5H", "5C", "5D", "8C", "9S"},
			ThreeOfAKind,
		},
		{
			"CanFlush",
			[]string{"8C", "9C", "5C", "6C", "KC"},
			Flush,
		},
		{
			"CanFourOfAKind",
			[]string{"5H", "5C", "5S", "5D", "8C"},
			FourOfAKind,
		},
		{
			"CanStraight",
			[]string{"5H", "6C", "7S", "8D", "9C"},
			Straight,
		},
		{
			"CanStraightFromAceToFive",
			[]string{"AH", "2C", "3S", "4D", "5C"},
			Straight,
		},
		{
			"CanStraightFlush",
			[]string{"AH", "2H", "3H", "4H", "5H"},
			StraightFlush,
		},
		{
			"CanFullHouse",
			[]string{"AC", "AH", "AS", "QH", "QD"},
			FullHouse,
		},
		{
			"CanRoyalFlush",
			[]string{"AC", "KC", "QC", "JC", "TC"},
			RoyalFlush,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ph := NewPokerHand(test.hand)
			assertString(t, test.combo, ph.GetCombo())
		})
	}
}

func assertString(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
