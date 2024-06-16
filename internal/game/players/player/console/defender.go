//nolint:forbidigo
package console

import (
	"fmt"
	"slices"

	"github.com/samber/lo"

	"github.com/vasilesk/fool/internal/game/cards/beaten"
	"github.com/vasilesk/fool/pkg/card"
)

type defender struct {
	*console
}

func newDefender(c *console) *defender {
	return &defender{console: c}
}

func (d *defender) AnswerMove(move []card.Card) ([]card.Card, bool) {
	sortCards(d.cards, d.trump)

	fmt.Println(d.name, "defending from", move, "you have", d.cards, "trump is", d.trump)

	possibleAnswers := lo.Filter(d.cards, func(answer card.Card, _ int) bool {
		return lo.SomeBy(move, func(m card.Card) bool {
			return beaten.Beats(answer, m, d.trump)
		})
	})

	cards, err := readCardsByNumbers(possibleAnswers)
	if err != nil {
		fmt.Println("failed to read cards")

		return nil, false
	}

	if len(cards) == 0 {
		return nil, false
	}

	d.cards = lo.Filter(d.cards, func(item card.Card, _ int) bool {
		return !slices.Contains(cards, item)
	})

	return cards, true
}
