package card

import (
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	s Suit
	w Weight
}

func New(s Suit, w Weight) Card {
	return Card{
		s: s,
		w: w,
	}
}

func NewFromString(st string) (Card, error) {
	const representationMinLen = 2

	st = strings.ToUpper(st)

	if len(st) < representationMinLen {
		return Card{}, fmt.Errorf(
			"card string representation should have len %d but it has len %d",
			representationMinLen,
			len(st),
		)
	}

	wStr, sStr := st[0:len(st)-1], st[len(st)-1]

	var w Weight

	switch wStr {
	case "A":
		w = WeightAce
	case "K":
		w = WeightKing
	case "Q":
		w = WeightQueen
	case "J":
		w = WeightJack
	default:
		wInt, err := strconv.Atoi(wStr)
		if err != nil {
			return Card{}, fmt.Errorf("converting weight %q to int: %w", wStr, err)
		}

		if wInt < 6 || wInt > 10 {
			return Card{}, fmt.Errorf("converting weight %q to int: inapropriate value", wStr)
		}

		w = Weight(wInt)
	}

	var s Suit

	switch sStr {
	case 'H':
		s = SuitHearts
	case 'S':
		s = SuitSpades
	case 'D':
		s = SuitDiamonds
	case 'C':
		s = SuitClubs
	default:
		return Card{}, fmt.Errorf("suit %q can't be parsed", s)
	}

	return New(s, w), nil
}

func (c Card) Suit() Suit {
	return c.s
}

func (c Card) Weight() Weight {
	return c.w
}

func (c Card) String() string {
	return c.Weight().String() + c.Suit().String()
}

const (
	SuitHearts Suit = iota
	SuitSpades
	SuitDiamonds
	SuitClubs
)

type Suit uint8

func (s Suit) String() string {
	switch s {
	case SuitHearts:
		return "♥"
	case SuitSpades:
		return "♠"
	case SuitDiamonds:
		return "♦"
	case SuitClubs:
		return "♣"
	}

	return ""
}

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

type Weight uint8

func (w Weight) String() string {
	switch w { //nolint:exhaustive
	case WeightJack:
		return "J"
	case WeightQueen:
		return "Q"
	case WeightKing:
		return "K"
	case WeightAce:
		return "A"
	default:
		return strconv.Itoa(int(w))
	}
}
