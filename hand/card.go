package hand

var valueMap = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type Card struct {
	value int
	suit  byte
}

func NewCard(s string) Card {
	return Card{
		value: valueMap[s[0]],
		suit:  s[1],
	}
}