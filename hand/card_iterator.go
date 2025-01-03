package hand

import "iter"

func CardIterator() iter.Seq[Card] {
	return func(yield func(Card) bool) {
		for _, suit := range "HDSC" {
			for _, value := range "23456789TJQKA" {
				if !yield(NewCard(string(value) + string(suit))) {
					return
				}
			}
		}
	}
}
