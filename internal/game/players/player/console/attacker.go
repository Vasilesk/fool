//nolint:forbidigo
package console

import (
	"fmt"
	"slices"

	"github.com/samber/lo"

	"github.com/vasilesk/fool/internal/game/cards/beaten"
	"github.com/vasilesk/fool/pkg/card"
)

type attacker struct {
	*console
}

func newAttacker(c *console) *attacker {
	return &attacker{console: c}
}

func (a *attacker) MakeMove(bb []beaten.Beaten) ([]card.Card, error) { // todo: extract first move?
	sortCards(a.cards, a.trump)

	fmt.Println(a.name, "make your move, trump is", a.trump)

	possibleMoves := a.cards
	if bb != nil {
		possibleMoves = lo.Filter(possibleMoves, func(move card.Card, _ int) bool {
			return lo.SomeBy(beaten.GetCards(bb), func(b card.Card) bool {
				return b.Weight() == move.Weight()
			})
		})
	}

	cards, err := readCardsByNumbers(possibleMoves)
	if err != nil {
		fmt.Println("failed to read cards")

		return nil, fmt.Errorf("reading cards: %w", err)
	}

	fmt.Println("your move is", cards)

	a.cards = lo.Filter(a.cards, func(item card.Card, _ int) bool {
		return !slices.Contains(cards, item)
	})

	return cards, nil
}
