package standard

import (
	"math/rand"
	"time"

	"github.com/vasilesk/fool/internal/game/cards/deck"
	"github.com/vasilesk/fool/pkg/card"
)

type stdDeck struct {
	cards []card.Card
	trump card.Card

	pos int
}

func (d *stdDeck) TakeMax(lim int) []card.Card {
	cardsLeft := len(d.cards) - d.pos

	if cardsLeft == 0 {
		return nil
	}

	if lim > cardsLeft {
		lim = cardsLeft
	}

	res := d.cards[d.pos : d.pos+lim]

	d.pos += lim

	return res
}

func (d *stdDeck) Trump() card.Card {
	return d.trump
}

func New() deck.Deck {
	return NewFromSeed(time.Now().UnixNano())
}

func NewFromSeed(s int64) deck.Deck {
	d := newOrdered()

	rand.Seed(s)

	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})

	return d
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
