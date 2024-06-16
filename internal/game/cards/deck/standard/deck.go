package standard

import (
	"math/rand"
	"slices"

	"github.com/vasilesk/fool/internal/game/cards/deck"
	"github.com/vasilesk/fool/pkg/card"
)

type stdDeck struct {
	cards []card.Card
	trump card.Card

	pos int
}

func NewDeck() deck.Deck {
	d := newOrdered()

	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})

	return d
}

func (d *stdDeck) TakeMax(lim int) []card.Card {
	cardsLeft := len(d.cards) - d.pos

	if cardsLeft == 0 {
		return nil
	}

	if lim > cardsLeft {
		lim = cardsLeft
	}

	res := slices.Concat(d.cards[d.pos : d.pos+lim])

	d.pos += lim

	return res
}

func (d *stdDeck) Trump() card.Card {
	return d.trump
}

func newOrdered() *stdDeck {
	suits := []card.Suit{
		card.SuitHearts,
		card.SuitSpades,
		card.SuitDiamonds,
		card.SuitClubs,
	}

	weights := []card.Weight{
		card.WeightSix,
		card.WeightSeven,
		card.WeightEight,
		card.WeightNine,
		card.WeightTen,
		card.WeightJack,
		card.WeightQueen,
		card.WeightKing,
		card.WeightAce,
	}

	cards := make([]card.Card, 0, len(suits)*len(weights))

	for _, s := range suits {
		for _, w := range weights {
			cards = append(cards, card.New(s, w))
		}
	}

	return &stdDeck{cards: cards, trump: cards[(len(cards) - 1)], pos: 0}
}
