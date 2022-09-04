package beaten

import (
	"fmt"

	"github.com/vasilesk/fool/pkg/card"
)

type Beaten struct {
	move   card.Card
	answer card.Card
}

func (cc Beaten) GetMove() card.Card {
	return cc.move
}

func (cc Beaten) GetAnswer() card.Card {
	return cc.answer
}

func NewBeaten(move card.Card, answer card.Card, trump card.Suit) (Beaten, error) {
	if move.Suit() != trump && answer.Suit() == trump {
		return Beaten{move: move, answer: answer}, nil
	}

	if answer.Suit() == move.Suit() && answer.Weight() > move.Weight() {
		return Beaten{move: move, answer: answer}, nil
	}

	return Beaten{}, fmt.Errorf("invalid answer %v to card %v", answer, move)
}

func GetCards(bb []Beaten) []card.Card {
	const x2 = 2

	cards := make([]card.Card, 0, len(bb)*x2)

	for _, b := range bb {
		cards = append(cards, b.GetMove(), b.GetAnswer())
	}

	return cards
}
