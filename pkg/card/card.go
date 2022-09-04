package card

type Card struct {
	s Suit
	w Weight
}

func (c Card) Suit() Suit {
	return c.s
}

func (c Card) Weight() Weight {
	return c.w
}

func New(s Suit, w Weight) Card {
	return Card{
		s: s,
		w: w,
	}
}

type Suit uint8

const (
	SuitHearts Suit = iota
	SuitSpades
	SuitDiamonds
	SuitClubs
)

type Weight uint8

const (
	WeightSix Weight = iota + 6
	WeightSeven
	WeightEight
	WeightNine
	WeightTen
	WeightJack
	WeightQueen
	WeightKing
	WeightAce
)
